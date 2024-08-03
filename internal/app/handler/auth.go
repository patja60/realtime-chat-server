package auth

import (
	auth "github.com/patja60/realtime-chat-server/internal/app/usecase"
	database "github.com/patja60/realtime-chat-server/pkg"
)

type Handler struct {
	db          *database.DB
	authUsecase auth.Usecase
}
