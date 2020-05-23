package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Cateogory represents the structure of our resource
	Cateogory struct {
		ID      bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`
		Name    string          `json:"name,omitempty" bson:"name,omitempty"`
		NameAR  string          `json:"nameAR,omitempty" bson:"nameAR,omitempty"`
		IsFirst bool            `json:"isFirst,omitempty" bson:"isFirst,omitempty"`
		Boxes   []bson.ObjectId `json:"boxes,omitempty" bson:"boxes,omitempty"`
	}
)

//IsCateogoryExist ... validates the id is for User
func IsCateogoryExist(id string, Session *mgo.Session) (bool, Cateogory) {
	oid := bson.ObjectIdHex(id)

	cat := Cateogory{}
	Session.DB("dibs").C("cateogories").FindId(oid).One(&cat)
	if cat.Name == "" {
		return false, cat
	}
	return true, cat
}
