package mocks

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
	"github.com/stretchr/testify/mock"
)

var _ repositories.UserRepositoryInterface = (*MockUserRepository)(nil)

type MockSwipeRepository struct {
	mock.Mock
}

//func (m *MockSwipeRepository) GetDailySwipe(ctx context.Context, userID uint) (int64, error) {
//	args := m.Called(ctx, userID)
//	return args.Get(0).(int64), args.Error(1)
//}

func (m *MockSwipeRepository) IncrementDailySwipe(ctx context.Context, userID uint) (int64, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockSwipeRepository) CreateSwipe(
	ctx context.Context,
	userID uint,
	swipedUserID uint,
	swipedDirection models.SwipeDirectionEnum,
) (*models.Swipe, error) {
	args := m.Called(ctx, userID, swipedUserID, swipedDirection)
	return args.Get(0).(*models.Swipe), args.Error(1)
}

func (m *MockSwipeRepository) SwipeMatch(ctx context.Context, swiperID uint, swipedID uint) (*models.Swipe, error) {
	args := m.Called(ctx, swiperID, swipedID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Swipe), args.Error(1)
}

func (m *MockSwipeRepository) CheckReverseSwipe(
	ctx context.Context,
	userID uint,
	targetUserID uint,
	swipedDirection models.SwipeDirectionEnum,
) (*models.Swipe, error) {
	args := m.Called(ctx, userID, targetUserID, swipedDirection)
	return args.Get(0).(*models.Swipe), args.Error(1)
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) EmailExist(ctx context.Context, email string) bool {
	args := m.Called(ctx, email)
	return args.Bool(0)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserById(id uint) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}
