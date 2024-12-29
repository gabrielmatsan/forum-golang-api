package usecases

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/repositories"
	usecaseserror "github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/application/use-cases-error"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/criptography"
	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
)

type RegisterStudentRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Resposta do caso de uso: ou um erro ou o estudante criado

type CreateStudentUseCase struct {
	studentRepository repositories.StudentsRepository
	hashGenerator     criptography.HashGenerator
}

func NewRegisterStudentUseCase(studentRepository repositories.StudentsRepository, hashGenerator criptography.HashGenerator) *CreateStudentUseCase {
	return &CreateStudentUseCase{
		studentRepository: studentRepository,
		hashGenerator:     hashGenerator,
	}
}

func (uc *CreateStudentUseCase) Execute(ctx context.Context, req RegisterStudentRequest) error {
	log.Printf("Request received: %+v", req)

	existingStudent, err := uc.studentRepository.FindByEmail(ctx, req.Email)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("Error checking existing student: %v", err)
		return fmt.Errorf("failed to check existing student: %w", err)
	}

	// verifica se a senha é válida
	if len(req.Password) < 6 {
		return usecaseserror.NewWeakPasswordError()
	}

	if existingStudent != nil {
		log.Printf("Student already exists: %+v", existingStudent)
		return usecaseserror.NewEmailAlreadyUsedError(req.Email)
	}

	log.Printf("Creating new student")

	hashedPassword, err := uc.hashGenerator.Hash(req.Password)

	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return fmt.Errorf("failed to hash password: %w", err)
	}

	newStudent := models.NewStudent(models.StudentProps{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})

	if err := uc.studentRepository.CreateStudent(ctx, newStudent); err != nil {
		log.Printf("Failed to save student: %v", err)
		return fmt.Errorf("failed to save student: %w", err)
	}

	log.Printf("Student created successfully")
	return nil
}
