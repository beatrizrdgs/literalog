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

func (r *SeriesRepository) Create(ctx context.Context, series *models.Series) error {
	_, err := r.collection.InsertOne(ctx, series)
	if err != nil {
		return fmt.Errorf("error creating series: %w", err)
	}
	return nil
}

func (r *SeriesRepository) Update(ctx context.Context, series *models.Series) error {
	filter := bson.M{"_id": series.ID}
	update := bson.M{"$set": series}
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

func (r *SeriesRepository) GetByID(ctx context.Context, id string) (*models.Series, error) {
	filter := bson.M{"_id": id}
	series := new(models.Series)
	if err := r.collection.FindOne(ctx, filter).Decode(series); err != nil {
		return nil, fmt.Errorf("error getting series: %w", err)
	}
	return series, nil
}

func (r *SeriesRepository) GetByName(ctx context.Context, name string) (*models.Series, error) {
	filter := bson.M{"name": name}
	series := new(models.Series)
	if err := r.collection.FindOne(ctx, filter).Decode(series); err != nil {
		return nil, fmt.Errorf("error getting series: %w", err)
	}
	return series, nil
}

func (r *SeriesRepository) GetAll(ctx context.Context) ([]models.Series, error) {
	series := make([]models.Series, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting series: %w", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &series); err != nil {
		return nil, fmt.Errorf("error getting series: %w", err)
	}

	return series, nil
}
