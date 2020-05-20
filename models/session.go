package models

import mgo "gopkg.in/mgo.v2"

//GetSession ... return a mongo session
func GetSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session

}
