package auth

import auth "github.com/patja60/realtime-chat-server/internal/app/repository"

type Usecase struct {
	repo *auth.Repository
}

func NewAuthUsecase(repo *auth.Repository) *Usecase {
	return &Usecase{repo: repo}
}
