package author

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, author *models.Author) error
	Update(ctx context.Context, author *models.Author) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*models.Author, error)
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

func (s *service) Create(ctx context.Context, author *models.Author) error {
	if err := s.validator.Validate(author); err != nil {
		return err
	}
	return s.repository.Create(ctx, author)
}

func (s *service) Update(ctx context.Context, author *models.Author) error {
	if err := s.validator.Validate(author); err != nil {
		return err
	}
	return s.repository.Update(ctx, author)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyID
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetByID(ctx context.Context, id string) (*models.Author, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	author, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.validator.Validate(author); err != nil {
		return nil, err
	}

	return author, nil
}

func (s *service) GetAll(ctx context.Context) ([]models.Author, error) {
	authors, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, author := range authors {
		if err := s.validator.Validate(&author); err != nil {
			return nil, err
		}
	}

	return authors, nil
}
