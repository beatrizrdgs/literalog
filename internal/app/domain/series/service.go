package series

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, series *models.Series) error
	Update(ctx context.Context, series *models.Series) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Series, error)
	GetAll(ctx context.Context) ([]models.Series, error)
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

func (s *service) Create(ctx context.Context, series *models.Series) error {
	if err := s.validator.Validate(series); err != nil {
		return err
	}
	return s.repository.Create(ctx, series)
}

func (s *service) Update(ctx context.Context, series *models.Series) error {
	if err := s.validator.Validate(series); err != nil {
		return err
	}
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

	series, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.validator.Validate(series); err != nil {
		return nil, err
	}

	return s.repository.GetById(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]models.Series, error) {
	series, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for i := range series {
		if err := s.validator.Validate(&series[i]); err != nil {
			return nil, err
		}
	}

	return s.repository.GetAll(ctx)
}
