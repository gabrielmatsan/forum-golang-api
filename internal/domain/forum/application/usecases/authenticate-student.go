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
)

type AuthenticateStudentRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateStudentResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthenticateStudentUseCase struct {
	studentRepo repositories.StudentsRepository
	hashCompare criptography.HashCompare
	encrypter   criptography.Encrypter
}

func NewAuthenticateStudentUseCase(
	studentRepo repositories.StudentsRepository,
	hashCompare criptography.HashCompare,
	encrypter criptography.Encrypter,
) *AuthenticateStudentUseCase {

	return &AuthenticateStudentUseCase{
		studentRepo: studentRepo,
		hashCompare: hashCompare,
		encrypter:   encrypter,
	}
}

func (uc *AuthenticateStudentUseCase) Execute(
	ctx context.Context,
	req AuthenticateStudentRequest,
) (*AuthenticateStudentResponse, error) {

	student, err := uc.studentRepo.FindByEmail(ctx, req.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Error checking existing student: %v", err)
			return nil, usecaseserror.NewWrongCredentialError()
		}
		log.Printf("Error checking existing student: %v", err)
		return nil, fmt.Errorf("failed to check existing student: %w", err)
	}

	if student == nil {
		return nil, usecaseserror.NewWrongCredentialError()
	}

	isMatch, err := uc.hashCompare.Compare(req.Password, student.GetPassword())

	if err != nil {
		log.Printf("Error comparing passwords: %v", err)
		return nil, fmt.Errorf("failed to compare passwords: %w", err)
	}

	if !isMatch {
		return nil, usecaseserror.NewWrongCredentialError()
	}

	payload := map[string]interface{}{
		"id": student.GetID(),
	}

	token, err := uc.encrypter.Encrypt(payload)

	if err != nil {
		log.Printf("Error generating token: %v", err)
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	return &AuthenticateStudentResponse{AccessToken: token}, nil
}
