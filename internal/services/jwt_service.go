package services

import (
	"datingapp/internal/models"
	"github.com/go-chi/jwtauth"
	"time"
)

type JWTConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenTTL     time.Duration
	RefreshTokenTTL    time.Duration
}

type JWTService struct {
	accessAuth  *jwtauth.JWTAuth
	refreshAuth *jwtauth.JWTAuth
	config      JWTConfig
}

func (s *JWTService) TokenAuth() *jwtauth.JWTAuth {
	return s.accessAuth
}

func NewJWTService(config JWTConfig) *JWTService {
	return &JWTService{
		accessAuth:  jwtauth.New("HS256", []byte(config.AccessTokenSecret), nil),
		refreshAuth: jwtauth.New("HS256", []byte(config.RefreshTokenSecret), nil),
		config:      config,
	}
}

func (s *JWTService) GenerateTokenPair(userID uint) (*models.TokenResponse, error) {
	now := time.Now()

	accessClaims := map[string]interface{}{
		"user_id": userID,
		"exp":     now.Add(s.config.AccessTokenTTL).Unix(),
		"iat":     now.Unix(),
		"type":    "access",
	}
	_, accessToken, err := s.accessAuth.Encode(accessClaims)

	if err != nil {
		return nil, err
	}

	refreshClaims := map[string]interface{}{
		"user_id": userID,
		"exp":     now.Add(s.config.RefreshTokenTTL).Unix(),
		"iat":     now.Unix(),
		"type":    "refresh",
	}
	_, refreshToken, err := s.refreshAuth.Encode(refreshClaims)

	if err != nil {
		return nil, err
	}

	return &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.config.AccessTokenTTL.Seconds()),
	}, nil
}

func (s *JWTService) VerifyRefreshToken(tokenString string) (uint, error) {
	token, err := s.refreshAuth.Decode(tokenString)
	if err != nil {
		return 0, err
	}

	claims := token.PrivateClaims()
	//if !ok {
	//	return 0, errors.New("invalid token claims")
	//}

	userID, _ := claims["user_id"].(float64)
	//if !ok {
	//	return 0, errors.New("invalid user_id in token")
	//}

	return uint(userID), nil
}
