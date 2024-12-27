package repositories

import (
	"context"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/repositories"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
	"github.com/gabrielmatsan/forum-golang-api/internal/infra/db/mappers"
	pgstore "github.com/gabrielmatsan/forum-golang-api/internal/infra/db/sqlc"
	"github.com/google/uuid"
)

type SQLCStudentsRepository struct {
	queries *pgstore.Queries
}

func NewSQLCStudentsRepository(queries *pgstore.Queries) repositories.StudentsRepository {
	return &SQLCStudentsRepository{queries: queries}
}

func (r *SQLCStudentsRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	return r.queries.CreateStudent(ctx, pgstore.CreateStudentParams{
		ID:       uuid.MustParse(student.GetID()),
		Name:     student.GetName(),
		Email:    student.GetEmail(),
		Password: student.GetPassword(),
	})
}

func (r *SQLCStudentsRepository) FindByEmail(ctx context.Context, email string) (*models.Student, error) {
	pgStudent, err := r.queries.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainStudent(&pgStudent), nil
}

func (r *SQLCStudentsRepository) FindById(ctx context.Context, id string) (*models.Student, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	pgStudent, err := r.queries.FindById(ctx, uid)
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainStudent(&pgStudent), nil
}

func (r *SQLCStudentsRepository) UpdateStudent(ctx context.Context, student *models.Student) error {
	return r.queries.UpdateStudent(ctx, pgstore.UpdateStudentParams{
		ID:       uuid.MustParse(student.GetID()),
		Name:     student.GetName(),
		Email:    student.GetEmail(),
		Password: student.GetPassword(),
	})
}
