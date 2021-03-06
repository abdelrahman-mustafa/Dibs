package models

import (
	"fmt"

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
		Long           float64       `json:"long,omitempty" bson:"long,omitempty"`
		Lat            float64       `json:"lat,omitempty" bson:"lat,omitempty"`
		From           string        `json:"from,omitempty" bson:"from,omitempty"`
		To             string        `json:"to,omitempty" bson:"to,omitempty"`
		FromHour       int           `json:"fromHour,omitempty" bson:"fromHour,omitempty"`
		ToHour         int           `json:"toHour,omitempty" bson:"toHour,omitempty"`
		Price          int           `json:"price,omitempty" bson:"price,omitempty"`
		OriginalPrice  int           `json:"originalPrice,omitempty" bson:"originalPrice,omitempty"`
		Banner         string        `json:"banner,omitempty" bson:"banner,omitempty"`
		Logo           string        `json:"logo,omitempty" bson:"logo,omitempty"`
		Contact        string        `json:"contact,omitempty" bson:"contact,omitempty"`
		Description    string        `json:"description,omitempty" bson:"description,omitempty"`
		DescriptionAR  string        `json:"descriptionAR,omitempty" bson:"descriptionAR,omitempty"`
		IsActive       bool          `json:"isActive,omitempty" bson:"isActive"`
		MinBoxes       int           `json:"minBoxes,omitempty" bson:"minBoxes,omitempty"`
		TomorrowBoxes  int           `json:"tomorrowBoxes,omitempty" bson:"tomorrowBoxes"`
		Location       GeoJSON       `bson:"location,omitempty" json:"location,omitempty"`
		Type           string        `json:"type,omitempty" bson:"type,omitempty"`
		Cateogories    []Cateogory   `json:"cateogories,omitempty" bson:"cateogories,omitempty"`
		IsFavorite     bool          `json:"isFavorite,omitempty"`
		Distance       float64       `json:"distance,omitempty"`
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

//IsBoxExist ... validates the username is for box
func IsBoxExist(username string, Session *mgo.Session) (bool, string, bson.ObjectId) {

	fmt.Println("Start find the box", username)
	box := Box{}
	error := Session.DB("dibs").C("boxes").Find(bson.M{"username": username}).One(&box)
	fmt.Println("Box is found ", box)
	if error != nil {
		return false, "", ""
	}
	if box.Username != "" {
		return true, box.Password, box.ID

	}
	return false, "", ""

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
		IsActive       int           `json:"isActive,omitempty" bson:"isActive,omitempty"`
	}
)
