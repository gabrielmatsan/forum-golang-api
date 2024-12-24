package tests

import (
	"context"
	"testing"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	inmemoryrepositories "github.com/gabrielmatsan/forum-golang-api/utils/in-memory-repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudentUseCase(t *testing.T) {
	studenteRepo := inmemoryrepositories.NewInMemoryStudentsRepository()
	useCase := usecases.NewRegisterStudentUseCase(studenteRepo)

	ctx := context.Background()

	t.Run("should be able to register a new student", func(t *testing.T) {
		request := usecases.RegisterStudentRequest{
			Name:     "Gabriel",
			Email:    "gabriel@el.com",
			Password: "123456",
		}

		response := useCase.Execute(ctx, request)

		assert.True(t, response.IsRight())

		// Captura o estudante e trata o erro
		student, err := response.RightValue()
		assert.NoError(t, err)    // Garante que não houve erro
		assert.NotNil(t, student) // Verifica que o estudante não é nulo

		assert.Equal(t, request.Name, (*student).GetName()) // Compara os valores
		assert.Equal(t, request.Email, (*student).GetEmail())
		assert.Equal(t, request.Password, (*student).GetPassword())

		assert.NotEqual(t, "", (*student).GetID()) // Verifica se o ID foi gerado
	})
}
