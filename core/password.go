package core

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		log.Fatalf("Canot hash password, [%v]", err)
	}

	return string(b)
}

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
