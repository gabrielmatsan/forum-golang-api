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

func (f *FakeEncrypter) Validate(token string) (map[string]interface{}, error) {
	var payload map[string]interface{}

	// Decodifica o token JSON para o mapa
	err := json.Unmarshal([]byte(token), &payload)
	if err != nil {
		return nil, err // Retorna erro se o token não for um JSON válido
	}

	return payload, nil
}
