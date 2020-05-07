package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Box represents the structure of our resource
	Box struct {
		ID             bson.ObjectId `json:"id,omitempty" bson:"_id" `
		Name           string        `json:"name,omitempty" bson:"name"`
		Username       string        `json:"username,omitempty" bson:"username"`
		Password       string        `json:"password,omitempty" bson:"password"`
		AvailableBoxes int           `json:"availableBoxes,omitempty" bson:"availableBoxes"`
		Long           string        `json:"long,omitempty" bson:"long"`
		Lat            string        `json:"lat,omitempty" bson:"lat"`
		From           string        `json:"from,omitempty" bson:"from"`
		To             string        `json:"to,omitempty" bson:"to"`
		Price          int           `json:"price,omitempty" bson:"price"`
		Banner         string        `json:"banner,omitempty" bson:"banner"`
		Logo           string        `json:"logo,omitempty" bson:"logo"`
		Contact        string        `json:"contact,omitempty" bson:"contact"`
		Description    string        `json:"description,omitempty" bson:"description"`
		IsActive       bool          `json:"isActive,omitempty" bson:"isActive"`
	}
)

//IsBox ... validates the id is for box
func IsBox(id string) bool {
	oid := bson.ObjectIdHex(id)
	session := GetSession()
	if err := session.DB("dibs").C("boxes").FindId(oid); err != nil {
		return false
	}
	return true
}
