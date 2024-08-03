package auth

import (
	database "github.com/patja60/realtime-chat-server/pkg"
	"github.com/patja60/realtime-chat-server/pkg/auth"
)

type AuthRepository interface {
	Signup(email, password string) error
}

type authRepositoryImpl struct {
	db *database.DB
}

func NewRepository(db *database.DB) AuthRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) Signup(email, password string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	return r.db.CreateUser(email, hashedPassword)
}
