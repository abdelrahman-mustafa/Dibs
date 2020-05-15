package routes

import (
	"../controllers"
	"../helpers"
)

//Box ... instance
var Box = controllers.NewBoxController()

func init() {
	R.POST("/admin/box", helpers.Authenticate(Box.CreateBox))
	R.PUT("/admin/box/:id", helpers.Authenticate(Box.UpdateBox))
	R.GET("/admin/boxes", helpers.Authenticate(Box.GetBoxes))
	R.GET("/admin/box/:id", helpers.Authenticate(Box.GetBox))
	R.DELETE("/admin/box/:id", helpers.Authenticate(Box.DeleteBox))

	// user
	R.GET("/user/box/:id", Box.GetBox)

}
