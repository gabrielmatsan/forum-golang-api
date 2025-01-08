package cryptoinfra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBCryptHasher_Compare(t *testing.T) {
	hasher := NewBCryptHasher(10)

	password := "securepassword123"
	hashedPassword, _ := hasher.Hash(password)

	// Comparação bem-sucedida
	isMatch, err := hasher.Compare(hashedPassword, password)
	assert.NoError(t, err)
	assert.True(t, isMatch)

	// Comparação com senha incorreta
	isMatch, err = hasher.Compare(hashedPassword, "wrongpassword")
	assert.NoError(t, err)
	assert.False(t, isMatch)
}
