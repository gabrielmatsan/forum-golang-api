package modules

import (
	cryptoinfra "github.com/gabrielmatsan/forum-golang-api/internal/infra/crypto-infra"
)

type CryptographyModule struct {
	Hasher *cryptoinfra.BCryptHasher
}

func NewCryptographyModule() *CryptographyModule {
	var HASH_SALT_LENGTH = 10

	return &CryptographyModule{
		Hasher: cryptoinfra.NewBCryptHasher(HASH_SALT_LENGTH),
	}
}
