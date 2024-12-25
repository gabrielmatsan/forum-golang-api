package inmemoryrepositories

import (
	"context"
	"sync"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
)

type InMemoryStudentsRepository struct {
	Mu       sync.Mutex
	Students []*models.Student
}

func NewInMemoryStudentsRepository() *InMemoryStudentsRepository {
	return &InMemoryStudentsRepository{
		Students: make([]*models.Student, 0),
	}
}

func (r *InMemoryStudentsRepository) CreateStudent(ctx context.Context, student *models.Student) error {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	r.Students = append(r.Students, student)
	return nil // Sempre retorna nil neste caso
}

func (r *InMemoryStudentsRepository) FindByEmail(ctx context.Context, email string) (*models.Student, error) {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	for _, s := range r.Students {
		if s.GetEmail() == email {
			return s, nil
		}
	}

	return nil, nil
}

func (r *InMemoryStudentsRepository) FindById(ctx context.Context, id string) (*models.Student, error) {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	for _, s := range r.Students {
		if s.GetID() == id {
			return s, nil
		}
	}

	return nil, nil
}
