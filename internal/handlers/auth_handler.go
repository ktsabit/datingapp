package handlers

import (
	"datingapp/internal/configs"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type AuthHandler struct {
	userRepo   *repositories.UserRepository
	jwtService *configs.JWTService
}

func NewAuthHandler(userRepo *repositories.UserRepository, jwtService *configs.JWTService) *AuthHandler {
	return &AuthHandler{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	tokens, err := h.jwtService.GenerateTokenPair(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		log.Fatalf("Failed to encode tokens: %v", err)
		return
	}
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req models.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := h.jwtService.VerifyRefreshToken(req.RefreshToken)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	if _, err = h.userRepo.GetUserById(userID); err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	tokens, err := h.jwtService.GenerateTokenPair(userID)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tokens)
	if err != nil {
		log.Fatalf("Failed to encode tokens: %v", err)
		return
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {}
