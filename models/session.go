package models

import mgo "gopkg.in/mgo.v2"

//Session ...
var Session *mgo.Session

//InitDB ... return a mongo session
func InitDB() {
	// Connect to our local mongo
	var err error
	Session, err = mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo runnig?
	if err != nil {
		panic(err)
	}

}
