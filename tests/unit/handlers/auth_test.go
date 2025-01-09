package handlers_test

import (
	"bytes"
	"datingapp/internal/handlers"
	"datingapp/internal/models"
	"datingapp/internal/services"
	"datingapp/tests/mocks"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Register(t *testing.T) {
	mockService := new(mocks.MockUserService)
	handler := handlers.NewUserHandler(mockService)

	registerReq := models.RegisterRequest{
		Email:    "test@test.com",
		Password: "abc123",
		Name:     "John doe",
	}

	mockService.On("Register", mock.Anything, registerReq).Return(models.User{}, nil)

	body, _ := json.Marshal(registerReq)
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler.Register(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
}

func TestAuthHandler_Login(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockJWTService := new(mocks.MockJWTService)
	handler := handlers.NewAuthHandler(mockRepo, mockJWTService)

	loginReq := models.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	hashedPassword, _ := services.GetPasswordHash("password123")
	mockUser := &models.User{
		Email:    loginReq.Email,
		Password: hashedPassword,
	}

	mockRepo.On("GetUserByEmail", loginReq.Email).Return(mockUser, nil)
	mockJWTService.On("GenerateTokenPair", mock.AnythingOfType("uint")).Return(
		&models.TokenResponse{
			AccessToken:  "test-access-token",
			RefreshToken: "test-refresh-token",
			ExpiresIn:    900,
		}, nil)

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
	mockJWTService.AssertExpectations(t)

	// user not found
	mockRepo = new(mocks.MockUserRepository)
	mockJWTService = new(mocks.MockJWTService)
	handler = handlers.NewAuthHandler(mockRepo, mockJWTService)

	mockRepo.On("GetUserByEmail", loginReq.Email).Return(nil, errors.New("user not found"))

	req = httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(body))
	w = httptest.NewRecorder()

	handler.Login(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")
	mockRepo.AssertExpectations(t)
	mockJWTService.AssertNotCalled(t, "GenerateTokenPair", mock.Anything)

}

func TestAuthHandler_RefreshToken(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockJWTService := new(mocks.MockJWTService)
	handler := handlers.NewAuthHandler(mockRepo, mockJWTService)

	refreshReq := models.RefreshRequest{
		RefreshToken: "test-refresh-token",
	}

	//mockJWTService.On("VerifyRefreshToken", "test-refresh-token").Return(uint(1), nil)
	mockJWTService.On("GenerateTokenPair", mock.AnythingOfType("uint")).Return(
		&models.TokenResponse{
			AccessToken:  "test-access-token",
			RefreshToken: "test-refresh-token",
			ExpiresIn:    900,
		}, nil)
	mockRepo.On("GetUserById", mock.AnythingOfType("uint")).Return(
		&models.User{
			Email: "test@example.com",
		}, nil)

	body, _ := json.Marshal(refreshReq)
	req := httptest.NewRequest(http.MethodPost, "/auth/refresh", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Refresh(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
	mockJWTService.AssertExpectations(t)

	// user not found
	mockRepo = new(mocks.MockUserRepository)
	mockJWTService = new(mocks.MockJWTService)
	handler = handlers.NewAuthHandler(mockRepo, mockJWTService)

	mockRepo.On("GetUserById", mock.AnythingOfType("uint")).Return(nil, errors.New("user not found"))

	req = httptest.NewRequest(http.MethodPost, "/auth/refresh", bytes.NewBuffer(body))
	w = httptest.NewRecorder()

	handler.Refresh(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "User not found")
	mockRepo.AssertExpectations(t)
	mockJWTService.AssertNotCalled(t, "GenerateTokenPair", mock.Anything)

}

func TestAuthHandler_RefreshToken_InvalidToken(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockJWTService := new(mocks.MockJWTService)
	handler := handlers.NewAuthHandler(mockRepo, mockJWTService)

	refreshReq := models.RefreshRequest{
		RefreshToken: "invalid-refresh-token",
	}

	//mockJWTService.On("VerifyRefreshToken", "invalid-refresh-token").Return(uint(0), assert.AnError)
	body, _ := json.Marshal(refreshReq)
	req := httptest.NewRequest(http.MethodPost, "/auth/refresh", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Refresh(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	mockRepo.AssertExpectations(t)
	mockJWTService.AssertExpectations(t)
}
