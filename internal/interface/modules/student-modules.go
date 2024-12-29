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

func NewStudentsModule(db *sqlc.Queries) *StudentsModule {
	// Inicializa o repositório SQLC
	studentRepo := repositories.NewSQLCStudentsRepository(db)

	// Inicializa o caso de uso
	createStudentUseCase := usecases.NewRegisterStudentUseCase(studentRepo)

	// Inicializa o controlador
	createStudentController := dto.NewCreateStudentController(createStudentUseCase)

	return &StudentsModule{
		CreateStudentController: createStudentController,
	}
}
