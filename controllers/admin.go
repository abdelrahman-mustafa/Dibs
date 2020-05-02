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
	Username string `json:"username" bson:"username"`
	Role     string `json:"role" bson:"role"`
	Password string `json:"password" bson:"password"`
}

type (

	// AdminController represents the controller for operating on the admin resource
	AdminController struct {
		session *mgo.Session
	}
)

//SetupResponse ...
func SetupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// NewAdminController ... returns a instance of UserController structure
func NewAdminController(s *mgo.Session) *AdminController {
	return &AdminController{s}
}

// CreateAdmin ... creates a new admin resource
func (ad AdminController) CreateAdmin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// tokn := r.Header.Get("Authorization")

	// _, err := helpers.VerifyToken(tokn)

	// if err != nil {
	// 	w.WriteHeader(404)
	// 	w.Write([]byte("Missing token"))
	// } else {
	SetupResponse(&w, r)
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

	// fmt.Fprintf(w, "%s", uj)
}

// Signin ... sign in as Admin
func (ad AdminController) Signin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Start Sign in")
	signIn := SignIn{}
	json.NewDecoder(r.Body).Decode(&signIn)

	isAdmin, password, id, role := models.IsAdminExist(signIn.Username)
	if isAdmin != true {
		if signIn.Username == "admin" && signIn.Password == "admin" {
			// create user first time

			admin := models.Admin{}
			admin.Username = signIn.Username
			admin.Password = signIn.Password
			admin.Role = "admin"

			//create id
			admin.ID = bson.NewObjectId()

			// write struct of admin to DB
			ad.session.DB("dibs").C("admins").Insert(admin)

			token := helpers.GenerateToken(admin.ID, admin.Role)
			fmt.Println("Token Generated is", token)

			Res := SignInResponse{}
			Res.ID = admin.ID
			Res.Token = token
			output, _ := json.Marshal(Res)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			fmt.Fprintf(w, "%s", output)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			w.Write([]byte("Not Authorized"))
		}
	} else {
		err := helpers.Compare(password, signIn.Password)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			w.Write([]byte("Not Authorized"))
		} else {
			token := helpers.GenerateToken(id, role)
			Res := SignInResponse{}
			Res.ID = id
			Res.Token = token
			output, _ := json.Marshal(Res)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			fmt.Fprintf(w, "%s", output)

		}
	}
}
