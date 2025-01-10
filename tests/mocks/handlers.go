package mocks

import (
	"github.com/stretchr/testify/mock"
	"net/http"
)

type MockAuthHandler struct {
	mock.Mock
}

func (m *MockAuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockAuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

type MockUserHandler struct {
	mock.Mock
}

func (m *MockUserHandler) Register(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

type MockSwipeHandler struct {
	mock.Mock
}

func (m *MockSwipeHandler) HandleSwipe(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}
