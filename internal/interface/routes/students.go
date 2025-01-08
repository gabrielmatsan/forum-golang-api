package routes

import (
	"github.com/gabrielmatsan/forum-golang-api/internal/interface/api/rest/dto"
	"github.com/gin-gonic/gin"
)

func RegisterStudentsRoutes(
	r *gin.Engine,
	createStudentController *dto.CreateStudentController,
	authController *dto.AuthenticateStudentController,
) {
	studentsRoutes := r.Group("/students")
	{
		// Rota para criar um estudante
		studentsRoutes.POST("/", createStudentController.Handle)

		// Rota para autenticar um estudante (login)
		studentsRoutes.POST("/login", authController.Handle)
	}
}