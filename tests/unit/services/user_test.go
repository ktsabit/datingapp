package services_test

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/services"
	"datingapp/tests/mocks"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Register(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := services.NewUserService(mockRepo)

	ctx := context.Background()
	req := models.RegisterRequest{
		Email:    "test@example.com",
		Password: "password123",
		Name:     "Test User",
	}

	mockRepo.On("EmailExist", ctx, req.Email).Return(false)
	mockRepo.On("CreateUser", ctx, mock.AnythingOfType("*models.User")).Return(nil)

	user, err := userService.Register(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, req.Email, user.Email)
	assert.Equal(t, req.Name, user.Name)
	mockRepo.AssertExpectations(t)

	// email exists
	mockRepo = new(mocks.MockUserRepository)
	userService = services.NewUserService(mockRepo)

	mockRepo.On("EmailExist", ctx, req.Email).Return(true)

	_, err = userService.Register(ctx, req)
	assert.Error(t, err)
	assert.Equal(t, "email exist", err.Error())
	mockRepo.AssertCalled(t, "EmailExist", ctx, req.Email)
	mockRepo.AssertNotCalled(t, "CreateUser", ctx, mock.AnythingOfType("*models.User"))
}
