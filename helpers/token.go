package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

//Claims ... represent the object of JWT
type Claims struct {
	Username bson.ObjectId `json:"username"`
	Role     string        `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("dibs_2020")

//GenerateToken ...
func GenerateToken(username bson.ObjectId, role string) string {
	expirationTime := time.Now().Add(300 * 30 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

//Data ... represtedn the return data by token verfication
type Data struct {
	Username bson.ObjectId `json:"username"`
	Role     string        `json:"role"`
}

//VerifyToken ...
func VerifyToken(token string) bool {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false
	}
	if !tkn.Valid {
		return false
	}

	return true
}
