package routes

import (
	"../controllers"
	"../helpers"
	"../models"
)

//User ... instance
var User = controllers.NewUserController(models.GetSession())

func init() {
	R.POST("/user", User.CreateUser)
	R.POST("/user/signin", User.Signin)
	R.GET("/user/info/:id", helpers.Authenticate(User.GetUser))
	R.HEAD("/user/:email", User.IsUserExistByEmail)

}
