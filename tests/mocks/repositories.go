package mocks

import (
	"context"
	"datingapp/internal/models"
	"datingapp/internal/repositories"
	"github.com/stretchr/testify/mock"
)

var _ repositories.UserRepositoryInterface = (*MockUserRepository)(nil)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) EmailExist(ctx context.Context, email string) bool {
	args := m.Called(ctx, email)
	return args.Bool(0)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserById(id uint) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}
