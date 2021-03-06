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
	R.PUT("/user/info/:id", helpers.Authenticate(User.UpdateUser))
	R.POST("/user/signin", User.Signin)
	R.POST("/user/facebook/signin", User.SigninWithFB)
	R.GET("/user/info/:id", helpers.Authenticate(User.GetUser))
	R.GET("/user/favorites/:id", helpers.Authenticate(User.GetUserFavorites))
	R.PUT("/user/favorites/add/:userID/:id", helpers.Authenticate(User.AddUserFavorite))
	R.GET("/user/orders/:id", helpers.Authenticate(User.GetUserOrders))
	R.DELETE("/user/favorites/remove/:userID/:id", helpers.Authenticate(User.DeleteUserFavorite))
	R.HEAD("/user/:email", User.IsUserExistByEmail)

}
