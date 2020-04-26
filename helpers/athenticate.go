package helpers

import (
	"net/http"
)

//Authenticate ... Middleware for request authentication
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokn := r.Header.Get("Authorization")
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
