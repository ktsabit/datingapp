package services_test

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/services"
	"datingapp/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
)

func TestSwipeService_CreateSwipe(t *testing.T) {
	mockRepo := new(mocks.MockSwipeRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	service := services.NewSwipeService(mockRepo, mockUserRepo)
	ctx := context.Background()

	successfulSwipe := &models.Swipe{
		UserID:         1,
		TargetID:       2,
		SwipeDirection: models.SwipeRight,
	}

	t.Run("Non-Premium User Within Limit", func(t *testing.T) {
		mockUserRepo.On("GetUserById", uint(1)).Return(&models.User{
			Model:     gorm.Model{ID: 1},
			IsPremium: false,
		}, nil)

		mockRepo.On("IncrementDailySwipe", ctx, uint(1)).Return(int64(5), nil)

		mockRepo.On("CreateSwipe", ctx, uint(1), uint(2), models.SwipeRight).
			Return(successfulSwipe, nil)
		mockRepo.On("CheckReverseSwipe", ctx, uint(1), uint(2), models.SwipeRight).
			Return(&models.Swipe{}, nil)
		mockRepo.On("SwipeMatch", ctx, uint(1), uint(2)).
			Return(successfulSwipe, nil)

		swipe, err := service.CreateSwipe(ctx, 1, 2, models.SwipeRight)
		assert.NoError(t, err)
		assert.NotNil(t, swipe)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertExpectations(t)
		mockUserRepo.AssertExpectations(t)
	})

	//t.Run("Non-Premium User Exceeding Limit", func(t *testing.T) {
	//	mockUserRepo.On("GetUserById", uint(1)).Return(&models.User{
	//		Model:     gorm.Model{ID: 1},
	//		IsPremium: false,
	//	}, nil)
	//
	//	mockRepo.On("IncrementDailySwipe", ctx, uint(1)).Return(int64(11), nil)
	//
	//	swipe, err := service.CreateSwipe(ctx, 1, 2, models.SwipeRight)
	//	assert.Error(t, err)
	//	assert.Nil(t, swipe)
	//	assert.Contains(t, err.Error(), "daily swipe limit exceeded")
	//	mockRepo.AssertNotCalled(t, "CreateSwipe")
	//	mockRepo.AssertNotCalled(t, "CheckReverseSwipe")
	//	mockRepo.AssertNotCalled(t, "SwipeMatch")
	//	mockRepo.AssertExpectations(t)
	//	mockUserRepo.AssertExpectations(t)
	//})

	t.Run("Premium User Bypassing Limit", func(t *testing.T) {
		mockUserRepo.On("GetUserById", uint(1)).Return(&models.User{
			Model:     gorm.Model{ID: 1},
			IsPremium: true,
		}, nil)

		mockRepo.On("CreateSwipe", ctx, uint(1), uint(2), models.SwipeRight).
			Return(successfulSwipe, nil)
		mockRepo.On("CheckReverseSwipe", ctx, uint(1), uint(2), models.SwipeRight).
			Return(nil, nil)

		swipe, err := service.CreateSwipe(ctx, 1, 2, models.SwipeRight)
		assert.NoError(t, err)
		assert.NotNil(t, swipe)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "IncrementDailySwipe")
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("User Match", func(t *testing.T) {
		mockUserRepo.On("GetUserById", mock.AnythingOfType("uint")).Return(&models.User{
			Model:     gorm.Model{ID: 1},
			IsPremium: true,
		}, nil)

		matchedSwipe := &models.Swipe{
			UserID:         2,
			TargetID:       1,
			SwipeDirection: models.SwipeRight,
			IsMatch:        true,
		}

		mockRepo.On("CreateSwipe", ctx, mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), models.SwipeRight).
			Return(successfulSwipe, nil)
		mockRepo.On("CheckReverseSwipe", ctx, mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), models.SwipeRight).
			Return(matchedSwipe, nil)
		mockRepo.On("SwipeMatch", ctx, uint(2), uint(1)).
			Return(nil, nil)
		mockRepo.On("SwipeMatch", ctx, uint(1), uint(2)).
			Return(&models.Swipe{UserID: 1, TargetID: 2, IsMatch: true}, nil)

		_, _ = service.CreateSwipe(ctx, 2, 1, models.SwipeRight)
		swipe, err := service.CreateSwipe(ctx, 1, 2, models.SwipeRight)
		assert.NoError(t, err)
		assert.NotNil(t, swipe)
		assert.True(t, swipe.IsMatch)
		mockRepo.AssertExpectations(t)
		mockRepo.AssertNotCalled(t, "IncrementDailySwipe")
		mockUserRepo.AssertExpectations(t)
	})
}
