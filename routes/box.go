package routes

import (
	"../controllers"
	"../helpers"
	"../models"
)

//Box ... instance
var Box = controllers.NewBoxController(models.GetSession())

func init() {
	R.POST("/admin/box", helpers.Authenticate(Box.CreateBox))
	R.PUT("/admin/box/:id", helpers.Authenticate(Box.UpdateBox))
	R.GET("/admin/boxes", helpers.Authenticate(Box.GetBoxes))

	// user
	R.GET("/user/box/:id", Box.GetBox)

}
