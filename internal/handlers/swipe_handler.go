package handlers

import (
	"datingapp/internal/models"
	"datingapp/internal/services"
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
)

type SwipeHandler struct {
	Service services.SwipeServiceInterface
}

func NewSwipeHandler(service services.SwipeServiceInterface) *SwipeHandler {
	return &SwipeHandler{Service: service}
}

func (sh *SwipeHandler) HandleSwipe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := r.Context().Value("userID").(uint)
	var req models.SwipeRequest

	if err := render.DecodeJSON(r.Body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	swipe, err := sh.Service.CreateSwipe(ctx, userID, req.SwipedUserID, req.SwipedDirection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(swipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
