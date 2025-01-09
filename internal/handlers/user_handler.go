package handlers

import (
	"datingapp/internal/models"
	"datingapp/internal/services"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req models.RegisterRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.Service.Register(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
