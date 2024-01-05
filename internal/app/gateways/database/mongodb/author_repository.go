package mongodb

import (
	"context"
	"fmt"

	"github.com/literalog/library/internal/app/domain/author"
	"github.com/literalog/library/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorRepository struct {
	collection *mongo.Collection
}

func NewAuthorRepository(collection *mongo.Collection) author.Repository {
	return &AuthorRepository{
		collection: collection,
	}
}

func (r *AuthorRepository) Create(ctx context.Context, a *models.Author) error {
	_, err := r.collection.InsertOne(ctx, a)
	if err != nil {
		return fmt.Errorf("error creating author: %w", err)
	}
	return nil
}

func (r *AuthorRepository) Update(ctx context.Context, a *models.Author) error {
	filter := bson.M{"_id": a.Id}
	update := bson.M{"$set": a}
	if _, err := r.collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("error updating author: %w", err)
	}
	return nil
}

func (r *AuthorRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("error deleting author: %w", err)
	}
	return nil
}

func (r *AuthorRepository) GetById(ctx context.Context, id string) (*models.Author, error) {
	filter := bson.M{"_id": id}
	a := new(models.Author)
	if err := r.collection.FindOne(ctx, filter).Decode(a); err != nil {
		return nil, fmt.Errorf("error getting author: %w", err)
	}
	return a, nil
}

func (r *AuthorRepository) GetAll(ctx context.Context) ([]models.Author, error) {
	aa := make([]models.Author, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting authors: %w", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &aa); err != nil {
		return nil, fmt.Errorf("error getting authors: %w", err)
	}

	return aa, nil
}
