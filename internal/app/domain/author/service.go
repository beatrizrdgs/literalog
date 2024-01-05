package author

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, a *models.Author) error
	Update(ctx context.Context, a *models.Author) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Author, error)
	GetAll(ctx context.Context) ([]models.Author, error)
}

type service struct {
	repository Repository
	validator  Validator
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, a *models.Author) error {
	if err := s.validator.Validate(a); err != nil {
		return err
	}
	return s.repository.Create(ctx, a)
}

func (s *service) Update(ctx context.Context, a *models.Author) error {
	return s.repository.Update(ctx, a)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyId
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Author, error) {
	if id == "" {
		return nil, ErrEmptyId
	}
	return s.repository.GetById(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]models.Author, error) {
	return s.repository.GetAll(ctx)
}
