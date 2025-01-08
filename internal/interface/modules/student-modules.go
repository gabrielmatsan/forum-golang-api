package modules

import (
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	"github.com/gabrielmatsan/forum-golang-api/internal/infra/db/repositories"
	sqlc "github.com/gabrielmatsan/forum-golang-api/internal/infra/db/sqlc"
	"github.com/gabrielmatsan/forum-golang-api/internal/interface/api/rest/dto"
)

type StudentsModule struct {
	CreateStudentController       *dto.CreateStudentController
	AuthenticateStudentController *dto.AuthenticateStudentController
}

func NewStudentsModule(db *sqlc.Queries, cryptoModule *CryptographyModule) *StudentsModule {
	// Repositories
	studentRepo := repositories.NewSQLCStudentsRepository(db)

	// UseCases
	createStudentUseCase := usecases.NewRegisterStudentUseCase(studentRepo, cryptoModule.Hasher)
	authenticateStudentUseCase := usecases.NewAuthenticateStudentUseCase(studentRepo, cryptoModule.Hasher, cryptoModule.Encrypter)

	// Controllers
	createStudentController := dto.NewCreateStudentController(createStudentUseCase)
	authenticateStudentController := dto.NewAuthenticateStudentController(authenticateStudentUseCase)

	return &StudentsModule{
		CreateStudentController:       createStudentController,
		AuthenticateStudentController: authenticateStudentController,
	}
}
