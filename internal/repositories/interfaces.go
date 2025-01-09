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
