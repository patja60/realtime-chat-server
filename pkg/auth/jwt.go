package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("realtime-chat-secret")

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(userID string) (string, error) {
	claims := &Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtKey)
}
