package helpers

import (
	"encoding/json"
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

type response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

//SendResponse ...
func (Res *ResController) SendResponse(message string, code int) {
	res := response{message, code}
	output, _ := json.Marshal(res)
	Res.Res.Header().Set("Content-Type", "application/json")
	Res.Res.WriteHeader(code)
	fmt.Fprintf(Res.Res, "%s", output)
}
