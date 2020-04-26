package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		ID        bson.ObjectId `json:"id,omitempty" bson:"_id" `
		Name      string        `json:"name,omitempty" bson:"name"`
		Username  string        `json:"username,omitempty" bson:"username"`
		Password  string        `json:"password,omitempty" bson:"password"`
		Phone     int64         `json:"phone,omitempty" bson:"phone"`
		Email     string        `json:"email,omitempty" bson:"email"`
		Area      string        `json:"area,omitempty" bson:"area"`
		Orders    []string      `json:"orders,omitempty" bson:"orders"`
		Favorites []string      `json:"favorites,omitempty" bson:"favorites"`
	}
)

//IsUser ... validates the id is for User
func IsUser(id string) bool {

	oid := bson.ObjectIdHex(id)

	session := GetSession()
	if err := session.DB("dibs").C("users").FindId(oid); err != nil {
		return false
	}
	return true
}

//GetUser ... validates the id is for User
func GetUser(id string) User {
	fmt.Println("Start with id  is ", id)

	session := GetSession()
	user := User{}
	oid := bson.ObjectIdHex(id)
	session.DB("dibs").C("users").FindId(oid).One(&user)
	fmt.Println("Start with data ", user)
	return user
}

//IsUserExist ... validates the id is for User
func IsUserExist(username string) (bool, string, bson.ObjectId) {

	fmt.Println("Start find the user", username)
	session := GetSession()
	user := User{}
	error := session.DB("dibs").C("users").Find(bson.M{"username": username}).One(&user)
	fmt.Println("User is found ", user)
	fmt.Println("Start find the user", user.Username)
	if error != nil {
		return false, "", ""
	}

	return true, user.Password, user.ID
}
