package genre

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, g *models.Genre) error
	Update(ctx context.Context, g *models.Genre) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Genre, error)
	GetByName(ctx context.Context, name string) (*models.Genre, error)
	GetAll(ctx context.Context) ([]models.Genre, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(ctx context.Context, g *models.Genre) error {
	return s.repository.Create(ctx, g)
}

func (s *service) Update(ctx context.Context, g *models.Genre) error {
	return s.repository.Update(ctx, g)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Genre, error) {
	return s.repository.GetById(ctx, id)
}

func (s *service) GetByName(ctx context.Context, name string) (*models.Genre, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *service) GetAll(ctx context.Context) ([]models.Genre, error) {
	return s.repository.GetAll(ctx)
}
