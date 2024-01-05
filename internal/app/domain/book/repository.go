package book

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, b *models.Book) error
	Update(ctx context.Context, b *models.Book) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
}
