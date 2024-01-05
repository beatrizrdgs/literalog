package author

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, a *models.Author) error
	Update(ctx context.Context, a *models.Author) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Author, error)
	GetAll(ctx context.Context) ([]models.Author, error)
}
