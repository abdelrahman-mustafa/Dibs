package controllers

import (
	"encoding/json"
	"fmt"
	"time"
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

type paymentResponse struct {
	Iframe  string
	OrderID bson.ObjectId
}

// NewOrderController ... returns a instance of UserController structure
func NewOrderController(session *mgo.Session) *OrderController {
	return &OrderController{session}
}

// CreateOrder ... creates a new Order resource
func (ad OrderController) CreateOrder(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// get user id from header
	userid := r.Header.Get("userID")
	fmt.Println("Start Get from id  is ", userid)
	if userid == "" {
		res := helpers.ResController{Res: w}
		res.SendResponse("Not Authorized", 404)
		return
	}
	user := models.GetUser(userid, ad.session)
	if user.Username == "" {
		res := helpers.ResController{Res: w}
		res.SendResponse("Not Authorized", 404)
		return
	}
	if user.Phone == "" {
		res := helpers.ResController{Res: w}
		res.SendResponse("Please update your phone number", 403)
		return
	}
	Order := models.Order{}
	//prase json  of body and attach to admoin struct
	json.NewDecoder(r.Body).Decode(&Order)

	Box := models.Box{}

	oid := Order.Box

	// write struct of admni to DB
	ad.session.DB("dibs").C("boxes").FindId(oid).One(&Box)

	if Box.Name == "" {
		w.Header().Set("Content-Type", "appliBoxion/json")
		w.WriteHeader(401)
		fmt.Fprintf(w, "%s", "Not valid Box")
		return
	}
	//create id
	Order.ID = bson.NewObjectId()
	Order.Status = "Done"
	if Order.Count == 0 {
		Order.Count = 1
	}
	// if Box.AvailableBoxes < Order.Count {
	// 	w.Header().Set("Content-Type", "appliBoxion/json")
	// 	w.WriteHeader(401)
	// 	fmt.Fprintf(w, "%s", "Not valid Box Count")
	// 	return
	// }

	Order.Price = Box.Price * Order.Count * 100

	newPay := helpers.Pay{
		Email:     user.Email,
		FirstName: user.Username,
		LastName:  user.Username,
		Phone:     user.Phone,
	}
	newPay.PayAuth()
	newPay.PlaceOrder(80, 11)
	newPay.GetToken()
	frame := newPay.BuildIFrame()
	// update box
	Box.AvailableBoxes = Box.AvailableBoxes - Order.Count
	out := bson.M{"$set": Box}
	// write struct of admni to DB
	ad.session.DB("dibs").C("boxes").UpdateId(oid, out)
	Order.PaymentID = newPay.OrderID
	Order.CreatedAt =  time.Now().String()
	Order.BoxDetails =  Box
	err := ad.session.DB("dibs").C("orders").Insert(Order)
		uid := bson.ObjectIdHex(userid)
	err = ad.session.DB("dibs").C("users").UpdateId(uid, bson.M{
		"$addToSet": bson.M{
			"orders": Order.ID,
		},
	})
		if err != nil {
		println("Order Error: ", err.Error())
	}
	res := helpers.ResController{Res: w}
	pRes := paymentResponse{
		frame,
		Order.ID,
	}
	output, _ := json.Marshal(&pRes)

	println("OUTPUT: ", string(output))
	if err != nil {
		println("Error : ", err.Error())
		res.SendResponse("Internal Server Error", 504)
		return
	}
	// here apply the payment implementation
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}

// GetOrders ... creates a new Order resource
func (ad OrderController) GetOrders(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// get user id from header
	userid := p.ByName("userID")
	fmt.Println("Start Get from id  is ", userid)

	user := models.GetUser(userid, ad.session)
	if user.Username == "" {
		res := helpers.ResController{Res: w}
		res.SendResponse("Not Fount", 404)
		return
	}

	var results []bson.M
	err := ad.session.DB("dibs").C("orders").Find(bson.M{}).All(&results)
	res := helpers.ResController{Res: w}

	if err != nil {
		res.SendResponse("Internal Server Error", 504)
	}
	// here apply the payment implementation
	output, _ := json.Marshal(results)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", output)
}
