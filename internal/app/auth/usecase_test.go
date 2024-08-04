package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/patja60/realtime-chat-server/pkg/auth"
)

type mockAuthRepository struct {
	users map[string]*User
}

func (m *mockAuthRepository) CreateUser(email, passwordHash string) error {
	if _, exists := m.users[email]; exists {
		return errors.New("user already exists")
	}
	m.users[email] = &User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	return nil
}

func (m *mockAuthRepository) GetUserByEmail(email string) (*User, error) {
	user, exists := m.users[email]
	if !exists {
		return nil, nil
	}

	return user, nil
}

func TestAuthUsecase_Signup(t *testing.T) {
	t.Run("SignupSuccess", func(t *testing.T) {
		repo := &mockAuthRepository{users: make(map[string]*User)}
		usecase := NewAuthUsecase(repo)

		t.Run("Should have 1 user after signup", func(t *testing.T) {
			email := "user@example.com"
			err := usecase.Signup(email, "password123")
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			if repo.users[email] == nil {
				t.Fatalf("Expected user to be created, got error: %v", err)
			}

			if len(repo.users) != 1 {
				t.Fatalf("Expected to have 1 user to be created, got error: %v", err)
			}
		})

		t.Run("DuplicateSignup", func(t *testing.T) {
			email := "user@example.com"
			_ = usecase.Signup(email, "password123")
			err := usecase.Signup(email, "password123")
			if err == nil {
				t.Fatalf("Expected error for duplicate signup, got nil")
			}

			if len(repo.users) != 1 {
				t.Fatalf("Expected to have 1 user to be created, got error: %v", err)
			}
		})

	})
}

func TestAuthUsecase_Signin(t *testing.T) {
	t.Run("ValidSignin", func(t *testing.T) {
		repo := &mockAuthRepository{users: make(map[string]*User)}
		usecase := NewAuthUsecase(repo)

		email := "user@example.com"
		password := "password123"
		err := usecase.Signup(email, password)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		tokenString, err := usecase.Signin(email, password)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		claims := &auth.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})

		t.Run("Should get correct token string", func(t *testing.T) {

			if !token.Valid {
				t.Fatalf("Expected no error, got %v", err)
			}

			if claims.UserID != "" {
				t.Fatalf("Expected userID to be not empty")
			}

			if claims.ExpiresAt.Unix() <= time.Now().Unix() {
				t.Fatalf("Expected token to have a future expiration time")
			}
		})
	})
}
