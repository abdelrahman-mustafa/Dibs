package routes

import (
	"../controllers"
	"../models"
)

//Box ... instance
var Box = controllers.NewBoxController(models.GetSession())

func init() {
	R.POST("/admin/box", Box.CreateBox)
	R.PUT("/admin/box/:id", Box.UpdateBox)
	R.GET("/admin/boxes", Box.GetBoxes)

	// user
	R.GET("/user/box/:id", Box.GetBox)

}
