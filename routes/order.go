package routes

import (
	"../controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

//InitOrder ... instance
func InitOrder(R *httprouter.Router, session *mgo.Session) {
	Order := controllers.NewOrderController(session)
	R.POST("/order", Order.CreateOrder)
}
