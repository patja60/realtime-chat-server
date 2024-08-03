package auth

import (
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	authUsecase AuthUsecase
}

func NewHandler(authUsecase AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.authUsecase.Signup(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
