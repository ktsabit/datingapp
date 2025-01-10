package repositories

import (
	"context"
	"datingapp/internal/models"
	"gorm.io/gorm"
)

type SwipeRepository struct {
	DB *gorm.DB
}

func NewSwipeRepository(db *gorm.DB) *SwipeRepository {
	return &SwipeRepository{DB: db}
}

func (r *SwipeRepository) CreateSwipe(
	ctx context.Context,
	userID uint,
	swipedUserID uint,
	swipedDirection models.SwipeDirectionEnum,
) (*models.Swipe, error) {
	swipe := &models.Swipe{
		UserID:         userID,
		TargetID:       swipedUserID,
		SwipeDirection: swipedDirection,
	}

	err := r.DB.WithContext(ctx).Create(swipe).Error

	if err != nil {
		return nil, err
	}
	return swipe, nil
}

func (r *SwipeRepository) SwipeMatch(ctx context.Context, swiperID uint, swipedID uint) (*models.Swipe, error) {
	res := r.DB.WithContext(ctx).Model(&models.Swipe{}).Where(
		"(user_id = ? AND target_id = ?) OR (target_id = ? AND user_id = ?)",
		swiperID, swipedID, swipedID, swiperID,
	)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return &models.Swipe{}, nil
}

func (r *SwipeRepository) CheckReverseSwipe(
	ctx context.Context,
	userID uint,
	targetUserID uint,
	swipedDirection models.SwipeDirectionEnum,
) (*models.Swipe, error) {
	var swipe models.Swipe
	err := r.DB.WithContext(ctx).
		Where("user_id = ? AND target_id = ? AND swipe_direction = ?", targetUserID, userID, swipedDirection).
		First(&swipe).Error

	if err != nil {
		return nil, err
	}

	return &swipe, nil
}
