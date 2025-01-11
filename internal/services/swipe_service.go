package services

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
	"errors"
)

type SwipeService struct {
	SwipeRepo  repositories.SwipeRepositoryInterface
	UserRepo   repositories.UserRepositoryInterface
	SwipeLimit int
}

func NewSwipeService(
	s repositories.SwipeRepositoryInterface,
	u repositories.UserRepositoryInterface) *SwipeService {
	return &SwipeService{
		SwipeRepo:  s,
		UserRepo:   u,
		SwipeLimit: 10,
	}
}

func (s *SwipeService) CreateSwipe(
	ctx context.Context,
	userID uint,
	swipedID uint,
	swipeDirection models.SwipeDirectionEnum,
) (*models.Swipe, error) {
	user, err := s.UserRepo.GetUserById(userID)
	if err != nil {
		return nil, err
	}

	if !user.IsPremium {
		count, err := s.SwipeRepo.IncrementDailySwipe(ctx, userID)
		if err != nil {
			return nil, err
		}

		if count > 10 {
			return nil, errors.New("daily swipe limit exceeded")
		}
	}

	_, err = s.SwipeRepo.CreateSwipe(ctx, userID, swipedID, swipeDirection)
	if err != nil {
		return nil, err
	}

	swipe, err := s.SwipeRepo.CheckReverseSwipe(ctx, userID, swipedID, swipeDirection)
	if err != nil {
		return nil, err
	}

	if swipe == nil {
		return nil, nil
	}

	swipe, err = s.SwipeRepo.SwipeMatch(ctx, userID, swipedID)
	if err != nil {
		return nil, err
	}
	return swipe, nil
}
