package auth

import "github.com/patja60/realtime-chat-server/pkg/auth"

type AuthUsecase interface {
	Signup(email, password string) error
}

type authUsecaseImpl struct {
	authRepo AuthRepository
}

func NewAuthUsecase(authRepo AuthRepository) AuthUsecase {
	return &authUsecaseImpl{authRepo: authRepo}
}

func (u *authUsecaseImpl) Signup(email, password string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	return u.authRepo.Signup(email, hashedPassword)
}
