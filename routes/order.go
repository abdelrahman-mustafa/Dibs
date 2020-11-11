package routes

import (
	"../controllers"
	"../helpers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

//InitOrder ... instance
func InitOrder(R *httprouter.Router, session *mgo.Session) {
	Order := controllers.NewOrderController(session)
	R.POST("/order", helpers.Authenticate(Order.CreateOrder))
	R.GET("/order/:id", helpers.Authenticate(Order.GetOrder))
	R.GET("/order/list", helpers.Authenticate(Order.GetOrders))

}
