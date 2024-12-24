package models

import (
	"time"

	"github.com/gabrielmatsan/forum-golang-api/internal/core/entities"
)

type AnswerProps struct {
	AuthorID   *entities.UniqueEntityID
	QuestionID *entities.UniqueEntityID
	Content    string
	CreatedAt  time.Time
	UpdateAt   *time.Time
	// After, add Attachments []Attachment
}

type Answer struct {
	*entities.Entity[AnswerProps]
}

func NewAnswer(props AnswerProps, id ...*entities.UniqueEntityID) *Answer {
	var entityID *entities.UniqueEntityID
	if len(id) > 0 {
		entityID = id[0]
	} else {
		entityID = entities.NewUniqueEntityID()
	}
	return &Answer{
		Entity: entities.NewEntity(props, entityID),
	}
}

func (a *Answer) GetAuthorID() string {
	return a.Props().AuthorID.ToString()
}

func (a *Answer) GetQuestionID() string {
	return a.Props().QuestionID.ToString()
}

func (a *Answer) GetContent() string {
	return a.Props().Content
}

func (a *Answer) GetCreatedAt() time.Time {
	return a.Props().CreatedAt
}

func (a *Answer) GetUpdateAt() *time.Time {
	return a.Props().UpdateAt
}

func (a *Answer) touch() {
	now := time.Now()
	a.Props().UpdateAt = &now
}

func (a *Answer) SetContent(content string) {
	a.Props().Content = content
	a.touch()
}


