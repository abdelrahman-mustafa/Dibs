package main

import (
	"fmt"
	"log"
	"net/http"

	"./models"
	"./routes"
	"github.com/gorilla/handlers"
	"github.com/robfig/cron"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	R := routes.InitRouter(models.GetSession())
	c := cron.New()
	c.AddFunc("@every 1d", updateAllBoxes)
	c.Start()
	handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(R)))

}

func updateAllBoxes() {
	fmt.Println("i was call")
	session := models.GetSession()
	var boxes []models.Box
	session.DB("dibs").C("boxes").Find(bson.M{}).All(&boxes)
	for _, box := range boxes {
		if box.TomorrowBoxes != 0 {
			box.AvailableBoxes = box.TomorrowBoxes
			box.TomorrowBoxes = box.MinBoxes
		} else {
			box.AvailableBoxes = box.MinBoxes
		}
		out := bson.M{"$set": box}
		session.DB("dibs").C("boxes").UpdateId(box.ID, out)
	}
	session.Close()
}
