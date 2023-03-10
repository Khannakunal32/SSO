package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	structures "github.com/khannakunal32/sso/src"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	var Credentials structures.Credentials

	// decoding the json passed by user into credentials
	err := json.NewDecoder(r.Body).Decode(&Credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// fetching the password from db or map
	expectedPassword, ok := structures.Users[Credentials.Username]

	if !ok || expectedPassword != Credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// expiration time for token
	expirationTime := time.Now().Add(time.Minute * 120)

	// creating claims and expiration of token
	claims := &structures.Claims{
		Username: Credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// jwt token generation with algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(structures.JwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error while token string")
		return
	}

	// setting cookie in browser for automatic signin
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}
