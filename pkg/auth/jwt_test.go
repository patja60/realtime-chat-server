package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWTToken(t *testing.T) {
	t.Run("GenerateJWTToken", func(t *testing.T) {
		t.Run("GenerateJWTTokenSuccess", func(t *testing.T) {
			userID := "userID"
			tokenString, err := GenerateJWTToken(userID)

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return JwtKey, nil
			})

			if tokenString == "" {
				t.Fatalf("Expected a valid token, got an empty string")
			}

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			t.Run("Token should be valid", func(t *testing.T) {
				if !token.Valid {
					t.Fatalf("Expected a valid token, got an invalid token")
				}
			})

			t.Run("Token should have correct userID", func(t *testing.T) {
				if claims.UserID != userID {
					t.Fatalf("Expected userID to be %v, got %v", userID, claims.UserID)
				}
			})

			t.Run("Token should have future expired time", func(t *testing.T) {
				if claims.ExpiresAt.Unix() <= time.Now().Unix() {
					t.Fatalf("Expected token to have a future expiration time")
				}
			})
		})
	})
}
