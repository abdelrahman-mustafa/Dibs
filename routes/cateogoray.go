package routes

import (
	"../controllers"
	"../helpers"
)

//Cateogory ... instance

func init() {
	// admin routes
	Cateogory := controllers.NewCatController()

	R.POST("/admin/cateogory", helpers.Authenticate(Cateogory.CreateCateogory))
	R.PUT("/admin/cateogory/:id", helpers.Authenticate(Cateogory.UpdateCateogory))
	R.GET("/admin/cateogories", helpers.Authenticate(Cateogory.GetCateogories))
	R.GET("/admin/cateogory/:id", helpers.Authenticate(Cateogory.GetCateogory))
	R.DELETE("/admin/cateogory/:id", helpers.Authenticate(Cateogory.DeleteCateogory))

	// user routes
	R.GET("/user/cateogories", Cateogory.GetCateogories)
	R.GET("/user/cateogory/:id", Cateogory.GetCateogory)
	//R.GET("/user/cateogory/", Cateogory.GetCateogory)

}
