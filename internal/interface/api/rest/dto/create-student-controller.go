package dto

import (
	"net/http"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateStudentController struct {
	useCase *usecases.CreateStudentUseCase
}

type studentBodyValidationSchema struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewCreateStudentController(useCase *usecases.CreateStudentUseCase) *CreateStudentController {
	return &CreateStudentController{
		useCase: useCase,
	}
}

func (ctrl *CreateStudentController) Handle(c *gin.Context) {
	validate := validator.New()

	var body studentBodyValidationSchema

	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := validate.Struct(body); err != nil {

		validationError := err.(validator.ValidationErrors)
		errorMessage := make(map[string]string)

		for _, validationErr := range validationError {
			errorMessage[validationErr.Field()] = validationErr.Tag()
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	err := ctrl.useCase.Execute(c.Request.Context(), usecases.RegisterStudentRequest{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student created successfully"})
}
