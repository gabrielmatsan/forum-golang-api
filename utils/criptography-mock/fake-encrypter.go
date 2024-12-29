package criptographymock

import (
	"encoding/json"
)

type FakeEncrypter struct{}

func (f *FakeEncrypter) Encrypt(payload map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
