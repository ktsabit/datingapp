package repositories

import (
	"context"
	"datingapp/internal/models"
)

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
	EmailExist(ctx context.Context, email string) bool
}

type SwipeRepositoryInterface interface {
	CreateSwipe(
		ctx context.Context,
		userID uint,
		swipedUserID uint,
		swipedDirection models.SwipeDirectionEnum,
	) (*models.Swipe, error)
	CheckReverseSwipe(
		ctx context.Context,
		swipeID uint,
		swipeUserID uint,
		swipeDirection models.SwipeDirectionEnum,
	) (*models.Swipe, error)
	SwipeMatch(ctx context.Context, swiperID uint, swipedID uint) (*models.Swipe, error)
	IncrementDailySwipe(ctx context.Context, userID uint) (int64, error)
	//GetDailySwipe(ctx context.Context, userID uint) (int64, error)
}
