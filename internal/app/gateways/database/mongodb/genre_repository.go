package mongodb

import (
	"context"
	"fmt"

	"github.com/literalog/library/internal/app/domain/genre"
	"github.com/literalog/library/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GenreRepository struct {
	collection *mongo.Collection
}

func NewGenreRepository(collection *mongo.Collection) genre.Repository {
	return &GenreRepository{
		collection: collection,
	}
}

func (r *GenreRepository) Create(ctx context.Context, g *models.Genre) error {
	_, err := r.collection.InsertOne(ctx, g)
	if err != nil {
		return fmt.Errorf("error creating genre: %w", err)
	}
	return nil
}

func (r *GenreRepository) Update(ctx context.Context, g *models.Genre) error {
	filter := bson.M{"_id": g.Id}
	update := bson.M{"$set": g}
	if _, err := r.collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("error updating genre: %w", err)
	}
	return nil
}

func (r *GenreRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("error deleting genre: %w", err)
	}
	return nil
}

func (r *GenreRepository) GetById(ctx context.Context, id string) (*models.Genre, error) {
	filter := bson.M{"_id": id}
	g := new(models.Genre)
	if err := r.collection.FindOne(ctx, filter).Decode(g); err != nil {
		return nil, fmt.Errorf("error getting genre: %w", err)
	}
	return g, nil
}

func (r *GenreRepository) GetByName(ctx context.Context, name string) (*models.Genre, error) {
	filter := bson.M{"name": name}
	g := new(models.Genre)
	if err := r.collection.FindOne(ctx, filter).Decode(g); err != nil {
		return nil, fmt.Errorf("error getting genre: %w", err)
	}
	return g, nil
}

func (r *GenreRepository) GetAll(ctx context.Context) ([]models.Genre, error) {
	gg := make([]models.Genre, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting genre: %w", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &gg); err != nil {
		return nil, fmt.Errorf("error getting genre: %w", err)
	}

	return gg, nil
}
