package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Admin represents the structure of our resource
	Admin struct {
		ID       bson.ObjectId `json:"id" bson:"_id"`
		Username string        `json:"username" bson:"username"`
		Password string        `json:"password" bson:"password"`
		Role     string        `json:"role" bson:"role"`
	}
)

//IsAdmin ... validates the id is for admin
func IsAdmin(id string) bool {

	session := GetSession()
	if err := session.DB("dibs").C("admins").FindId(id); err != nil {
		return false
	}
	return true
}
