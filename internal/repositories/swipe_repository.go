package repositories

import (
	"context"
	"datingapp/internal/models"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type SwipeRepository struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewSwipeRepository(
	db *gorm.DB,
	r *redis.Client) *SwipeRepository {
	return &SwipeRepository{DB: db, Redis: r}
}

//func (r *SwipeRepository) GetDailySwipe(
//	ctx context.Context,
//	userID uint,
//) (int64, error) {
//	key := fmt.Sprintf("swipes:%d", userID)
//	count, err := r.Redis.Get(ctx, key).Int64()
//	if err != nil {
//		return 0, err
//	}
//	return count, nil
//}

func (r *SwipeRepository) IncrementDailySwipe(
	ctx context.Context,
	userID uint,
) (int64, error) {
	key := fmt.Sprintf("swipes:%d", userID)

	exists, err := r.Redis.Exists(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if exists == 0 {
		now := time.Now()
		next := now.AddDate(0, 0, 1)
		ttl := next.Sub(now)

		err := r.Redis.Set(ctx, key, 0, ttl).Err()
		if err != nil {
			return 0, err
		}
	}

	return r.Redis.Incr(ctx, key).Result()
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
	).Update("is_match", true)

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
