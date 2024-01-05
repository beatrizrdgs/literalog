package series

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, s *models.Series) error
	Update(ctx context.Context, s *models.Series) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Series, error)
	GetAll(ctx context.Context) ([]models.Series, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, series *models.Series) error {
	return s.repository.Create(ctx, series)
}

func (s *service) Update(ctx context.Context, series *models.Series) error {
	return s.repository.Update(ctx, series)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyId
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Series, error) {
	if id == "" {
		return nil, ErrEmptyId
	}
	return s.repository.GetById(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]models.Series, error) {
	return s.repository.GetAll(ctx)
}
