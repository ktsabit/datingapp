package services

import (
	"context"
	"datingapp/internal/models"
	"github.com/go-chi/jwtauth"
)

type JWTServiceInterface interface {
	TokenAuth() *jwtauth.JWTAuth
	GenerateTokenPair(userID uint) (*models.TokenResponse, error)
	VerifyRefreshToken(tokenString string) (uint, error)
}

type UserServiceInterface interface {
	Register(ctx context.Context, userReq models.RegisterRequest) (models.User, error)
}
