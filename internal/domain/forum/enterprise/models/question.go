package models

import (
	"strings"
	"time"

	"github.com/gabrielmatsan/forum-golang-api/internal/core/entities"
	. "github.com/gabrielmatsan/forum-golang-api/internal/core/entities"

	. "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models/value-object"
)

type QuestionProps struct {
	authorID     *UniqueEntityID
	bestAnswerID *UniqueEntityID
	Title        string
	Content      string
	Slug         *Slug
	CreatedAt    time.Time
	UpdateAt     *time.Time
}

type Question struct {
	*entities.Entity[QuestionProps]
}

func NewQuestion(props QuestionProps, id ...*UniqueEntityID) *Question {
	return &Question{
		Entity: NewEntity(props, id...),
	}
}

func (q *Question) GetAuthorID() string {
	return q.Props().authorID.ToString()
}

func (q *Question) GetBestAnswerID() string {
	return q.Props().bestAnswerID.ToString()
}

func (q *Question) GetTitle() string {
	return q.Props().Title
}

func (q *Question) GetContent() string {
	return q.Props().Content
}

func (q *Question) GetSlug() string {
	return q.Props().Slug.Value()
}

func (q *Question) GetCreatedAt() time.Time {
	return q.Props().CreatedAt
}

func (q *Question) GetUpdateAt() *time.Time {
	return q.Props().UpdateAt
}

func (q *Question) GetExcerpt() string {
	if len(q.Props().Content) <= 120 {
		return q.Props().Content
	}
	return strings.TrimSpace(q.Props().Content[:120]) + "..."
}

func (q *Question) touch() {
	now := time.Now()
	q.Props().UpdateAt = &now
}

func (q *Question) SetTitle(title string) {
	q.Props().Title = title
	q.Props().Slug = q.Props().Slug.CreateSlugFromText(title)
	q.touch()
}

func (q *Question) SetContent(content string) {
	q.Props().Content = content
	q.touch()
}

func (q *Question) SetBestAnswerID(newBestAnswerID *UniqueEntityID) {
	if newBestAnswerID != nil && !newBestAnswerID.Equals(q.Props().bestAnswerID) {
		// TODO: Adiciona o evento de notificação
	}

	if newBestAnswerID != q.Props().bestAnswerID {
		q.Props().bestAnswerID = newBestAnswerID
		q.touch()
	}

}
