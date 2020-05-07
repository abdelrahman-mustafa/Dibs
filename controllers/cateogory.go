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

// CreateCateogory ... creates a new Cat resource
func (ad CatController) CreateCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	Cat := models.Cateogory{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Cat)

	//create id
	Cat.ID = bson.NewObjectId()

	// write struct of admni to DB
	ad.session.DB("dibs").C("cateogories").Insert(Cat)

	// convert struct to JSON
	output, _ := json.Marshal(Cat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

	// fmt.Fprintf(w, "%s", uj)
}

// UpdateCateogory ... update  Cat resource
func (ad CatController) UpdateCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cateogoryID := p.ByName("id")
	isExist := models.IsCateogoryExist(cateogoryID)
	if isExist == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "It is not valid cateogory")
		return
	}
	cat := models.Cateogory{}
	json.NewDecoder(r.Body).Decode(&cat)

	ad.session.DB("dibs").C("cateogories").UpdateId(cateogoryID, cat)

	w.Header().Set("Content-Type", "appliBoxion/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", "The cateogory is updated successfully")

}

// GetCateogory ... get  Cat resource
func (ad CatController) GetCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cateogoryID := p.ByName("id")
	println("Finding id", cateogoryID)

	isExist := models.IsCateogoryExist(cateogoryID)
	if isExist == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "It is not valid cateogory")
		return
	}
	oid := bson.ObjectIdHex(cateogoryID)

	cat := models.Cateogory{}
	ad.session.DB("dibs").C("cateogories").FindId(oid).One(&cat)
	output, _ := json.Marshal(cat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}

// GetCateogories ... get  Cat resource
func (ad CatController) GetCateogories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// cat := []models.Cateogory{}
	var results []bson.M

	ad.session.DB("dibs").C("cateogories").Pipe([]bson.M{
		{
			"$lookup": bson.M{
				"from":         "Boxes",
				"localField":   "boxes",
				"foreignField": "_id",
				"as":           "boxes",
			},
		},
	}).All(&results)

	output, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}

// DeleteCateogory ... delete  Cat resource
func (ad CatController) DeleteCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cateogoryID := p.ByName("id")
	println("Finding id", cateogoryID)

	isExist := models.IsCateogoryExist(cateogoryID)
	if isExist == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "It is not valid cateogory")
		return
	}
	oid := bson.ObjectIdHex(cateogoryID)

	err := ad.session.DB("dibs").C("cateogories").RemoveId(oid)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "There is a problem")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", "The cateogory is deleted successfully")
}
