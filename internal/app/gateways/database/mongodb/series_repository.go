package mongodb

import (
	"context"
	"fmt"

	"github.com/literalog/library/internal/app/domain/series"
	"github.com/literalog/library/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SeriesRepository struct {
	collection *mongo.Collection
}

func NewSeriesRepository(collection *mongo.Collection) series.Repository {
	return &SeriesRepository{
		collection: collection,
	}
}

func (r *SeriesRepository) Create(ctx context.Context, s *models.Series) error {
	_, err := r.collection.InsertOne(ctx, s)
	if err != nil {
		return fmt.Errorf("error creating series: %w", err)
	}
	return nil
}

func (r *SeriesRepository) Update(ctx context.Context, s *models.Series) error {
	filter := bson.M{"_id": s.Id}
	update := bson.M{"$set": s}
	if _, err := r.collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("error updating series: %w", err)
	}
	return nil
}

func (r *SeriesRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("error deleting series: %w", err)
	}
	return nil
}

func (r *SeriesRepository) GetById(ctx context.Context, id string) (*models.Series, error) {
	filter := bson.M{"_id": id}
	s := new(models.Series)
	if err := r.collection.FindOne(ctx, filter).Decode(s); err != nil {
		return nil, fmt.Errorf("error getting series: %w", err)
	}
	return s, nil
}

func (r *SeriesRepository) GetAll(ctx context.Context) ([]models.Series, error) {
	ss := make([]models.Series, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting series: %w", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &ss); err != nil {
		return nil, fmt.Errorf("error getting series: %w", err)
	}

	return ss, nil
}
