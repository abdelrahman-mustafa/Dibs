package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

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
	SetupResponse(&w, r)

	//create id
	Box.ID = bson.NewObjectId()

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

	// validate id and return the object of id
	SetupResponse(&w, r)

	// edit  the new changes in the object
	// update the doc in DB
	Box := models.Box{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Box)

	// write struct of admni to DB
	ad.session.DB("dibs").C("Boxes").UpdateId(p.ByName("id"), Box)

	// convert struct to JSON
	output, _ := json.Marshal(Box)
	w.Header().Set("Content-Type", "appliBoxion/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

	// fmt.Fprintf(w, "%s", uj)
}
