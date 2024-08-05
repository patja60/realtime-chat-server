package auth

import (
	"errors"

	"github.com/patja60/realtime-chat-server/pkg/auth"
)

type SigninUsecaseDTO struct {
	Token  string
	UserID string
}

type AuthUsecase interface {
	Signup(email, password string) error
	Signin(email, password string) (string, error)
	// Signout(userID string) error
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

	return u.authRepo.CreateUser(email, hashedPassword)
}

func (u *authUsecaseImpl) Signin(email, password string) (string, error) {
	// get user from user repo
	user, err := u.authRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	// compare password
	isMatch := auth.CompareHashAndPassword(password, user.PasswordHash)
	if !isMatch {
		return "", errors.New("password not match")
	}

	// create token
	token, err := auth.GenerateJWTToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
