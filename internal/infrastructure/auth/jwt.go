package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

const year = time.Hour * 24 * 365

type JWTService struct {
	secretKey []byte
}

func NewJwtService(secretKey string) *JWTService {
	return &JWTService{secretKey: []byte(secretKey)}
}

func (s *JWTService) CreatePermanentToken(profileId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"profileId": profileId.String(),
		"exp":       time.Now().Add(100 * year).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *JWTService) ProfileIdFromToken(tokenString string) *uuid.UUID {
	claims := &jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})

	if err != nil {
		return nil
	}

	profileId, ok := (*claims)["profileId"].(string)
	if !ok {
		return nil
	}

	u, err := uuid.Parse(profileId)
	if err != nil {
		return nil
	}
	return &u
}
