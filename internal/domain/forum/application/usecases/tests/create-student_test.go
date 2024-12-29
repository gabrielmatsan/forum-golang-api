package tests

import (
	"context"
	"testing"

	usecaseserror "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/use-cases-error"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	criptographymock "github.com/gabrielmatsan/forum-golang-api/utils/criptography-mock"
	inmemoryrepositories "github.com/gabrielmatsan/forum-golang-api/utils/in-memory-repositories"
	"github.com/stretchr/testify/assert"
)

func setupCreateStudentTest() (*inmemoryrepositories.InMemoryStudentsRepository, *usecases.CreateStudentUseCase) {
	studenteRepo := inmemoryrepositories.NewInMemoryStudentsRepository()
	fakeHasher := criptographymock.FakeHasher{}
	useCase := usecases.NewRegisterStudentUseCase(studenteRepo, &fakeHasher)
	return studenteRepo, useCase
}

func TestCreateStudentUseCase(t *testing.T) {
	ctx := context.Background()

	t.Run("should be able to register a new student with hashed password", func(t *testing.T) {
		studentRepo, useCase := setupCreateStudentTest()

		request := usecases.RegisterStudentRequest{
			Name:     "Gabriel",
			Email:    "gabriel@el.com",
			Password: "123456",
		}

		err := useCase.Execute(ctx, request)
		assert.NoError(t, err)

		studentRepo.Mu.Lock()
		defer studentRepo.Mu.Unlock()

		assert.Equal(t, 1, len(studentRepo.Students)) // Verifica se um estudante foi salvo
		student := studentRepo.Students[0]

		assert.Equal(t, request.Name, (*student).GetName()) // Compara os valores
		assert.Equal(t, request.Email, (*student).GetEmail())

		expectedHashedPassword := "123456-hashed"
		assert.Equal(t, expectedHashedPassword, (*student).GetPassword())

		assert.NotEqual(t, "", (*student).GetID()) // Verifica se o ID foi gerado
	})

	t.Run("should not be able to register a student with an email already in use", func(t *testing.T) {

		studenteRepo, useCase := setupCreateStudentTest()
		request := usecases.RegisterStudentRequest{
			Name:     "Gabriel",
			Email:    "gabriel@ro.com",
			Password: "123456",
		}
		response := useCase.Execute(ctx, request)

		assert.NoError(t, response)

		// Tenta registrar o mesmo estudante
		err := useCase.Execute(ctx, request)
		assert.Error(t, err)
		assert.IsType(t, usecaseserror.NewEmailAlreadyUsedError(request.Email), err)
		studenteRepo.Mu.Lock()
		defer studenteRepo.Mu.Unlock()

		assert.Equal(t, 1, len(studenteRepo.Students))
	})

	t.Run("should not be able to register a student with a weak password", func(t *testing.T) {
		_, useCase := setupCreateStudentTest()
		request := usecases.RegisterStudentRequest{
			Name:     "Gabriel",
			Email:    "gabriel@ht.com",
			Password: "123", //Weak password
		}

		err := useCase.Execute(ctx, request)

		assert.Error(t, err)
		assert.IsType(t, usecaseserror.NewWeakPasswordError(), err)

		if err != nil {
			// Valida que é o erro correto
			useCaseError := err.(*usecaseserror.UseCaseError)
			assert.Equal(t, "WEAK_PASSWORD", useCaseError.Code)
			assert.Equal(t, "Password is too weak", useCaseError.Message)
		}
	})
}
