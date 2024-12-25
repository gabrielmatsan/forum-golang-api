package usecases

import (
	"context"

	"github.com/gabrielmatsan/forum-golang-api/internal/core/entities"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/repositories"
	usecaseserror "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/use-cases-error"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
)

type RegisterStudentRequest struct {
	Name     string
	Email    string
	Password string
}

// Resposta do caso de uso: ou um erro ou o estudante criado
type RegisterStudentResponse = entities.Either[*usecaseserror.UseCaseError, *models.Student]

type CreateStudentUseCase struct {
	studentRepository repositories.StudentsRepository
}

func NewRegisterStudentUseCase(studentRepository repositories.StudentsRepository) *CreateStudentUseCase {
	return &CreateStudentUseCase{
		studentRepository: studentRepository,
	}
}

func (uc *CreateStudentUseCase) Execute(ctx context.Context, req RegisterStudentRequest) RegisterStudentResponse {
	existingStudent, _ := uc.studentRepository.FindByEmail(ctx, req.Email)

	if existingStudent != nil {
		return entities.Left[*usecaseserror.UseCaseError, *models.Student](
			usecaseserror.NewEmailAlreadyUsedError(req.Email),
		)
	}

	if len(req.Password) < 6 {
		return entities.Left[*usecaseserror.UseCaseError, *models.Student](
			usecaseserror.NewWeakPasswordError(),
		)
	}

	// Validar o email
	

	newStudent := models.NewStudent(models.StudentProps{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	uc.studentRepository.CreateStudent(ctx, newStudent)

	return entities.Right[*usecaseserror.UseCaseError, *models.Student](
		newStudent,
	)
}
