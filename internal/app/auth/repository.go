package auth

import database "github.com/patja60/realtime-chat-server/pkg"

type AuthRepository interface {
	Signup(email, hashedPassword string) error
}

type authRepositoryImpl struct {
	db *database.DB
}

func NewRepository(db *database.DB) AuthRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) Signup(email, hashedPassword string) error {
	_, err := r.db.Conn.Exec("INSERT INTO users(email, passwordHash) VALUES ($1, $2)", email, hashedPassword)
	return err
}
