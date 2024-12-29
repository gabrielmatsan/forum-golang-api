package tests

import (
	"context"
	"testing"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
	criptographymock "github.com/gabrielmatsan/forum-golang-api/utils/criptography-mock"
	inmemoryrepositories "github.com/gabrielmatsan/forum-golang-api/utils/in-memory-repositories"
	"github.com/stretchr/testify/assert"
)

func setupAuthenticateStudentTest() (*inmemoryrepositories.InMemoryStudentsRepository, *usecases.AuthenticateStudentUseCase) {
	studenteRepo := inmemoryrepositories.NewInMemoryStudentsRepository()
	fakeHasher := criptographymock.FakeHasher{}
	fakeEncrypt := criptographymock.FakeEncrypter{}
	useCase := usecases.NewAuthenticateStudentUseCase(studenteRepo, &fakeHasher, &fakeEncrypt)
	return studenteRepo, useCase
}

func TestAuthenticateStudentUseCase(t *testing.T) {
	ctx := context.Background()

	t.Run("should be able to authenticate a student", func(t *testing.T) {
		studentRepo, useCase := setupAuthenticateStudentTest()

		studentRepo.CreateStudent(ctx, models.NewStudent(models.StudentProps{
			Name:     "Gabriel",
			Email:    "gabriel@el.com",
			Password: "123456-hashed", // Senha já deve estar hasheada
		}))

		request := usecases.AuthenticateStudentRequest{
			Email:    "gabriel@el.com",
			Password: "123456",
		}

		res, err := useCase.Execute(ctx, request)
		assert.NoError(t, err)

		// Validações
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.AccessToken)
		t.Logf("Access Token: %s", res.AccessToken)
	})

}
