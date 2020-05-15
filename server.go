package main

import (
	"log"
	"net/http"

	"./models"
	"./routes"
	"github.com/gorilla/handlers"
)

func main() {

	models.InitDB()
	handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(routes.R)))

}
