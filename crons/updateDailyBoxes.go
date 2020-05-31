package crons

import (
	"../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UpdateAllBoxes ...
func UpdateAllBoxes(session *mgo.Session) {
	var boxes []models.Box
	session.DB("dibs").C("boxes").Find(bson.M{}).All(boxes)
	for _, box := range boxes {
		if box.TomorrowAvailableBoxes != 0 {
			box.AvailableBoxes = box.TomorrowAvailableBoxes
		} else {
			box.AvailableBoxes = 0
		}
		out := bson.M{"$set": box}
		session.DB("dibs").C("boxes").UpdateId(box.ID, out)
	}
}
