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
type RegisterStudentResponse = entities.Either[*usecaseserror.EmailAlreadyUsed, *models.Student]

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
		// Use apenas um único * para o tipo no Left e Right
		return entities.Left[*usecaseserror.EmailAlreadyUsed, *models.Student](
			usecaseserror.NewEmailAlreadyUsedError(req.Email),
		)
	}

	// Caso contrário, continue com a lógica para criar um novo estudante
	newStudent := models.NewStudent(models.StudentProps{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	uc.studentRepository.CreateStudent(ctx, newStudent)

	return entities.Right[*usecaseserror.EmailAlreadyUsed, *models.Student](
		newStudent,
	)
}
