package routes

import (
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

//InitRouter ... Instantiate a new router
func InitRouter(db *mgo.Session) *httprouter.Router {
	R := httprouter.New()
	InitBox(R, db)
	InitAdmin(R, db)
	InitCat(R, db)
	InitUser(R, db)

	return R
}
