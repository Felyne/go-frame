package main

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	appC := appContext{session.DB("test")}
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler, acceptHandler)
	router := NewRouter()
	router.Get("/teas/:id", commonHandlers.ThenFunc(appC.teaHandler))
	router.Put("/teas/:id", commonHandlers.Append(contentTypeHandler, bodyHandler(TeaResource{})).ThenFunc(appC.updateTeaHandler))
	router.Delete("/teas/:id", commonHandlers.ThenFunc(appC.deleteTeaHandler))
	router.Get("/teas", commonHandlers.ThenFunc(appC.teasHandler))
	router.Post("/teas", commonHandlers.Append(contentTypeHandler, bodyHandler(TeaResource{})).ThenFunc(appC.createTeaHandler))
	http.ListenAndServe(":8080", router)
}
