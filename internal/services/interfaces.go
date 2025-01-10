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

type SwipeServiceInterface interface {
	CreateSwipe(
		ctx context.Context,
		userID uint,
		swipedID uint,
		swipeDirection models.SwipeDirectionEnum,
	) (*models.Swipe, error)
}

type ProfileServiceInterface interface {
	GenerateFeed(ctx context.Context, userID uint) ([]*models.User, error)
}
