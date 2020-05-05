package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Cateogory represents the structure of our resource
	Cateogory struct {
		ID    bson.ObjectId   `json:"id,omitempty" bson:"_id"`
		Name  string          `json:"name,omitempty" bson:"name"`
		Boxes []bson.ObjectId `json:"boxes,omitempty" bson:"boxes"`
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
