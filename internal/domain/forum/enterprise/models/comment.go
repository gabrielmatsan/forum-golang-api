package models

import (
	"time"

	"github.com/gabrielmatsan/forum-golang-api/internal/core/entities"
)

type CommentProps struct {
	AuthorID  *entities.UniqueEntityID
	Content   string
	CreatedAt time.Time
	UpdateAt  *time.Time
}

type Comment[T any] struct {
	*entities.Entity[T]
}

func NewComment[T CommentProps](props T, id ...*entities.UniqueEntityID) *Comment[T] {
	var entityID *entities.UniqueEntityID
	if len(id) > 0 {
		entityID = id[0]
	} else {
		entityID = entities.NewUniqueEntityID()
	}
	return &Comment[T]{
		Entity: entities.NewEntity(props, entityID),
	}
}
