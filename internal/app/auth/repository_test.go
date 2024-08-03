package auth

import "testing"

type mockDB struct {
	users map[string]*User
}

func TestAuthRepository_Signup(t *testing.T) {
	t.Run("Register", func(t *testing.T) {
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

		t.Run("DuplicateRegistration", func(t *testing.T) {
			email := "user@example.com"
			_ = usecase.Signup(email, "password123")
			err := usecase.Signup(email, "password123")
			if err == nil {
				t.Fatalf("Expected error for duplicate registration, got nil")
			}

			if len(repo.users) != 1 {
				t.Fatalf("Expected to have 1 user to be created, got error: %v", err)
			}
		})

	})
}
