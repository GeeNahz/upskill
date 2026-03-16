package helpers

import (
	"http/library/src/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CompareHashAndPassword(hashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func GenerateToken(user *models.User) (string, error) {
	secret := []byte("super-secret-key")
	method := jwt.SigningMethodHS256
	claims := jwt.MapClaims{
		"userId":   user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 168).Unix(), // 7 days
	}

	token, err := jwt.NewWithClaims(method, claims).SignedString(secret)
	if err != nil {
		return "", err
	}
	return token, nil
}
