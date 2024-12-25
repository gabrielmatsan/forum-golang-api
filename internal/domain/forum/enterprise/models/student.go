package models

import (
	"github.com/gabrielmatsan/forum-golang-api/internal/core/entities"
)

type StudentProps struct {
	Name     string
	Email    string
	Password string
}

type Student struct {
	*entities.Entity[StudentProps]
}

func NewStudent(props StudentProps, id ...*entities.UniqueEntityID) *Student {
	return &Student{
		Entity: entities.NewEntity(props, id...),
	}
}

// Métodos Getter
func (s *Student) GetName() string {
	return s.Props().Name
}

func (s *Student) GetEmail() string {
	return s.Props().Email
}

func (s *Student) GetPassword() string {
	return s.Props().Password
}

func (s *Student) GetID() string {
	return s.ID().ToString()
}
