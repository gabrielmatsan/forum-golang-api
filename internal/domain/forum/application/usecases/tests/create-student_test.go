package tests

import (
	"context"
	"fmt"
	"testing"

	usecaseserror "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/use-cases-error"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	inmemoryrepositories "github.com/gabrielmatsan/forum-golang-api/utils/in-memory-repositories"
	"github.com/stretchr/testify/assert"
)

func setup() (*inmemoryrepositories.InMemoryStudentsRepository, *usecases.CreateStudentUseCase) {
	studenteRepo := inmemoryrepositories.NewInMemoryStudentsRepository()
	useCase := usecases.NewRegisterStudentUseCase(studenteRepo)
	return studenteRepo, useCase
}

func TestCreateStudentUseCase(t *testing.T) {
	// studenteRepo := inmemoryrepositories.NewInMemoryStudentsRepository()
	// useCase := usecases.NewRegisterStudentUseCase(studenteRepo)

	ctx := context.Background()

	t.Run("should be able to register a new student", func(t *testing.T) {
		_, useCase := setup()

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

	t.Run("should not be able to register a student with an email already in use", func(t *testing.T) {

		studenteRepo, useCase := setup()
		request := usecases.RegisterStudentRequest{
			Name:     "Gabriel",
			Email:    "gabriel@ro.com",
			Password: "123456",
		}
		response := useCase.Execute(ctx, request)

		assert.True(t, response.IsRight())

		request = usecases.RegisterStudentRequest{
			Name:     "Hoalnf",
			Email:    "gabriel@ro.com",
			Password: "123456",
		}

		response = useCase.Execute(ctx, request)
		assert.True(t, response.IsLeft())

		studenteRepo.Mu.Lock()
		defer studenteRepo.Mu.Unlock()

		assert.Equal(t, 1, len(studenteRepo.Students))
	})

	t.Run("should not be able to register a student with a weak password", func(t *testing.T) {
		_, useCase := setup()
		request := usecases.RegisterStudentRequest{
			Name:     "Gabriel",
			Email:    "gabriel@ht.com",
			Password: "123",
		}

		response := useCase.Execute(ctx, request)

		assert.True(t, response.IsLeft())

		// Verifica o erro retornado
		errorValue, err := response.LeftValue()
		assert.NoError(t, err)

		if errorValue != nil {
			// Imprime o erro no log do teste
			t.Logf("Erro retornado: Code=%s, Message=%s", (*errorValue).Code, (*errorValue).Message)

			// Ou usando fmt.Println, se quiser saída direta no terminal
			fmt.Printf("Erro retornado: Code=%s, Message=%s\n", (*errorValue).Code, (*errorValue).Message)

			// Valida o código e mensagem do erro
			assert.Equal(t, usecaseserror.NewWeakPasswordError().Code, (*errorValue).Code)
			assert.Equal(t, usecaseserror.NewWeakPasswordError().Message, (*errorValue).Message)
		}
	})
}
