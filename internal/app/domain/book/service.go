package book

import (
	"context"
	"fmt"

	"github.com/literalog/library/internal/app/domain/author"
	"github.com/literalog/library/internal/app/domain/genre"
	"github.com/literalog/library/internal/app/domain/series"
	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, b *models.Book) error
	Update(ctx context.Context, b *models.Book) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
}

type service struct {
	repository    Repository
	authorService author.Service
	seriesService series.Service
	genreService  genre.Service
	validator     Validator
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, b *models.Book) error {
	_, err := s.authorService.GetById(ctx, b.AuthorId)
	if err != nil {
		return err
	}

	_, err = s.seriesService.GetById(ctx, b.SeriesId)
	if err != nil {
		return err
	}

	for _, genre := range b.Genre {
		_, err = s.genreService.GetByName(ctx, genre)
		if err != nil {
			return fmt.Errorf("error getting genre %s: %w", genre, err)
		}
	}

	if err := s.validator.Validate(b); err != nil {
		return err
	}
	return s.repository.Create(ctx, b)
}

func (s *service) Update(ctx context.Context, b *models.Book) error {
	_, err := s.authorService.GetById(ctx, b.AuthorId)
	if err != nil {
		return err
	}

	_, err = s.seriesService.GetById(ctx, b.SeriesId)
	if err != nil {
		return err
	}

	for _, genre := range b.Genre {
		_, err = s.genreService.GetByName(ctx, genre)
		if err != nil {
			return fmt.Errorf("error getting genre %s: %w", genre, err)
		}
	}

	if err := s.validator.Validate(b); err != nil {
		return err
	}

	return s.repository.Update(ctx, b)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyId
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Book, error) {
	if id == "" {
		return nil, ErrEmptyId
	}
	return s.repository.GetById(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]models.Book, error) {
	return s.repository.GetAll(ctx)
}
