package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtkey = []byte("my-secret-key")

func GenerateToken(UserID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}
