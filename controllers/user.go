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

// SignInResponse ... sign in as User
type SignInResponse struct {
	Token string        `json:"token" bson:"token"`
	ID    bson.ObjectId `json:"id" bson:"id"`
}

// SignInAsUser ... sign in as User
type SignInAsUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// SignInAsUserWithFB ... sign in as User
type SignInAsUserWithFB struct {
	FacebookID string `json:"facebookID" bson:"facebookID"`
}

type (

	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController ... returns a instance of UserController structure
func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

// CreateUser ... creates a new User resource
func (ad UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	User := models.User{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&User)

	//create id
	User.ID = bson.NewObjectId()

	fmt.Println(User.Username)
	fmt.Println(User.Password)

	encryptedPassword, er := helpers.Encrypt(User.Password)

	if er != nil {
		res := helpers.ResController{Res: w}
		res.SendResponse("Something goes wrong", 500)
		return
	}

	User.Password = encryptedPassword
	// write struct of admni to DB
	ad.session.DB("dibs").C("users").Insert(User)

	// build response for user
	token := helpers.GenerateToken(User.ID, "user")
	Res := SignInResponse{}
	Res.ID = User.ID
	Res.Token = token
	//

	// convert struct to JSON
	output, _ := json.Marshal(Res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)

}

//UpdateUser ... creates a new User resource
func (ad UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	User := models.User{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&User)

	//create id
	User.ID = bson.NewObjectId()

	if User.Password != "" {
		encryptedPassword, er := helpers.Encrypt(User.Password)
		if er != nil {
			res := helpers.ResController{Res: w}
			res.SendResponse("Something goes wrong", 500)
			return
		}
		User.Password = encryptedPassword

	}
	oid := bson.ObjectIdHex(p.ByName("id"))
	out := bson.M{"$set": User}

	// write struct of admni to DB
	ad.session.DB("dibs").C("users").UpdateId(oid, out)

	res := helpers.ResController{Res: w}
	res.SendResponse("The User is updated successfully", 200)

}

// Signin ... sign in as User
func (ad UserController) Signin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Println("Start Sign in")
	signIn := SignInAsUser{}
	json.NewDecoder(r.Body).Decode(&signIn)

	//verify username
	isValid, userPassword, userID := models.IsUserExist(signIn.Email, ad.session)
	if isValid == false {
		res := helpers.ResController{Res: w}

		res.SendResponse("Not valid Email", 401)
		return
	}

	err := helpers.Compare(userPassword, signIn.Password)
	if err != nil {
		res := helpers.ResController{Res: w}

		res.SendResponse("Not valid password", 401)
		return
	}

	token := helpers.GenerateToken(userID, "user")
	Res := SignInResponse{}
	Res.ID = userID
	Res.Token = token
	output, _ := json.Marshal(Res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)
}

// SigninWithFB ... sign in as User
func (ad UserController) SigninWithFB(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Println("Start Sign in with facebook")
	signIn := SignInAsUserWithFB{}
	json.NewDecoder(r.Body).Decode(&signIn)

	//verify username
	user := models.GetUserByFaceBookID(signIn.FacebookID, ad.session)
	if user.Name == "" {
		res := helpers.ResController{Res: w}

		res.SendResponse("Not valid Facebook Account", 401)
		return
	}

	token := helpers.GenerateToken(user.ID, "user")
	Res := SignInResponse{}
	Res.ID = user.ID
	Res.Token = token
	output, _ := json.Marshal(Res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", output)
}

// GetUser ... GetUser data
func (ad UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Println("Start Get user data")

	// get user id from header
	id := p.ByName("id")
	fmt.Println("Start Get from id  is ", id)

	user := models.GetUser(id, ad.session)
	if user.Username == "" {
		res := helpers.ResController{Res: w}
		res.SendResponse("Not Fount", 404)
		return
	}
	output, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}

// IsUserExistByEmail ... IsUserExistByEmail data
func (ad UserController) IsUserExistByEmail(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Println("Start Get user data")

	// get user id from header
	email := p.ByName("email")
	fmt.Println("Start Get from email  is ", email)

	user := models.GetUserByEmail(email, ad.session)
	res := helpers.ResController{Res: w}
	if user == false {
		res.SendResponse("Not Fount", 404)
		return
	}

	res.SendResponse("Found", 200)

}
