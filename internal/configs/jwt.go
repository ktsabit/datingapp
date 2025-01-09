package configs

import (
	"datingapp/internal/models"
	"github.com/go-chi/jwtauth"
	"os"
	"time"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_KEY")), nil)
}

type JWTConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenTTL     time.Duration // e.g., 15 minutes
	RefreshTokenTTL    time.Duration // e.g., 7 days
}

type JWTService struct {
	accessAuth  *jwtauth.JWTAuth
	refreshAuth *jwtauth.JWTAuth
	config      JWTConfig
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
	_, accessToken, _ := s.accessAuth.Encode(accessClaims)

	refreshClaims := map[string]interface{}{
		"user_id": userID,
		"exp":     now.Add(s.config.RefreshTokenTTL).Unix(),
		"iat":     now.Unix(),
		"type":    "refresh",
	}
	_, refreshToken, _ := s.refreshAuth.Encode(refreshClaims)

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
