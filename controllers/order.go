package controllers

import (
	"encoding/json"
	"fmt"

	"net/http"

	"../helpers"
	"../models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (

	// OrderController represents the controller for operating on the Order resource
	OrderController struct {
		session *mgo.Session
	}
)

// NewOrderController ... returns a instance of UserController structure
func NewOrderController(session *mgo.Session) *OrderController {
	return &OrderController{session}
}

// CreateOrder ... creates a new Order resource
func (ad OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// get user id from header
	userid := p.ByName("userID")
	fmt.Println("Start Get from id  is ", userid)

	user := models.GetUser(userid, ad.session)
	if user.Username == "" {
		res := helpers.ResController{Res: w}
		res.SendResponse("Not Fount", 404)
		return
	}
	Order := models.Order{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Order)

	//create id
	Order.ID = bson.NewObjectId()
	Order.Status = "Done"

	newPay := helpers.Pay{
		Email:     user.Email,
		FirstName: user.Username,
		LastName:  user.Username,
		Phone:     user.Phone,
	}
	newPay.PayAuth()
	newPay.PlaceOrder(Order.Price, 11)
	newPay.GetToken()
	frame := newPay.BuildIFrame()

	Order.PaymentID = newPay.OrderID
	err := ad.session.DB("dibs").C("orders").Insert(Order)
	ad.session.DB("dibs").C("users").UpdateId(userid, bson.M{
		"$addToSet": bson.M{
			"orders": Order,
		},
	})
	res := helpers.ResController{Res: w}

	if err != nil {
		res.SendResponse("Internal Server Error", 504)
	}
	// here apply the payment implementation

	res.SendResponse(frame, 200)
}

// GetOrder ... creates a new Order resource
func (ad OrderController) GetOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// get user id from header
	userid := p.ByName("userID")
	fmt.Println("Start Get from id  is ", userid)

	user := models.GetUser(userid, ad.session)
	if user.Username == "" {
		res := helpers.ResController{Res: w}
		res.SendResponse("Not Fount", 404)
		return
	}
	oid := bson.ObjectIdHex(p.ByName("id"))

	Order := &models.Order{}
	err := ad.session.DB("dibs").C("orders").FindId(oid).One(&Order)
	res := helpers.ResController{Res: w}

	if err != nil {
		res.SendResponse("Internal Server Error", 504)
	}
	// here apply the payment implementation
	output, _ := json.Marshal(Order)

	res.SendResponse(string(output), 200)
}

// //UpdateOrder ... creates a new Order resource
// func (ad OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	// query := r.URL.Query()

// 	// Order := &models.Order{}
// 	// err := ad.session.DB("dibs").C("orders").FindId(oid).One(&Order)
// 	// res := helpers.ResController{Res: w}

// 	// if err != nil {
// 	// 	res.SendResponse("Internal Server Error", 504)
// 	// }
// 	// // here apply the payment implementation
// 	// output, _ := json.Marshal(Order)

// 	// res.SendResponse(string(output), 200)
// }
