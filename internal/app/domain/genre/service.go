package genre

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, genre *models.Genre) error
	Update(ctx context.Context, genre *models.Genre) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*models.Genre, error)
	GetByName(ctx context.Context, name string) (*models.Genre, error)
	GetAll(ctx context.Context) ([]models.Genre, error)
}

type service struct {
	repository Repository
	validator  Validator
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(ctx context.Context, genre *models.Genre) error {
	if err := s.validator.Validate(genre); err != nil {
		return err
	}
	return s.repository.Create(ctx, genre)
}

func (s *service) Update(ctx context.Context, genre *models.Genre) error {
	if err := s.validator.Validate(genre); err != nil {
		return err
	}
	return s.repository.Update(ctx, genre)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyID
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetByID(ctx context.Context, id string) (*models.Genre, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	genre, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.validator.Validate(genre); err != nil {
		return nil, err
	}

	return s.repository.GetByID(ctx, id)
}

func (s *service) GetByName(ctx context.Context, name string) (*models.Genre, error) {
	if name == "" {
		return nil, ErrEmptyName
	}

	genre, err := s.repository.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}

	if err := s.validator.Validate(genre); err != nil {
		return nil, err
	}

	return s.repository.GetByName(ctx, name)
}

func (s *service) GetAll(ctx context.Context) ([]models.Genre, error) {
	genre, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for i := range genre {
		if err := s.validator.Validate(&genre[i]); err != nil {
			return nil, err
		}
	}

	return genre, nil
}
