package routes

import (
	"../controllers"
	"../helpers"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

//InitUser ... instance
func InitUser(R *httprouter.Router, session *mgo.Session) {
	User := controllers.NewUserController(session)
	R.POST("/user", User.CreateUser)
	R.POST("/user/signin", User.Signin)
	R.POST("/user/facebook/signin", User.SigninWithFB)
	R.GET("/user/info/:id", helpers.Authenticate(User.GetUser))
	R.HEAD("/user/:email", User.IsUserExistByEmail)

}
