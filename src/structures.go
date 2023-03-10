package src

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var SecretKey = goDotEnvVariable("secret_key")
var JwtKey = []byte(SecretKey)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
