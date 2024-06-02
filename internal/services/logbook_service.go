package services

import (
	"context"
	"errors"

	"github.com/beatrizrdgs/literalog/internal/models"
	"github.com/literalog/go-wise/wise"
)

type LogbookService struct {
	repo wise.MongoRepository[models.Logbook]
}

func NewLogbookService(repo wise.MongoRepository[models.Logbook]) (*LogbookService, error) {
	if repo == nil {
		return nil, errors.New("repository is required")
	}
	return &LogbookService{repo: repo}, nil
}

func (s *LogbookService) Add(ctx context.Context, logbook models.Logbook) error {
	return s.repo.InsertOne(ctx, logbook)
}

func (s *LogbookService) GetByUserId(ctx context.Context, userId string) ([]models.Logbook, error) {
	filters := map[string][]any{
		"user_id": {userId},
	}
	return s.repo.Find(ctx, filters)
}

func (s *LogbookService) GetByUserIdAndBookId(ctx context.Context, userId, bookId string) (models.Logbook, error) {
	filters := map[string][]any{
		"user_id": {userId},
		"book_id": {bookId},
	}
	return s.repo.FindOne(ctx, filters)
}

func (s *LogbookService) UpdateStatus(ctx context.Context, userId, bookId string, status models.Status) error {
	logbook, err := s.GetByUserIdAndBookId(ctx, userId, bookId)
	if err != nil {
		return err
	}
	logbook.Status = models.NewStatus(status.String())
	return s.Add(ctx, logbook)
}

func (s *LogbookService) DeleteByUserIdAndBookId(ctx context.Context, userId, bookId string) (models.Logbook, error) {
	filters := map[string][]any{
		"user_id": {userId},
		"book_id": {bookId},
	}
	return s.repo.DeleteOne(ctx, filters)
}
