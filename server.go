package main

import (
	"net/http"

	"./routes"
	"github.com/rs/cors"
)

func main() {

	// routes.R.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Header.Get("Access-Control-Request-Method") != "" {
	// 		// Set CORS headers
	// 		header := w.Header()
	// 		header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
	// 		header.Set("Access-Control-Allow-Origin", "*")
	// 	}

	// 	// Adjust status code to 204
	// 	w.WriteHeader(http.StatusNoContent)
	// })
	handler := cors.Default().Handler(routes.R)
	http.ListenAndServe(":3000", handler)
}
