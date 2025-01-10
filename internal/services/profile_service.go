package services

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
)

type ProfileService struct {
	Repository repositories.UserRepositoryInterface
}

func NewProfileService(r repositories.UserRepositoryInterface) *ProfileService {
	return &ProfileService{Repository: r}
}

func (fs *ProfileService) GenerateFeed(ctx context.Context, userID uint) ([]*models.User, error) {
	return nil, nil
}
