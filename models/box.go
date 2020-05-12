package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Box represents the structure of our resource
	Box struct {
		ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Name           string        `json:"name,omitempty" bson:"name,omitempty"`
		Username       string        `json:"username,omitempty" bson:"username,omitempty"`
		Password       string        `json:"password,omitempty" bson:"password,omitempty"`
		AvailableBoxes int           `json:"availableBoxes,omitempty" bson:"availableBoxes,omitempty"`
		Long           float64       `json:"long,omitempty" bson:"long,omitempty"`
		Lat            float64       `json:"lat,omitempty" bson:"lat,omitempty"`
		From           int           `json:"from,omitempty" bson:"from,omitempty"`
		To             int           `json:"to,omitempty" bson:"to,omitempty"`
		Price          int           `json:"price,omitempty" bson:"price,omitempty"`
		Banner         string        `json:"banner,omitempty" bson:"banner,omitempty"`
		Logo           string        `json:"logo,omitempty" bson:"logo,omitempty"`
		Contact        string        `json:"contact,omitempty" bson:"contact,omitempty"`
		Description    string        `json:"description,omitempty" bson:"description,omitempty"`
		IsActive       bool          `json:"isActive,omitempty" bson:"isActive,omitempty"`
	}
)

//IsBox ... validates the id is for box
func IsBox(id string) bool {
	oid := bson.ObjectIdHex(id)
	session := GetSession()
	box := Box{}
	session.DB("dibs").C("Boxes").FindId(oid).One(&box)

	if box.Name != "" || box.Username != "" {
		return true
	} else {
		return false
	}
}

type (

	//UserBox ... UserBox
	UserBox struct {
		ID             bson.ObjectId `json:"id,omitempty" bson:"_id" `
		Name           string        `json:"name,omitempty" bson:"name,omitempty"`
		Username       string        `json:"-" bson:"username,omitempty"`
		Password       string        `json:"-" bson:"password,omitempty"`
		AvailableBoxes int           `json:"availableBoxes,omitempty" bson:"availableBoxes,omitempty"`
		Long           string        `json:"long,omitempty" bson:"long,omitempty"`
		Lat            string        `json:"lat,omitempty" bson:"lat,omitempty"`
		From           string        `json:"from,omitempty" bson:"from,omitempty"`
		To             string        `json:"to,omitempty" bson:"to,omitempty"`
		Price          int           `json:"price,omitempty" bson:"price,omitempty"`
		Banner         string        `json:"banner,omitempty" bson:"banner,omitempty"`
		Logo           string        `json:"logo,omitempty" bson:"logo,omitempty"`
		Contact        string        `json:"contact,omitempty" bson:"contact,omitempty"`
		Description    string        `json:"description,omitempty" bson:"description,omitempty"`
		IsActive       bool          `json:"isActive,omitempty" bson:"isActive,omitempty"`
	}
)
