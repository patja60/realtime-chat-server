package auth

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
	}{
		{"NonEmptyPassword", "mysecretpassword"},
		{"AnotherPassword", "anotherpassword"},
		{"NumberPassword", "123456"},
		{"SymbolPassword", "asd!@#$asd"},
		{"EmptyPassword", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hashedPassword, err := HashPassword(tc.password)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			t.Run("Hashed password should not be empty", func(t *testing.T) {
				if len(hashedPassword) == 0 {
					t.Fatalf("Expected a hashed password, got an empty string")
				}
			})

			t.Run("Password should match to hashed password", func(t *testing.T) {
				if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(tc.password)); err != nil {
					t.Fatalf("Expected password hash to match")
				}
			})
		})
	}
}

func TestCompareHashAndPassword(t *testing.T) {
	t.Run("CompareHashAndPassword", func(t *testing.T) {
		t.Run("CompareHashAndPasswordSuccess", func(t *testing.T) {
			testCases := []struct {
				name     string
				password string
			}{
				{"NonEmptyPassword", "mysecretpassword"},
				{"AnotherPassword", "anotherpassword"},
				{"NumberPassword", "123456"},
				{"SymbolPassword", "asd!@#$asd"},
				{"EmptyPassword", ""},
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					hashedPassword, err := HashPassword(tc.password)
					if err != nil {
						t.Fatalf("Expected no error, got %v", err)
					}

					isMatch := CompareHashAndPassword(tc.password, hashedPassword)

					t.Run(fmt.Sprintf("Should return match for case %s", tc.name), func(t *testing.T) {
						if !isMatch {
							t.Fatalf("Expected password to match")
						}
					})

				})
			}
		})
		t.Run("CompareHashAndPasswordFailed", func(t *testing.T) {
			testCases := []struct {
				name     string
				password string
			}{
				{"NonEmptyPassword", "mysecretpassword"},
				{"AnotherPassword", "anotherpassword"},
				{"NumberPassword", "123456"},
				{"SymbolPassword", "asd!@#$asd"},
				{"EmptyPassword", ""},
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					notMatchHashedPassword, err := HashPassword(tc.password + "wrongPassMatch")
					if err != nil {
						t.Fatalf("Expected no error, got %v", err)
					}

					isMatch := CompareHashAndPassword(tc.password, notMatchHashedPassword)

					t.Run(fmt.Sprintf("Should return not match for case %s", tc.name), func(t *testing.T) {
						if isMatch {
							t.Fatalf("Expected password to match")
						}
					})

				})
			}
		})
	})
}
