package routes

import (
	"datingapp/internal/handlers"
	"datingapp/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"net/http"
)

func SetupRoutes(
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
	jwtService *services.JWTService,
) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		tokenAuth := jwtService.TokenAuth()

		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
			userHandler.Register(w, r)
		})
	})

	r.Group(func(r chi.Router) {
		r.Post("/auth/signup", userHandler.Register)
		r.Post("/auth/login", authHandler.Login)
		r.Post("/auth/refresh", authHandler.Refresh)
	})

	return r
}
