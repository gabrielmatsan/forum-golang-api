package usecases

import (
	"context"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/repositories"
	usecaseserror "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/use-cases-error"
)

type EditStudentRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EditStudentUseCase struct {
	studentRepository repositories.StudentsRepository
}

func NewEditStudenUseCase(studentRepository repositories.StudentsRepository) *EditStudentUseCase {
	return &EditStudentUseCase{
		studentRepository: studentRepository,
	}
}

func (uc *EditStudentUseCase) Execute(ctx context.Context, req EditStudentRequest) error {
	student, _ := uc.studentRepository.FindById(ctx, req.ID)

	if student == nil {
		return usecaseserror.ResourceNotFoundError("Student", req.ID)
	}

	// Valida e atualiza o email, se fornecido
	if req.Email != student.GetEmail() {
		existingStudent, err := uc.studentRepository.FindByEmail(ctx, req.Email)

		if err != nil {
			return usecaseserror.NewInternalError()
		}

		if existingStudent != nil {
			return usecaseserror.NewEmailAlreadyUsedError(req.Email)
		}
		student.SetEmail(req.Email)
	}

	// Valida e atualiza a senha, se fornecida
	if len(req.Password) < 6 {
		return usecaseserror.NewWeakPasswordError()
	}

	student.SetPassword(req.Password)

	// Salva as alterações no repositório
	if err := uc.studentRepository.UpdateStudent(ctx, student); err != nil {
		return usecaseserror.NewInternalError()
	}

	return nil
}
