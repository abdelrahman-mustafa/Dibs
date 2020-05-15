package routes

import (
	"../controllers"
)

//Admin ... instance
var Admin = controllers.NewAdminController()

func init() {
	R.POST("/admin", Admin.CreateAdmin)
	R.POST("/admin/signin", Admin.Signin)

}
