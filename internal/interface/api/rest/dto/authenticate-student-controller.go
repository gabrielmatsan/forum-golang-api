package dto

import (
	"net/http"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthenticateStudentController struct {
	useCase *usecases.AuthenticateStudentUseCase
}

type authenticateStudentBodyValidationSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewAuthenticateStudentController(useCase *usecases.AuthenticateStudentUseCase) *AuthenticateStudentController {
	return &AuthenticateStudentController{
		useCase: useCase,
	}
}

func (ctrl *AuthenticateStudentController) Handle(c *gin.Context) {
	validate := validator.New()

	var body authenticateStudentBodyValidationSchema

	if err := c.ShouldBindJSON(&body); err != nil {
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

	// Receber os dois valores retornados pelo use case
	response, err := ctrl.useCase.Execute(c.Request.Context(), usecases.AuthenticateStudentRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Retornar o token gerado no response
	c.JSON(http.StatusOK, gin.H{
		"access_token": response.AccessToken,
	})
}
