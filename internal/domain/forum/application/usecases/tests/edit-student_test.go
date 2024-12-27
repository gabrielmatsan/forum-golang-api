package tests

import (
	"context"
	"testing"

	usecaseserror "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/use-cases-error"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
	inmemoryrepositories "github.com/gabrielmatsan/forum-golang-api/utils/in-memory-repositories"
	"github.com/stretchr/testify/assert"
)

func setupEditStudentTest() (*inmemoryrepositories.InMemoryStudentsRepository, *usecases.EditStudentUseCase) {
	studenteRepo := inmemoryrepositories.NewInMemoryStudentsRepository()
	useCase := usecases.NewEditStudenUseCase(studenteRepo)
	return studenteRepo, useCase
}

func TestEditStudentUseCase(t *testing.T) {

	ctx := context.Background()

	t.Run("should be able to edit email and password", func(t *testing.T) {
		studenteRepo, useCase := setupEditStudentTest()

		student := models.NewStudent(models.StudentProps{
			Name:     "Gabriel",
			Email:    "gabriel@li.com",
			Password: "123456",
		})

		studenteRepo.CreateStudent(ctx, student)

		request := usecases.EditStudentRequest{
			ID:       student.GetID(),
			Name:     "Gabriel",
			Email:    "gabriel@el.com",
			Password: "gabriel123",
		}

		err := useCase.Execute(ctx, request)
		assert.NoError(t, err)

		studenteRepo.Mu.Lock()
		defer studenteRepo.Mu.Unlock()

		assert.Equal(t, 1, len(studenteRepo.Students))
		s := studenteRepo.Students[0]
		assert.Equal(t, s.GetEmail(), request.Email)
		assert.Equal(t, s.GetPassword(), request.Password)
	})

	t.Run("should not be able to edit email to an already used email", func(t *testing.T) {
		studenteRepo, useCase := setupEditStudentTest()

		student1 := models.NewStudent(models.StudentProps{
			Name:     "Gabriel",
			Email:    "gabriel@hot.com",
			Password: "123456",
		})

		studenteRepo.CreateStudent(ctx, student1)

		student2 := models.NewStudent(models.StudentProps{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Password: "123456",
		})

		studenteRepo.CreateStudent(ctx, student2)

		request := usecases.EditStudentRequest{
			ID:       student1.GetID(),
			Name:     "Gabriel",
			Email:    "johndoe@example.com",
			Password: "123456",
		}

		err := useCase.Execute(ctx, request)

		assert.Error(t, err)
		assert.IsType(t, usecaseserror.NewEmailAlreadyUsedError(request.Email), err)
	})
}
