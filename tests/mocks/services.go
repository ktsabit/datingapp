package mocks

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/services"
	"errors"
	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/mock"
)

type MockSwipeService struct {
	mock.Mock
}

func (m *MockSwipeService) CreateSwipe(
	ctx context.Context,
	userID uint,
	swipedID uint,
	swipeDirection models.SwipeDirectionEnum,
) (*models.Swipe, error) {
	args := m.Called(ctx, userID, swipedID, swipeDirection)
	return args.Get(0).(*models.Swipe), args.Error(1)
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(ctx context.Context, req models.RegisterRequest) (models.User, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(models.User), args.Error(1)
}

var _ services.JWTServiceInterface = (*MockJWTService)(nil)

type MockJWTService struct {
	mock.Mock
}

func (m *MockJWTService) TokenAuth() *jwtauth.JWTAuth {
	args := m.Called()
	return args.Get(0).(*jwtauth.JWTAuth)
}

func (m *MockJWTService) VerifyRefreshToken(tokenString string) (uint, error) {
	if tokenString == "test-refresh-token" {
		return 1, nil
	}
	return 0, errors.New("invalid token")
}

func (m *MockJWTService) GenerateTokenPair(userID uint) (*models.TokenResponse, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.TokenResponse), args.Error(1)
}
