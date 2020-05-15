package routes

import (
	"../controllers"
	"../helpers"
)

//User ... instance

func init() {
	User := controllers.NewUserController()
	R.POST("/user", User.CreateUser)
	R.POST("/user/signin", User.Signin)
	R.GET("/user/info/:id", helpers.Authenticate(User.GetUser))
	R.HEAD("/user/:email", User.IsUserExistByEmail)
}
