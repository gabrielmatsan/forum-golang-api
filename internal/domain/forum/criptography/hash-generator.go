package criptography

type HashGenerator interface {
	Hash(string) (string, error)
}
