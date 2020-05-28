package routes

import (
	"../controllers"
	"../helpers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

//InitCat ... instance
func InitCat(R *httprouter.Router, session *mgo.Session) {
	// admin routes
	Cateogory := controllers.NewCatController(session)

	R.POST("/admin/cateogory", helpers.Authenticate(Cateogory.CreateCateogory))
	R.PUT("/admin/cateogory/:id", helpers.Authenticate(Cateogory.UpdateCateogory))
	R.PUT("/admin/cateogory/:id/priority", helpers.Authenticate(Cateogory.UpdateCateogoryPriority))
	R.GET("/admin/cateogories", helpers.Authenticate(Cateogory.GetCateogories))
	R.GET("/admin/cateogory/:id", helpers.Authenticate(Cateogory.GetCateogory))
	R.DELETE("/admin/cateogory/:id", helpers.Authenticate(Cateogory.DeleteCateogory))

	// user routes
	R.GET("/user/cateogories", Cateogory.GetCateogories)
	// R.GET("/user/cateogory/:id", Cateogory.GetCateogory)
	//R.GET("/user/cateogory/", Cateogory.GetCateogory)

}
