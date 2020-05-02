package main

import (
	"net/http"

	"./routes"
	"github.com/rs/cors"
)

func main() {

	handler := cors.Default().Handler(routes.R)
	http.ListenAndServe(":3000", handler)
}
