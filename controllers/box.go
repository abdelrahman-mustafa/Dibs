package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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

//BoxSignIn ...
type BoxSignIn struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// NewBoxController ... returns a instance of UserController structure
func NewBoxController(session *mgo.Session) *BoxController {
	return &BoxController{session}
}

// CreateBox ... creates a new Box resource
func (ad BoxController) CreateBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	Box := models.Box{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Box)

	isExist, _, _ := models.IsBoxExist(Box.Username, ad.session)
	if isExist {
		w.Header().Set("Content-Type", "appliBoxion/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "There is a user with that name please create a unique username")
		return
	}

	//create id
	Box.ID = bson.NewObjectId()

	encryptedPassword, _ := helpers.Encrypt(Box.Password)
	Box.Password = encryptedPassword
	Box.Location.Type = "Point"
	Box.Location.Coordinates = []float64{Box.Long, Box.Lat}
	Box.AvailableBoxes = Box.MinBoxes
	Box.TomorrowBoxes = Box.MinBoxes

	// write struct of admni to DB

	println("Data", Box.Password)

	layOut := "2006-01-02T15:04:05Z"
	toTimeStamp, errTo := time.Parse(layOut, Box.To)

	if errTo != nil {
		fmt.Println(errTo)
		Box.ToHour = 0
	}

	tohour, _, _ := toTimeStamp.Clock()
	Box.ToHour = tohour
	fromTimeStamp, errFrom := time.Parse(layOut, Box.From)

	if errFrom != nil {
		fmt.Println(errFrom)
		Box.FromHour = 0
	}

	fromHour, _, _ := fromTimeStamp.Clock()
	Box.FromHour = fromHour
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

// SignIn ... updates a new Box resource
func (ad BoxController) SignIn(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	boxAdmin := SignIn{}
	json.NewDecoder(r.Body).Decode(&boxAdmin)
	isExist, pass, id := models.IsBoxExist(boxAdmin.Username, ad.session)
	if isExist == false {
		w.Header().Set("Content-Type", "appliBoxion/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "Not valid Box")
		return
	}
	err := helpers.Compare(pass, boxAdmin.Password)
	if err != nil {
		res := helpers.ResController{Res: w}

		res.SendResponse("Not valid password", 401)
		return
	}

	token := helpers.GenerateToken(id, "box")
	Res := SignInResponse{}
	Res.ID = id
	Res.Token = token
	output, _ := json.Marshal(Res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

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

	if Box.Lat != 0 {
		println("LAT", Box.Lat)
		Box.Location.Type = "Point"
		Box.Location.Coordinates = []float64{Box.Long, Box.Lat}
	}
	println("LAT", (Box.Location.Coordinates))

	if Box.AvailableBoxes != 0 {
		Box.MinBoxes = Box.AvailableBoxes
	}

	layOut := "2006-01-02T15:04:05"
	if Box.To != "" {
		toTimeStamp, errTo := time.Parse(layOut, Box.To)

		if errTo != nil {
			fmt.Println(errTo)
			Box.ToHour = 0
		}

		tohour, _, _ := toTimeStamp.Clock()
		Box.ToHour = tohour
	}
	if Box.From != "" {
		fromTimeStamp, errFrom := time.Parse(layOut, Box.From)

		if errFrom != nil {
			fmt.Println(errFrom)
			Box.FromHour = 0
		}

		fromHour, _, _ := fromTimeStamp.Clock()
		Box.FromHour = fromHour
	}

	out := bson.M{"$set": Box}

	// write struct of admni to DB
	ad.session.DB("dibs").C("boxes").UpdateId(oid, out)

	res := helpers.ResController{Res: w}
	res.SendResponse("The Box is updated successfully", 200)

	// fmt.Fprintf(w, "%s", uj)
}

// UpdateActivationBox ... updates a new Box resource
func (ad BoxController) UpdateActivationBox(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

	oid := bson.ObjectIdHex(p.ByName("id"))
	out := bson.M{"$set": Box}

	// write struct of admni to DB
	println(Box.IsActive)
	ad.session.DB("dibs").C("boxes").UpdateId(oid, out)

	res := helpers.ResController{Res: w}
	res.SendResponse("The Box is updated successfully", 200)

	// fmt.Fprintf(w, "%s", uj)
}

type queryBody struct {
	From int      `json:"from,omitempty"`
	To   int      `json:"to,omitempty"`
	Type []string `json:"type,omitempty"`
	Name string   `json:"name,omitempty"`
}

// SearchBoxes ... get  box resource
func (ad BoxController) SearchBoxes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// cat := []models.Cateogory{}
	queryValues := r.URL.Query()

	lat, _ := strconv.ParseFloat(queryValues.Get("lat"), 64)
	long, _ := strconv.ParseFloat(queryValues.Get("long"), 64)

	queryBox := queryBody{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&queryBox)

	var andQuery []bson.M
	var query bson.M
	andQuery = append(andQuery, bson.M{"isActive": true})
	if queryBox.From != 0 {
		println("from", queryBox.From)
		andQuery = append(andQuery, bson.M{"fromHour": bson.M{
			"$lte": queryBox.To,
		}})
		andQuery = append(andQuery, bson.M{"fromHour": bson.M{
			"$gte": queryBox.From,
		}})

	}
	var cateogories []bson.ObjectId

	if len(queryBox.Type) > 0 {

		for _, item := range queryBox.Type {
			cateogories = append(cateogories, bson.ObjectIdHex(item))
		}
		andQuery = append(andQuery, bson.M{"type": bson.M{
			"$in": cateogories,
		}})
	}

	if queryBox.Name != "" {
		andQuery = append(andQuery, bson.M{"name": bson.RegEx{Pattern: queryBox.Name, Options: "i"}})
	}

	var results []bson.M
	query = bson.M{
		"$and": andQuery,
	}
	err := ad.session.DB("dibs").C("boxes").Pipe([]bson.M{
		{
			"$geoNear": bson.M{
				"near":          bson.M{"type": "Point", "coordinates": []float64{long, lat}},
				"distanceField": "distance",
				"spherical":     true,
			},
		},
		{
			"$match": query,
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

// GetBoxesByUser ... get  box resource
func (ad BoxController) GetBoxesByUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// cat := []models.Cateogory{}
	queryValues := r.URL.Query()

	lat, _ := strconv.ParseFloat(queryValues.Get("lat"), 64)
	long, _ := strconv.ParseFloat(queryValues.Get("long"), 64)
	var results []bson.M

	err := ad.session.DB("dibs").C("boxes").Pipe([]bson.M{
		{
			"$geoNear": bson.M{
				"near":          bson.M{"type": "Point", "coordinates": []float64{long, lat}},
				"distanceField": "distance",
				"spherical":     true,
			},
		},
		{
			"$match": bson.M{
				"isActive": true,
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

// GetBoxes ... get  box resource
func (ad BoxController) GetBoxes(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// cat := []models.Cateogory{}
	queryValues := r.URL.Query()

	lat, _ := strconv.ParseFloat(queryValues.Get("lat"), 64)
	long, _ := strconv.ParseFloat(queryValues.Get("long"), 64)
	var results []bson.M

	err := ad.session.DB("dibs").C("boxes").Pipe([]bson.M{
		{
			"$geoNear": bson.M{
				"near":          bson.M{"type": "Point", "coordinates": []float64{long, lat}},
				"distanceField": "distance",
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

	isExist, cat := models.IsCateogoryExist(p.ByName("id"), ad.session)
	if isExist != false && cat.Name != "" && cat.Name == "In Your Hood" {
		err := ad.session.DB("dibs").C("boxes").Pipe([]bson.M{
			{
				"$geoNear": bson.M{
					"near":          bson.M{"type": "Point", "coordinates": []float64{long, lat}},
					"distanceField": "distance",
					"maxDistance":   5000,
					"spherical":     true,
				},
			},
		}).All(&results)
		if err != nil {
			panic(err)
		}
	} else if isExist != false && cat.Name != "" && cat.Name == "JUST Missed" {
		err := ad.session.DB("dibs").C("boxes").Pipe([]bson.M{
			{
				"$geoNear": bson.M{
					"near":          bson.M{"type": "Point", "coordinates": []float64{long, lat}},
					"distanceField": "distance",
					"maxDistance":   5000,
					"spherical":     true,
				},
			},
			{
				"$match": bson.M{"availableBoxes": 0},
			},
		}).All(&results)
		if err != nil {
			panic(err)
		}
	} else {
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
