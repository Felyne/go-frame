package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Repo
type Tea struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
}

type TeasCollection struct {
	Data []Tea `json:"data"`
}

type TeaResource struct {
	Data Tea `json:"data"`
}

type TeaRepo struct {
	coll *mgo.Collection
}

func (r *TeaRepo) All() (TeasCollection, error) {
	result := TeasCollection{[]Tea{}}
	err := r.coll.Find(nil).All(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *TeaRepo) Find(id string) (TeaResource, error) {
	result := TeaResource{}
	err := r.coll.FindId(bson.ObjectIdHex(id)).One(&result.Data)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *TeaRepo) Create(tea *Tea) error {
	id := bson.NewObjectId()
	_, err := r.coll.UpsertId(id, tea)
	if err != nil {
		return err
	}

	tea.Id = id

	return nil
}

func (r *TeaRepo) Update(tea *Tea) error {
	err := r.coll.UpdateId(tea.Id, tea)
	if err != nil {
		return err
	}

	return nil
}

func (r *TeaRepo) Delete(id string) error {
	err := r.coll.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}

	return nil
}
