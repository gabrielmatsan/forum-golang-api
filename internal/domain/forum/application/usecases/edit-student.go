package usecases

import (
	"github.com/gabrielmatsan/forum-golang-api/internal/core/entities"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/repositories"
	usecaseserror "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/use-cases-error"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
)

type EditStudentRequest struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type EditStudentResponse = entities.Either[*usecaseserror.UseCaseError, *models.Student]

type EditStudentUseCase struct {
	studentRepository repositories.StudentsRepository
}

func NewEditStudenUseCase(studentRepository repositories.StudentsRepository) *EditStudentUseCase {
	return &EditStudentUseCase{
		studentRepository: studentRepository,
	}
}

// func (uc *EditStudentUseCase) Execute(ctx context.Context, req EditStudentRequest) EditStudentResponse {
// 	student, _ := uc.studentRepository.FindById(ctx, req.ID)

// 	if student == nil {
// 		// Retorna um erro de recurso não encontrado
// 		return entities.Left[*usecaseserror.UseCaseError, *models.Student](
// 			usecaseserror.ResourceNotFoundError("Student", req.ID),
// 		)
// 	}

// }
