package routes

import (
	"../controllers"
	"../models"
)

//User ... instance
var User = controllers.NewUserController(models.GetSession())

func init() {
	R.POST("/user", User.CreateUser)
	R.POST("/user/signin", User.Signin)
	R.GET("/user/info/:id", User.GetUser)

}
