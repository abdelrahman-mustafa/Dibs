package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	//Order ...
	Order struct {
		ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Box       bson.ObjectId `json:"box" bson:"box,omitempty"`
		Status    string        `json:"status" bson:"status,omitempty"`
		Price     int           `json:"price" bson:"price,omitempty"`
		PaymentID int           `json:"paymentId" bson:"paymentId"`
		Count     int           `json:"count" bson:"count,omitempty"`
		Discount  int           `json:"discount" bson:"discount,omitempty"`
	}
)

//OrderID   string        `json:"orderId" bson:"orderId"`
