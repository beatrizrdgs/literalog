package authors

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, author *models.Author) error
	Update(ctx context.Context, author *models.Author) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*models.Author, error)
	GetByName(ctx context.Context, name string) (*models.Author, error)
	GetAll(ctx context.Context) ([]models.Author, error)
}
