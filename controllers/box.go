package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"net/http"

	"../helpers"
	"../models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (

	// BoxController represents the controller for operating on the Box resource
	BoxController struct {
		session *mgo.Session
	}
)

// NewBoxController ... returns a instance of UserController structure
func NewBoxController(session *mgo.Session) *BoxController {
	return &BoxController{session}
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
	Box.Location.Type = "Point"
	Box.Location.Coordinates = []float64{Box.Long, Box.Lat}
	// write struct of admni to DB

	println("Data", Box.Password)
	err := ad.session.DB("dibs").C("boxes").Insert(Box)
	if err != nil {
		panic(err)
	}

	index := mgo.Index{
		Key:  []string{"$2dsphere:location"},
		Bits: 26,
	}

	err = ad.session.DB("dibs").C("boxes").EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	// convert struct to JSON

	res := helpers.ResController{Res: w}
	res.SendResponse("The Box is created successfully", 200)
	// fmt.Fprintf(w, "%s", uj)
}

// UpdateBox ... updates a new Box resource
func (ad BoxController) UpdateBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	isExist := models.IsBox(p.ByName("id"), ad.session)
	if isExist == false {
		res := helpers.ResController{Res: w}
		res.SendResponse("Not valid Box", 401)
		return
	}
	// validate id and return the object of id

	// edit  the new changes in the object
	// update the doc in DB
	Box := models.Box{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Box)

	if Box.Password != "" {
		encryptedPassword, _ := helpers.Encrypt(Box.Password)
		Box.Password = encryptedPassword
	}
	oid := bson.ObjectIdHex(p.ByName("id"))
	out := bson.M{"$set": Box}

	if Box.Lat != 0 {
		Box.Location.Type = "Point"
		Box.Location.Coordinates = []float64{Box.Lat, Box.Long}
	}

	// write struct of admni to DB
	ad.session.DB("dibs").C("boxes").UpdateId(oid, out)

	res := helpers.ResController{Res: w}
	res.SendResponse("The Box is updated successfully", 200)

	// fmt.Fprintf(w, "%s", uj)
}

// GetBoxes ... get  box resource
func (ad BoxController) GetBoxes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// cat := []models.Cateogory{}
	var results []bson.M

	err := ad.session.DB("dibs").C("boxes").Pipe([]bson.M{
		{
			"$geoNear": bson.M{
				"near":          bson.M{"type": "Point", "coordinates": []float64{139.701642, 35.690647}},
				"maxDistance":   200000000000,
				"distanceField": "dist.location",
				"spherical":     true,
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

// GetBox ... updates a new Box resource
func (ad BoxController) GetBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// validate id and return the object of id

	// edit  the new changes in the object
	// update the doc in DB
	isExist := models.IsBox(p.ByName("id"), ad.session)
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
	ad.session.DB("dibs").C("boxes").FindId(oid).One(&Box)

	// convert struct to JSON
	output, _ := json.Marshal(Box)
	w.Header().Set("Content-Type", "appliBoxion/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

	// fmt.Fprintf(w, "%s", uj)
}

// GetBoxesByCateogory ... get  box resource
func (ad BoxController) GetBoxesByCateogory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	queryValues := r.URL.Query()

	lat, _ := strconv.ParseFloat(queryValues.Get("lat"), 64)
	long, _ := strconv.ParseFloat(queryValues.Get("long"), 64)

	var results []bson.M
	oid := bson.ObjectIdHex(p.ByName("id"))

	err := ad.session.DB("dibs").C("boxes").Pipe([]bson.M{
		{
			"$geoNear": bson.M{
				"near":          bson.M{"type": "Point", "coordinates": []float64{long, lat}},
				"distanceField": "distance",
				"spherical":     true,
			},
		},
		{
			"$match": bson.M{"cateogories._id": oid},
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

// DeleteBox ...  Delete Box resource
func (ad BoxController) DeleteBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	isExist := models.IsBox(p.ByName("id"), ad.session)
	res := helpers.ResController{Res: w}
	if isExist == false {
		res.SendResponse("Not valid Box", 401)
		return
	}

	oid := bson.ObjectIdHex(p.ByName("id"))
	ad.session.DB("dibs").C("boxes").RemoveId(oid)

	res.SendResponse("The box is Deleted successfully", 200)

}
