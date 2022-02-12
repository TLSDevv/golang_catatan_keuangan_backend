package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordToHash(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashPassword)
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
