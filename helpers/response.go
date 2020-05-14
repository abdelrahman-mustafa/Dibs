package helpers

import (
	"fmt"
	"net/http"
)

//SendBadRequest ...
type (

	// ResController represents the controller for operating on the Cat resource
	ResController struct {
		Res http.ResponseWriter
	}
)

//SendBadRequest ...
func (Res *ResController) SendBadRequest(message string) {
	Res.Res.Header().Set("Content-Type", "application/json")
	Res.Res.WriteHeader(201)
	fmt.Fprintf(Res.Res, "%s", message)
}
