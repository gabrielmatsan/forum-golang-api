package mappers

import (
	"github.com/gabrielmatsan/forum-golang-api/internal/core/entities"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
	pgstore "github.com/gabrielmatsan/forum-golang-api/internal/infra/db/sqlc"
	"github.com/google/uuid"
)

// Mapper: pgstore.Student -> models.Student
func ToDomainStudent(pgStudent *pgstore.Student) *models.Student {
	return models.NewStudent(
		models.StudentProps{
			Name:     pgStudent.Name,
			Email:    pgStudent.Email,
			Password: pgStudent.Password,
		},
		entities.NewUniqueEntityID(pgStudent.ID.String()),
	)
}

// Mapper: pgstore.Student -> models.Student
func ToPgStoreStudent(domainStudent *models.Student) pgstore.Student {
	return pgstore.Student{
		ID:       uuid.MustParse(domainStudent.GetID()),
		Name:     domainStudent.GetName(),
		Email:    domainStudent.GetEmail(),
		Password: domainStudent.GetPassword(),
	}
}
