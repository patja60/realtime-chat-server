package auth

type AuthUsecase interface {
}

type authUsecaseImpl struct {
	authRepo AuthRepository
}

func NewAuthUsecase(authRepo AuthRepository) AuthUsecase {
	return &authUsecaseImpl{authRepo: authRepo}
}
