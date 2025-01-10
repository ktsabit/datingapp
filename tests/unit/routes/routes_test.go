package routes_test

import (
	"bytes"
	"datingapp/tests/mocks"
	"github.com/go-chi/jwtauth"

	"datingapp/internal/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSetupRoutes(t *testing.T) {
	mockUserHandler := new(mocks.MockUserHandler)
	mockAuthHandler := new(mocks.MockAuthHandler)
	mockSwipeHandler := new(mocks.MockSwipeHandler)
	mockJWTService := new(mocks.MockJWTService)

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	mockJWTService.On("TokenAuth").Return(tokenAuth)

	mockUserHandler.On("Register", mock.Anything, mock.Anything).Return(nil)
	mockAuthHandler.On("Login", mock.Anything, mock.Anything).Return(nil)
	mockAuthHandler.On("Refresh", mock.Anything, mock.Anything).Return(nil)

	r := routes.SetupRoutes(mockUserHandler, mockAuthHandler, mockJWTService, mockSwipeHandler)

	tests := []struct {
		method       string
		route        string
		expectedCode int
		body         string
	}{
		{method: http.MethodPost, route: "/auth/signup", expectedCode: http.StatusOK, body: `{"email":"test@example.com", "password":"password"}`},
		{method: http.MethodPost, route: "/auth/login", expectedCode: http.StatusOK, body: `{"email":"test@example.com", "password":"password"}`},
		{method: http.MethodPost, route: "/auth/refresh", expectedCode: http.StatusOK, body: `{"refreshToken":"dummyToken"}`},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(tc.method, tc.route, bytes.NewBuffer([]byte(tc.body)))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, tc.expectedCode, rr.Code)
	}
}
