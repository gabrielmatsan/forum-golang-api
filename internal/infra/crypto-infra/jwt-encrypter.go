package cryptoinfra

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtEncrypter struct {
	secret     string
	expiration time.Duration
}

func NewJwtEncrypter(secret string, expiration time.Duration) *JwtEncrypter {
	return &JwtEncrypter{
		secret:     secret,
		expiration: expiration,
	}
}

func (j *JwtEncrypter) Encrypt(payload map[string]interface{}, isRefresh bool) (string, error) {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(j.expiration).Unix(),
		"iat":  time.Now().Unix(),
		"type": "access",
	}

	if isRefresh {
		claims["exp"] = time.Now().Add(j.expiration * 24).Unix()
		claims["type"] = "refresh"
	}

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JwtEncrypter) Validate(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["type"] == "refresh" {
			// Logic for refresh token validation
			return claims, nil
		}

		if claims["type"] == "access" {
			// Logic for access token validation
			return claims, nil
		}

		return nil, errors.New("invalid token type")
	}
	return nil, errors.New("invalid token")
}
func (j *JwtEncrypter) GenerateRefreshToken(payload map[string]interface{}) (string, error) {
	return j.Encrypt(payload, true) // Pass true for refresh token
}
