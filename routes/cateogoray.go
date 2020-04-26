package routes

import (
	"../controllers"
	"../models"
)

//Cateogory ... instance
var Cateogory = controllers.NewCatController(models.GetSession())

func init() {
	R.POST("/cateogory", Cateogory.CreateCat)

}
