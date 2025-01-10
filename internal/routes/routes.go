package routes

import (
	"context"
	"datingapp/internal/handlers"
	"datingapp/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func UserIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", uint(userIDFloat))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SetupRoutes(
	userHandler handlers.UserHandlerInterface,
	authHandler handlers.AuthHandlerInterface,
	jwtService services.JWTServiceInterface,
	swipeHandler handlers.SwipeHandlerInterface,
) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		tokenAuth := jwtService.TokenAuth()
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(UserIDMiddleware)

		r.Post("/swipe", swipeHandler.HandleSwipe)
	})

	r.Group(func(r chi.Router) {
		r.Post("/auth/signup", userHandler.Register)
		r.Post("/auth/login", authHandler.Login)
		r.Post("/auth/refresh", authHandler.Refresh)
	})

	r.Get("/images/{filename}", serveImage)

	return r
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	filePath := "./uploads/profile_pictures/" + filename

	http.ServeFile(w, r, filePath)
}
