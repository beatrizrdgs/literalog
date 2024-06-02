package services

import (
	"context"
	"errors"

	"github.com/beatrizrdgs/literalog/internal/models"
	"github.com/literalog/go-wise/wise"
)

type BookService struct {
	repo wise.MongoRepository[models.Book]
}

func NewBookService(repo wise.MongoRepository[models.Book]) (*BookService, error) {
	if repo == nil {
		return nil, errors.New("repository is required")
	}
	return &BookService{repo: repo}, nil
}

func (s *BookService) Add(ctx context.Context, book models.Book) error {
	return s.repo.InsertOne(ctx, book)
}

func (s *BookService) GetById(ctx context.Context, id string) (models.Book, error) {
	book, err := s.repo.FindById(ctx, id)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}
