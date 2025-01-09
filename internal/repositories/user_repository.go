package repositories

import (
	"context"
	"datingapp/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepositoryImpl) GetUserById(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, "email = ?", email).Error
	return &user, err
}

func (r *UserRepositoryImpl) EmailExist(ctx context.Context, email string) bool {
	var count int64
	r.DB.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}
