package handlers

import "net/http"

type AuthHandlerInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}

type UserHandlerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	//UploadProfilePicture(w http.ResponseWriter, r *http.Request)
}

type SwipeHandlerInterface interface {
	HandleSwipe(w http.ResponseWriter, r *http.Request)
}
