package auth

import database "github.com/patja60/realtime-chat-server/pkg"

type AuthRepository interface {
	CreateUser(email, hashedPassword string) error
	GetUserByEmail(email string) (*User, error)
}

type authRepositoryImpl struct {
	db *database.DB
}

func NewAuthRepository(db *database.DB) AuthRepository {
	return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) CreateUser(email, hashedPassword string) error {
	_, err := r.db.Conn.Exec("INSERT INTO users(email, passwordHash) VALUES ($1, $2)", email, hashedPassword)
	return err
}

func (r *authRepositoryImpl) GetUserByEmail(email string) (*User, error) {
	user := &User{}
	row := r.db.Conn.QueryRow("SELECT id, email, passwordHash FROM user WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}
