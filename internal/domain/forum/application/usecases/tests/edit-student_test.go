package tests

import (
	"context"
	"testing"

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
			Password: "123456",
		}

		response := useCase.Execute(ctx, request)

		assert.True(t, response.IsRight())

		responseValue, err := response.RightValue()

		if err != nil {
			return
		}
		assert.Equal(t, (*responseValue).GetEmail(), "gabriel@el.com")
	})

}
