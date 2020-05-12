package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Cateogory represents the structure of our resource
	Cateogory struct {
		ID    bson.ObjectId   `json:"id,omitempty" bson:"_id,omitempty"`
		Name  string          `json:"name,omitempty" bson:"name,omitempty"`
		Boxes []bson.ObjectId `json:"boxes,omitempty" bson:"boxes,omitempty"`
	}
)

//IsCateogoryExist ... validates the id is for User
func IsCateogoryExist(id string) bool {
	oid := bson.ObjectIdHex(id)

	session := GetSession()
	cat := Cateogory{}
	session.DB("dibs").C("cateogories").FindId(oid).One(&cat)
	if cat.Name == "" {
		return false
	}
	return true
}
