package modules

import (
	"log"

	cryptoinfra "github.com/gabrielmatsan/forum-golang-api/internal/infra/crypto-infra"
)

type CryptographyModule struct {
	Hasher    *cryptoinfra.BCryptHasher
	Encrypter *cryptoinfra.JWTEncrypter
}

func NewCryptographyModule() *CryptographyModule {
	var HASH_SALT_LENGTH = 10

	hasher := cryptoinfra.NewBCryptHasher(HASH_SALT_LENGTH)
	encrypter, err := cryptoinfra.NewJWTEncrypterFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize JWT encrypter: %v", err)
	}

	return &CryptographyModule{
		Hasher:    hasher,
		Encrypter: encrypter,
	}
}
