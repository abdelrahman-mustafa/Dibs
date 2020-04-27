package models

import (
	"fmt"

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

//IsAdminExist ... validates the username is for admin
func IsAdminExist(username string) (bool, string, bson.ObjectId, string) {

	fmt.Println("Start find the admin", username)
	session := GetSession()
	admin := Admin{}
	error := session.DB("dibs").C("admins").Find(bson.M{"username": username}).One(&admin)
	fmt.Println("Admin is found ", admin)
	fmt.Println("Start find the Admin", admin.Username)
	if error != nil {
		return false, "", "", ""
	}

	return true, admin.Password, admin.ID, admin.Role
}
