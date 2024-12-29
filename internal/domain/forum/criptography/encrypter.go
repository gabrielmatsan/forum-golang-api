package criptography

type Encrypter interface {
	Encrypt(payload map[string]interface{}) (string, error)
}
