package main

import (
	"net/http"

	"./routes"
)

func main() {
	http.ListenAndServe(":3000", routes.R)
}
