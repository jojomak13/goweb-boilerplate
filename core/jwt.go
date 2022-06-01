package core

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewJwtToken(data interface{}) string {

	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("APP_KEY")))

	if err != nil {
		log.Fatalf("Cannot generate JWT token, [%v]", err)
	}

	return t
}
