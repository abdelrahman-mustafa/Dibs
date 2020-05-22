package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//GeoJSON ...
type GeoJSON struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}

type (
	// Box represents the structure of our resource
	Box struct {
		ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Name           string        `json:"name,omitempty" bson:"name,omitempty"`
		NameAR         string        `json:"nameAR,omitempty" bson:"nameAR,omitempty"`
		Username       string        `json:"username,omitempty" bson:"username,omitempty"`
		Password       string        `json:"password,omitempty" bson:"password,omitempty"`
		AvailableBoxes int           `json:"availableBoxes,omitempty" bson:"availableBoxes,omitempty"`
		Long           float64       `json:"long,omitempty," bson:"-"`
		Lat            float64       `json:"lat,omitempty" bson:"-"`
		From           int           `json:"from,omitempty" bson:"from,omitempty"`
		To             int           `json:"to,omitempty" bson:"to,omitempty"`
		Price          int           `json:"price,omitempty" bson:"price,omitempty"`
		Banner         string        `json:"banner,omitempty" bson:"banner,omitempty"`
		Logo           string        `json:"logo,omitempty" bson:"logo,omitempty"`
		Contact        string        `json:"contact,omitempty" bson:"contact,omitempty"`
		Description    string        `json:"description,omitempty" bson:"description,omitempty"`
		DescriptionAR  string        `json:"descriptionAR,omitempty" bson:"descriptionAR,omitempty"`
		IsActive       bool          `json:"isActive,omitempty" bson:"isActive,omitempty"`
		Location       GeoJSON       `bson:"location,omitempty" json:"location,omitempty"`
		Cateogories    []Cateogory   `json:"cateogories,omitempty" bson:"cateogories,omitempty"`
	}
)

//IsBox ... validates the id is for box
func IsBox(id string, Session *mgo.Session) bool {
	oid := bson.ObjectIdHex(id)
	box := Box{}
	Session.DB("dibs").C("boxes").FindId(oid).One(&box)

	if box.Name != "" || box.Username != "" {
		return true
	}
	return false

}

type (

	//UserBox ... UserBox
	UserBox struct {
		ID             bson.ObjectId `json:"id,omitempty" bson:"_id" `
		Name           string        `json:"name,omitempty" bson:"name,omitempty"`
		Username       string        `json:"-" bson:"username,omitempty"`
		Password       string        `json:"-" bson:"password,omitempty"`
		AvailableBoxes int           `json:"availableBoxes,omitempty" bson:"availableBoxes,omitempty"`
		Long           uint64        `json:"long,omitempty" bson:"long,omitempty"`
		Lat            uint64        `json:"lat,omitempty" bson:"lat,omitempty"`
		From           string        `json:"from,omitempty" bson:"from,omitempty"`
		To             string        `json:"to,omitempty" bson:"to,omitempty"`
		Price          float32       `json:"price,omitempty" bson:"price,omitempty"`
		Banner         string        `json:"banner,omitempty" bson:"banner,omitempty"`
		Logo           string        `json:"logo,omitempty" bson:"logo,omitempty"`
		Contact        string        `json:"contact,omitempty" bson:"contact,omitempty"`
		Description    string        `json:"description,omitempty" bson:"description,omitempty"`
		IsActive       bool          `json:"isActive,omitempty" bson:"isActive,omitempty"`
	}
)
