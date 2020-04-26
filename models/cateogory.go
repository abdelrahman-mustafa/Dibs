package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Cateogory represents the structure of our resource
	Cateogory struct {
		ID    bson.ObjectId `json:"id" bson:"_id"`
		Name  string        `json:"name" bson:"name"`
		Boxes []string      `json:"boxes" bson:"boxes"`
	}
)
