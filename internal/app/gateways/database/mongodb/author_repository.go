package mongodb

import (
	"context"
	"fmt"

	"github.com/literalog/library/internal/app/domain/authors"
	"github.com/literalog/library/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorRepository struct {
	collection *mongo.Collection
}

func NewAuthorRepository(collection *mongo.Collection) authors.Repository {
	return &AuthorRepository{
		collection: collection,
	}
}

func (r *AuthorRepository) Create(ctx context.Context, author *models.Author) error {
	_, err := r.collection.InsertOne(ctx, author)
	if err != nil {
		return fmt.Errorf("error creating author: %w", err)
	}
	return nil
}

func (r *AuthorRepository) Update(ctx context.Context, author *models.Author) error {
	filter := bson.M{"_id": author.ID}
	update := bson.M{"$set": author}
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

func (r *AuthorRepository) GetByID(ctx context.Context, id string) (*models.Author, error) {
	filter := bson.M{"_id": id}
	author := new(models.Author)
	if err := r.collection.FindOne(ctx, filter).Decode(author); err != nil {
		return nil, fmt.Errorf("error getting author: %w", err)
	}
	return author, nil
}

func (r *AuthorRepository) GetByName(ctx context.Context, name string) (*models.Author, error) {
	filter := bson.M{"name": name}
	author := new(models.Author)
	if err := r.collection.FindOne(ctx, filter).Decode(author); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, authors.ErrNotFound
		}
		return nil, fmt.Errorf("error getting author: %w", err)
	}
	return author, nil
}

func (r *AuthorRepository) GetAll(ctx context.Context) ([]models.Author, error) {
	authors := make([]models.Author, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting authors: %w", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &authors); err != nil {
		return nil, fmt.Errorf("error getting authors: %w", err)
	}

	return authors, nil
}
