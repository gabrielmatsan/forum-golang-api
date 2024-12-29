package modules

import (
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	"github.com/gabrielmatsan/forum-golang-api/internal/infra/db/repositories"
	sqlc "github.com/gabrielmatsan/forum-golang-api/internal/infra/db/sqlc"
	"github.com/gabrielmatsan/forum-golang-api/internal/interface/api/rest/dto"
)

type StudentsModule struct {
	CreateStudentController *dto.CreateStudentController
}

func NewStudentsModule(db *sqlc.Queries, cryptoModule *CryptographyModule) *StudentsModule {
	// Repositories
	studentRepo := repositories.NewSQLCStudentsRepository(db)

	// UseCases
	createStudentUseCase := usecases.NewRegisterStudentUseCase(studentRepo, cryptoModule.Hasher)

	// Controllers
	createStudentController := dto.NewCreateStudentController(createStudentUseCase)

	return &StudentsModule{
		CreateStudentController: createStudentController,
	}
}
