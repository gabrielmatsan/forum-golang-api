package cryptoinfra

import (
	"crypto/rsa"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTEncrypter struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewJWTEncrypterFromEnv carrega as chaves privada e pública de variáveis de ambiente.
func NewJWTEncrypterFromEnv() (*JWTEncrypter, error) {
	privateKeyContent := os.Getenv("JWT_PRIVATE_KEY")
	if privateKeyContent == "" {
		return nil, errors.New("missing JWT_PRIVATE_KEY environment variable")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyContent))
	if err != nil {
		return nil, err
	}

	publicKeyContent := os.Getenv("JWT_PUBLIC_KEY")
	if publicKeyContent == "" {
		return nil, errors.New("missing JWT_PUBLIC_KEY environment variable")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyContent))
	if err != nil {
		return nil, err
	}

	return &JWTEncrypter{privateKey: privateKey, publicKey: publicKey}, nil
}

func (j *JWTEncrypter) Encrypt(payload map[string]interface{}) (string, error) {
	claims := jwt.MapClaims(payload)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

func (j *JWTEncrypter) Validate(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
