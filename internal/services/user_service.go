package services

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{Repository: r}
}

func (s *UserService) Register(ctx context.Context, userReq models.RegisterRequest) (models.User, error) {
	hash, err := GetPasswordHash(userReq.Password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Email:    userReq.Email,
		Name:     userReq.Name,
		Password: hash,
	}

	err = s.Repository.CreateUser(ctx, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
