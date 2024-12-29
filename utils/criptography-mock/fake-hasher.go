package criptographymock

import "fmt"

type FakeHasher struct{}

func (f *FakeHasher) Hash(plain string) (string, error) {

	return fmt.Sprintf("%s-hashed", plain), nil
}

func (f *FakeHasher) Compare(plain, hash string) (bool, error) {
	return fmt.Sprintf("%s-hashed", plain) == hash, nil
}
