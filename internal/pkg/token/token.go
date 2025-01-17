package token

import (
	"authService/internal/config"
	"authService/internal/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Service struct {
	SecretKey string
}

func NewService(config *config.Config) *Service {
	return &Service{
		SecretKey: config.JwtKey,
	}
}

func (s *Service) GenerateAccessToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"token_type": "access",
		"user_id":    user.ID,
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Service) GenerateRefreshToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"token_type": "refresh",
		"user_id":    user.ID,
		"exp":        time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.SecretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *Service) CheckToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")

}
