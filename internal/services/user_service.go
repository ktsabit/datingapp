package services

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository repositories.UserRepositoryInterface
}

func NewUserService(r repositories.UserRepositoryInterface) *UserService {
	return &UserService{Repository: r}
}

func (s *UserService) Register(ctx context.Context, userReq models.RegisterRequest) (models.User, error) {
	hash, err := GetPasswordHash(userReq.Password)
	if err != nil {
		return models.User{}, err
	}

	emailExist := s.Repository.EmailExist(ctx, userReq.Email)
	if emailExist {
		return models.User{}, errors.New("email exist")
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
