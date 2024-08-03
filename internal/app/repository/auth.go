package auth

import (
	database "github.com/patja60/realtime-chat-server/pkg"
	"github.com/patja60/realtime-chat-server/pkg/auth"
)

type Repository struct {
	db *database.DB
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Signup(email, password string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	return r.db.CreateUser(email, hashedPassword)
}
