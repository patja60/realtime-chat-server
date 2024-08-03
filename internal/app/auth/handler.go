package auth

import (
	database "github.com/patja60/realtime-chat-server/pkg"
)

type Handler struct {
	db          *database.DB
	authUsecase AuthUsecase
}
