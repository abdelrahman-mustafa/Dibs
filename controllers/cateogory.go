package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"../helpers"
	"../models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

type (

	// CatController represents the controller for operating on the Cat resource
	CatController struct {
	}
)

// NewCatController ... returns a instance of UserController structure
func NewCatController() *CatController {
	return &CatController{}
}

// CreateCateogory ... creates a new Cat resource
func (ad CatController) CreateCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	Cat := models.Cateogory{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Cat)

	//create id
	Cat.ID = bson.NewObjectId()

	// write struct of admni to DB

	for _, id := range Cat.Boxes {
		println("Start update box", id)
		err := models.Session.DB("dibs").C("Boxes").UpdateId(id, bson.M{
			"$addToSet": bson.M{
				"cateogories": Cat,
			},
		})
		if err != nil {
			panic(err)
		}

	}
	models.Session.DB("dibs").C("cateogories").Insert(Cat)

	// convert struct to JSON
	// output, _ := json.Marshal(Cat)
	res := helpers.ResController{Res: w}
	res.SendResponse("the cateogory is created", 200)
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
	oid := bson.ObjectIdHex(p.ByName("id"))

	if len(cat.Boxes) != 0 {
		for _, id := range cat.Boxes {
			models.Session.DB("dibs").C("Boxes").UpdateId(id, bson.M{
				"$addToSet": bson.M{
					"cateogories": cat,
				},
			})

		}
	}

	out := bson.M{"$set": cat}
	models.Session.DB("dibs").C("cateogories").UpdateId(oid, out)

	res := helpers.ResController{Res: w}
	res.SendResponse("The cateogory is updated successfully", 200)
}

// GetCateogory ... get  Cat resource
func (ad CatController) GetCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cateogoryID := p.ByName("id")
	println("Finding id", cateogoryID)

	isExist := models.IsCateogoryExist(cateogoryID)
	if isExist == false {
		res := helpers.ResController{Res: w}
		res.SendResponse("It is not valid cateogory", 404)
		return
	}
	oid := bson.ObjectIdHex(cateogoryID)

	cat := models.Cateogory{}
	models.Session.DB("dibs").C("cateogories").FindId(oid).One(&cat)
	output, _ := json.Marshal(cat)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}

// GetCateogories ... get  Cat resource
func (ad CatController) GetCateogories(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// cat := []models.Cateogory{}
	var results []bson.M

	err := models.Session.DB("dibs").C("cateogories").Pipe([]bson.M{
		{
			"$lookup": bson.M{
				"from":         "Boxes",
				"localField":   "boxes",
				"foreignField": "_id",
				"as":           "boxes",
			},
		},
	}).All(&results)

	if err != nil {
		panic(err)
	}
	output, _ := json.Marshal(results)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}

// DeleteCateogory ... delete  Cat resource
func (ad CatController) DeleteCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cateogoryID := p.ByName("id")
	res := helpers.ResController{Res: w}
	isExist := models.IsCateogoryExist(cateogoryID)
	if isExist == false {
		res.SendResponse("It is not valid cateogory", 404)
		return
	}
	oid := bson.ObjectIdHex(cateogoryID)

	err := models.Session.DB("dibs").C("cateogories").RemoveId(oid)

	if err != nil {
		res.SendResponse("There is a problem, try later", 401)
		return
	}
	res.SendResponse("The cateogory is deleted successfully", 200)
}
