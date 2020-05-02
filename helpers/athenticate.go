package helpers

import (
	"net/http"
	"strings"
)

//Authenticate ... Middleware for request authentication
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokn := r.Header.Get("Authorization")
		tokn = strings.Replace(tokn, "bearer", "", 1)

		_, err := VerifyToken(tokn)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Not Authorized"))
			return
		}

		// find user if founct go to next if not return back

		next.ServeHTTP(w, r)

	}
}
