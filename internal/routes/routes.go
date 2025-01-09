package routes

import (
	"datingapp/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(
	userHandler *handlers.UserHandler,
	authHandler *handlers.AuthHandler,
) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	//tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	//r.Group(func(r chi.Router) {
	//	r.Use(jwtauth.Verifier(tokenAuth))
	//	r.Use(jwtauth.Authenticator)
	//	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
	//		userHandler.Register(w, r)
	//	})
	//})

	r.Group(func(r chi.Router) {
		r.Post("/auth/signup", userHandler.Register)
		r.Post("/auth/login", authHandler.Login)
		r.Post("/auth/refresh", authHandler.Refresh)
	})

	return r
}
