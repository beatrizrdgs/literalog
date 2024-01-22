package series

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, series *models.Series) error
	Update(ctx context.Context, series *models.Series) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*models.Series, error)
	GetByName(ctx context.Context, name string) (*models.Series, error)
	GetAll(ctx context.Context) ([]models.Series, error)
}
