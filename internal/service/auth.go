package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"log/slog"
)

const (
	secretKey = "GdeGeneratorrr"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"uid"`
}

type AuthService struct {
	logg *slog.Logger
}

func NewAuthService(logg *slog.Logger) *AuthService {
	return &AuthService{
		logg: logg,
	}
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	s.logg.Debug("DBG_token:", token)

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not valid")
	}

	s.logg.Debug("DBG_claims:", claims.UserID)

	return claims.UserID, nil
}
