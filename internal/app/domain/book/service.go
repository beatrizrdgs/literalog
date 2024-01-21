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
	Create(ctx context.Context, book *models.Book) error
	Update(ctx context.Context, book *models.Book) error
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

func (s *service) Create(ctx context.Context, book *models.Book) error {
	_, err := s.authorService.GetById(ctx, book.AuthorId)
	if err != nil {
		return err
	}

	_, err = s.seriesService.GetById(ctx, book.SeriesId)
	if err != nil {
		return err
	}

	for _, genre := range book.Genre {
		_, err = s.genreService.GetByName(ctx, genre)
		if err != nil {
			return err
		}
	}

	if err := s.validator.Validate(book); err != nil {
		return err
	}
	return s.repository.Create(ctx, book)
}

func (s *service) Update(ctx context.Context, book *models.Book) error {
	_, err := s.authorService.GetById(ctx, book.AuthorId)
	if err != nil {
		return err
	}

	_, err = s.seriesService.GetById(ctx, book.SeriesId)
	if err != nil {
		return fmt.Errorf("error getting series: %w", err)
	}

	for _, genre := range book.Genre {
		_, err = s.genreService.GetByName(ctx, genre)
		if err != nil {
			return err
		}
	}

	if err := s.validator.Validate(book); err != nil {
		return err
	}

	return s.repository.Update(ctx, book)
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

	book, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := s.validator.Validate(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *service) GetAll(ctx context.Context) ([]models.Book, error) {
	books, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		if err := s.validator.Validate(&book); err != nil {
			return nil, err
		}
	}

	return books, nil
}
