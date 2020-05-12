package helpers

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

//Authenticate ... Middleware for request authentication
func Authenticate(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokn := r.Header.Get("Authorization")
		tokn = strings.Replace(tokn, "bearer ", "", 1)

		isValid := VerifyToken(tokn)
		if isValid != true {
			w.WriteHeader(404)
			w.Write([]byte("Not Authorized"))
			return
		}

		// find user if founct go to next if not return back

		next(w, r, ps)

	}
}
