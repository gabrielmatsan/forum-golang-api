package routes

import (
	"github.com/gabrielmatsan/forum-golang-api/internal/interface/api/rest/dto"
	"github.com/gin-gonic/gin"
)

func RegisterStudentsRoutes(
	r *gin.Engine,
	studentController *dto.CreateStudentController,
) {

	studentsRoutes := r.Group("/students")
	{
		studentsRoutes.POST("/", studentController.Handle)
	}
}
