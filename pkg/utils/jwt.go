package utils

import (
	"Food-Delivery/config"
	"Food-Delivery/pkg/common"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Token struct {
	Access_token string `json:"access_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type TokePayload struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type claim struct {
	Payload TokePayload `json:"payload"`
	jwt.RegisteredClaims
}

func GenerateJwt(data TokePayload, cfg *config.Config) (*Token, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(time.Hour * 12))
	claim := claim{
		Payload: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        fmt.Sprintf("%d", jwt.NewNumericDate(time.Now()).UnixNano()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessToken, err := token.SignedString([]byte(cfg.App.Secret))
	if err != nil {
		return nil, err
	}
	return &Token{
		Access_token: accessToken,
		ExpiresAt:    expirationTime.Unix(),
	}, nil
}

func ValidateJwt(access_token string, cfg *config.Config) (*TokePayload, error) {
	//var registerClaim jwt.RegisteredClaims

	token, err := jwt.ParseWithClaims(access_token, &claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.App.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claim, ok := token.Claims.(*claim)

	if !ok {
		return nil, fmt.Errorf("invalid claim")
	}
	return &claim.Payload, nil
}

func RefreshJWT(accessToken string, cfg *config.Config) (*Token, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.App.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claim, ok := token.Claims.(*claim)

	if !ok {
		return nil, ErrInvalidToken
	}

	expirationTime := jwt.NewNumericDate(time.Now().Add(12 * time.Hour))
	claim.ExpiresAt = expirationTime
	claim.IssuedAt = jwt.NewNumericDate(time.Now())
	claim.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessToken, err = token.SignedString([]byte(cfg.App.Secret))
	if err != nil {
		return nil, err
	}

	return &Token{
		Access_token: accessToken,
		ExpiresAt:    expirationTime.Unix(),
	}, nil
}

var (
	ErrTokenNotFound = common.ErrUnauthorized(errors.New("token not found"))

	ErrEncodingToken = common.ErrUnauthorized(errors.New("error encoding token"))

	ErrInvalidToken = common.ErrUnauthorized(errors.New("invalid token"))
)
