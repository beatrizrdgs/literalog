package series

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Repository interface {
	Create(ctx context.Context, s *models.Series) error
	Update(ctx context.Context, s *models.Series) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Series, error)
	GetAll(ctx context.Context) ([]models.Series, error)
}
