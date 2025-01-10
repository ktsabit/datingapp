package services

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
)

type SwipeService struct {
	Repository repositories.SwipeRepositoryInterface
}

func NewSwipeService(r repositories.SwipeRepositoryInterface) *SwipeService {
	return &SwipeService{Repository: r}
}

func (s *SwipeService) CreateSwipe(
	ctx context.Context,
	userID uint,
	swipedID uint,
	swipeDirection models.SwipeDirectionEnum,
) (*models.Swipe, error) {
	_, err := s.Repository.CreateSwipe(ctx, userID, swipedID, swipeDirection)
	if err != nil {
		return nil, err
	}

	swipe, err := s.Repository.CheckReverseSwipe(ctx, userID, swipedID, swipeDirection)
	if err != nil {
		return nil, err
	}

	if swipe == nil {
		return nil, nil
	}

	_, err = s.Repository.SwipeMatch(ctx, userID, swipedID)
	if err != nil {
		return nil, err
	}
	return swipe, nil
}
