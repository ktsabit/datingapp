package handlers

import "net/http"

type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

type UserHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
}
