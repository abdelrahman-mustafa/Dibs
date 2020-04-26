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

// SignIn ... sign in as Admin
type SignIn struct {
	Username bson.ObjectId `json:"username" bson:"username"`
	Role     string        `json:"role" bson:"role"`
}

type (

	// AdminController represents the controller for operating on the admin resource
	AdminController struct {
		session *mgo.Session
	}
)

// NewAdminController ... returns a instance of UserController structure
func NewAdminController(s *mgo.Session) *AdminController {
	return &AdminController{s}
}

// CreateAdmin ... creates a new admin resource
func (ad AdminController) CreateAdmin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	tokn := r.Header.Get("Authorization")

	_, err := helpers.VerifyToken(tokn)

	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Missing token"))
	} else {

		admin := models.Admin{}
		//prase json  of body and attach to admoin struct
		json.NewDecoder(r.Body).Decode(&admin)

		//create id
		admin.ID = bson.NewObjectId()

		// write struct of admni to DB
		ad.session.DB("dibs").C("admins").Insert(admin)

		// convert struct to JSON
		output, _ := json.Marshal(admin)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		fmt.Fprintf(w, "%s", output)
	}
	// fmt.Fprintf(w, "%s", uj)
}

// Signin ... sign in as Admin
func (ad AdminController) Signin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Start Sign in")
	signIn := SignIn{}
	json.NewDecoder(r.Body).Decode(&signIn)

	token := helpers.GenerateToken(signIn.Username, signIn.Role)

	fmt.Println("Token Generated is", token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", token)
}
