package routes

import (
	"../controllers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

//InitAdmin ... instance
func InitAdmin(R *httprouter.Router, session *mgo.Session) {
	Admin := controllers.NewAdminController(session)
	R.POST("/admin", Admin.CreateAdmin)
	R.POST("/admin/signin", Admin.Signin)
}
