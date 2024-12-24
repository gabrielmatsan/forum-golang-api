package entities

import "github.com/google/uuid"

type UniqueEntityID struct {
	value string
}

func NewUniqueEntityID(value ...string) *UniqueEntityID {
	var id string
	if len(value) > 0 && value[0] != "" {
		id = value[0]
	} else {
		id = uuid.New().String() // Gera um UUID aleatório
	}

	return &UniqueEntityID{
		value: id,
	}
}

func (u *UniqueEntityID) ToString() string {
	return u.value
}

func (u *UniqueEntityID) Equals(other *UniqueEntityID) bool {
	if other == nil {
		return false
	}
	return u.value == other.value
}
