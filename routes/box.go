package routes

import (
	"../controllers"
	"../helpers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

//InitBox ... instance
func InitBox(R *httprouter.Router, session *mgo.Session) {
	Box := controllers.NewBoxController(session)

	R.POST("/admin/box", helpers.Authenticate(Box.CreateBox))
	R.POST("/box/signin", helpers.Authenticate(Box.SignIn))
	R.PUT("/admin/box/:id", helpers.Authenticate(Box.UpdateBox))
	R.PUT("/admin/box/:id/activation", helpers.Authenticate(Box.UpdateActivationBox))
	R.GET("/admin/boxes", helpers.Authenticate(Box.GetBoxes))
	R.GET("/admin/box/:id", helpers.Authenticate(Box.GetBox))
	R.DELETE("/admin/box/:id", helpers.Authenticate(Box.DeleteBox))

	// user
	R.GET("/user/box/:id", helpers.Authenticate(Box.GetBox))
	R.GET("/user/boxes/cateogory/:id", helpers.Authenticate(Box.GetBoxesByCateogory))
	R.GET("/user/boxes", helpers.Authenticate(Box.GetBoxesByUser))
	R.POST("/user/boxes/search", helpers.Authenticate(Box.SearchBoxes))

}
