package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	var bytes, err = bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}