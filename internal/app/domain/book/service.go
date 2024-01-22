package book

import (
	"context"

	"github.com/literalog/library/internal/app/domain/authors"
	"github.com/literalog/library/internal/app/domain/genre"
	"github.com/literalog/library/internal/app/domain/series"
	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, book *models.Book) error
	CreateByISBN(ctx context.Context, isbn string) error
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
}

type service struct {
	repository     Repository
	isbnRepository ISBNRepository
	authorService  authors.Service
	seriesService  series.Service
	genreService   genre.Service
	validator      Validator
}

func NewService(repo Repository, isbnRepo ISBNRepository, authorSvc authors.Service, seriesSvc series.Service, genreSvc genre.Service) Service {
	return &service{
		repository:     repo,
		isbnRepository: isbnRepo,
		authorService:  authorSvc,
		seriesService:  seriesSvc,
		genreService:   genreSvc,
		validator:      Validator{},
	}
}

func (s *service) Create(ctx context.Context, book *models.Book) error {
	if book.AuthorIDs == nil || len(book.AuthorIDs) == 0 {
		return ErrEmptyAuthors
	}

	for _, author := range book.AuthorIDs {
		_, err := s.authorService.GetByID(ctx, author)
		if err != nil {
			return err
		}
	}

	if book.SeriesID != "" {
		_, err := s.seriesService.GetByID(ctx, book.SeriesID)
		if err != nil {
			return err
		}
	}

	if book.Genre != nil {
		for _, genre := range book.Genre {
			_, err := s.genreService.GetByName(ctx, genre)
			if err != nil {
				return err
			}
		}
	}

	if err := s.validator.Validate(book); err != nil {
		return err
	}
	return s.repository.Create(ctx, book)
}

func (s *service) CreateByISBN(ctx context.Context, isbn string) error {
	if isbn == "" {
		return ErrEmptyISBN
	}

	book, err := s.isbnRepository.Get(ctx, isbn)
	if err != nil {
		return err
	}

	if err := s.validator.Validate(book); err != nil {
		return err
	}

	for i := range book.AuthorIDs {
		author, err := s.authorService.GetByName(ctx, book.AuthorIDs[i])
		switch {
		case err == authors.ErrNotFound:
			author = models.NewAuthor(book.AuthorIDs[i])
			if err := s.authorService.Create(ctx, author); err != nil {
				return err
			}
		case err != nil:
			return err
		}

		book.AuthorIDs[i] = author.ID
	}

	return s.Create(ctx, book)
}

func (s *service) Update(ctx context.Context, book *models.Book) error {
	for _, author := range book.AuthorIDs {
		_, err := s.authorService.GetByID(ctx, author)
		if err != nil {
			return err
		}
	}

	_, err := s.seriesService.GetByID(ctx, book.SeriesID)
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

	return s.repository.Update(ctx, book)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyID
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetByID(ctx context.Context, id string) (*models.Book, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	book, err := s.repository.GetByID(ctx, id)
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
