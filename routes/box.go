package routes

import (
	"../controllers"
	"../models"
)

//Box ... instance
var Box = controllers.NewBoxController(models.GetSession())

func init() {
	R.POST("/box", Box.CreateBox)
	R.PUT("/box/:id", Box.UpdateBox)

}
