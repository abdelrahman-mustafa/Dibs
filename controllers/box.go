package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"../helpers"
	"../models"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (

	// BoxController represents the controller for operating on the Box resource
	BoxController struct {
		session *mgo.Session
	}
)

// NewBoxController ... returns a instance of UserController structure
func NewBoxController(s *mgo.Session) *BoxController {
	return &BoxController{s}
}

// CreateBox ... creates a new Box resource
func (ad BoxController) CreateBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	Box := models.Box{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Box)

	//create id
	Box.ID = bson.NewObjectId()

	encryptedPassword, _ := helpers.Encrypt(Box.Password)
	Box.Password = encryptedPassword
	// write struct of admni to DB
	ad.session.DB("dibs").C("Boxes").Insert(Box)

	// convert struct to JSON
	output, _ := json.Marshal(Box)
	w.Header().Set("Content-Type", "appliBoxion/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

	// fmt.Fprintf(w, "%s", uj)
}

// UpdateBox ... updates a new Box resource
func (ad BoxController) UpdateBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isExist := models.IsBox(p.ByName("id"))
	if isExist == false {
		w.Header().Set("Content-Type", "appliBoxion/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "Not valid Box")
		return
	}
	// validate id and return the object of id

	// edit  the new changes in the object
	// update the doc in DB
	Box := models.Box{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Box)
	oid := bson.ObjectIdHex(p.ByName("id"))

	// write struct of admni to DB
	ad.session.DB("dibs").C("Boxes").UpdateId(oid, Box)

	// convert struct to JSON
	output, _ := json.Marshal(Box)
	w.Header().Set("Content-Type", "appliBoxion/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

	// fmt.Fprintf(w, "%s", uj)
}

// GetBoxes ... get  box resource
func (ad BoxController) GetBoxes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// cat := []models.Cateogory{}
	var results []bson.M

	ad.session.DB("dibs").C("Boxes").Find(bson.M{}).All(&results)
	output, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}

// GetBox ... updates a new Box resource
func (ad BoxController) GetBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// validate id and return the object of id

	// edit  the new changes in the object
	// update the doc in DB
	isExist := models.IsBox(p.ByName("id"))
	if isExist == false {
		w.Header().Set("Content-Type", "appliBoxion/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "Not valid Box")
		return
	}
	Box := models.Box{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Box)

	oid := bson.ObjectIdHex(p.ByName("id"))
	// write struct of admni to DB
	ad.session.DB("dibs").C("Boxes").FindId(oid).One(&Box)

	// convert struct to JSON
	output, _ := json.Marshal(Box)
	w.Header().Set("Content-Type", "appliBoxion/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

	// fmt.Fprintf(w, "%s", uj)
}

// DeleteBox ...  Delete Box resource
func (ad BoxController) DeleteBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	isExist := models.IsBox(p.ByName("id"))
	if isExist == false {
		w.Header().Set("Content-Type", "appliBoxion/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "Not valid Box")
		return
	}

	oid := bson.ObjectIdHex(p.ByName("id"))
	ad.session.DB("dibs").C("Boxes").RemoveId(oid)

	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", "Deleted")

}
