package inmemoryrepositories

import (
	"context"
	"sync"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
)

type InMemoryStudentsRepository struct {
	mu       sync.Mutex
	students []*models.Student
}

func NewInMemoryStudentsRepository() *InMemoryStudentsRepository {
	return &InMemoryStudentsRepository{
		students: make([]*models.Student, 0),
	}
}

func (r *InMemoryStudentsRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.students = append(r.students, student)
	return nil // Sempre retorna nil neste caso
}

func (r *InMemoryStudentsRepository) FindByEmail(ctx context.Context, email string) (*models.Student, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, s := range r.students {
		if s.GetEmail() == email {
			return s, nil
		}
	}

	return nil, nil
}
