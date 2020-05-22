package models

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		ID        bson.ObjectId `json:"id,omitempty" bson:"_id"`
		Name      string        `json:"name,omitempty" bson:"name,omitempty"`
		Username  string        `json:"username,omitempty" bson:"username,omitempty"`
		Password  string        `json:"password,omitempty" bson:"password,omitempty"`
		Phone     string        `json:"phone,omitempty" bson:"phone,omitempty"`
		Email     string        `json:"email,omitempty" bson:"email,omitempty"`
		Area      string        `json:"area,omitempty" bson:"area,omitempty"`
		Orders    []string      `json:"orders,omitempty" bson:"orders,omitempty"`
		Favorites []string      `json:"favorites,omitempty" bson:"favorites,omitempty"`
	}
)

//IsUser ... validates the id is for User
func IsUser(id string, Session *mgo.Session) bool {

	oid := bson.ObjectIdHex(id)

	if err := Session.DB("dibs").C("users").FindId(oid); err != nil {
		return false
	}
	return true
}

//GetUser ... validates the id is for User
func GetUser(id string, Session *mgo.Session) User {
	fmt.Println("Start with id  is ", id)

	user := User{}
	oid := bson.ObjectIdHex(id)
	Session.DB("dibs").C("users").FindId(oid).One(&user)
	fmt.Println("Start with data ", user)
	return user
}

//IsUserExist ... validates the id is for User
func IsUserExist(email string, Session *mgo.Session) (bool, string, bson.ObjectId) {

	fmt.Println("Start find the user", email)
	user := User{}
	error := Session.DB("dibs").C("users").Find(bson.M{"email": email}).One(&user)
	fmt.Println("User is found ", user)
	fmt.Println("Start find the user", user.Username)
	if error != nil {
		return false, "", ""
	}

	return true, user.Password, user.ID
}

//GetUserByEmail ... validates the id is for User
func GetUserByEmail(email string, Session *mgo.Session) bool {

	fmt.Println("Start find the user", email)
	user := User{}
	error := Session.DB("dibs").C("users").Find(bson.M{"email": email}).One(&user)
	fmt.Println("User is found ", user)
	fmt.Println("Start find the user", user.Username)
	if error != nil {
		return false
	}

	return true
}
