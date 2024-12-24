package repositories

import (
	"context"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
)

type StudentsRepository interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	FindByEmail(ctx context.Context, email string) (*models.Student, error)
}
