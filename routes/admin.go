package routes

import (
	"../controllers"
	"../models"
)

//Admin ... instance
var Admin = controllers.NewAdminController(models.GetSession())

func init() {
	R.POST("/admin", Admin.CreateAdmin)
	R.POST("/admin/signin", Admin.Signin)

}
