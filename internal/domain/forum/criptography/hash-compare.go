package criptography

type HashCompare interface {
	Compare(hash string, password string) (bool, error)
}
