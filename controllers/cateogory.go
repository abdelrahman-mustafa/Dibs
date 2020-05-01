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

	// CatController represents the controller for operating on the Cat resource
	CatController struct {
		session *mgo.Session
	}
)

// NewCatController ... returns a instance of UserController structure
func NewCatController(s *mgo.Session) *CatController {
	return &CatController{s}
}

// CreateCat ... creates a new Cat resource
func (ad CatController) CreateCat(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	SetupResponse(&w, r)

	Cat := models.Cateogory{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Cat)

	//create id
	Cat.ID = bson.NewObjectId()

	// write struct of admni to DB
	ad.session.DB("dibs").C("Cateogories").Insert(Cat)

	// convert struct to JSON
	output, _ := json.Marshal(Cat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

	// fmt.Fprintf(w, "%s", uj)
}
