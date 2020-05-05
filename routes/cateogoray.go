package routes

import (
	"../controllers"
	"../models"
)

//Cateogory ... instance
var Cateogory = controllers.NewCatController(models.GetSession())

func init() {
	// admin routes
	R.POST("/admin/cateogory", Cateogory.CreateCateogory)
	R.PUT("/admin/cateogory/:id", Cateogory.UpdateCateogory)
	R.GET("/admin/cateogories", Cateogory.GetCateogories)
	R.GET("/admin/cateogory/:id", Cateogory.GetCateogory)
	R.DELETE("/admin/cateogory/:id", Cateogory.DeleteCateogory)

	// user routes
	//R.GET("/user/cateogory/", Cateogory.GetCateogory)

}
