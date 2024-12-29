package cryptoinfra

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type BCryptHasher struct {
	cost int
}

func NewBCryptHasher(cost int) *BCryptHasher {
	return &BCryptHasher{
		cost: cost,
	}
}

func (b *BCryptHasher) Hash(plain string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plain), b.cost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (b *BCryptHasher) Compare(hashed, plain string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
