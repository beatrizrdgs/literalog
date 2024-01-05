package mongodb

import (
	"context"
	"fmt"

	"github.com/literalog/library/internal/app/domain/book"
	"github.com/literalog/library/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	collection *mongo.Collection
}

func NewBookRepository(collection *mongo.Collection) book.Repository {
	return &BookRepository{
		collection: collection,
	}
}

func (r *BookRepository) Create(ctx context.Context, b *models.Book) error {
	_, err := r.collection.InsertOne(ctx, b)
	if err != nil {
		return fmt.Errorf("error creating book: %w", err)
	}
	return nil
}

func (r *BookRepository) Update(ctx context.Context, b *models.Book) error {
	filter := bson.M{"_id": b.Id}
	update := bson.M{"$set": b}
	if _, err := r.collection.UpdateOne(ctx, filter, update); err != nil {
		return fmt.Errorf("error updating book: %w", err)
	}
	return nil
}

func (r *BookRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	if _, err := r.collection.DeleteOne(ctx, filter); err != nil {
		return fmt.Errorf("error deleting book: %w", err)
	}
	return nil
}

func (r *BookRepository) GetById(ctx context.Context, id string) (*models.Book, error) {
	filter := bson.M{"_id": id}
	b := new(models.Book)
	if err := r.collection.FindOne(ctx, filter).Decode(b); err != nil {
		return nil, fmt.Errorf("error getting book: %w", err)
	}
	return b, nil
}

func (r *BookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	bb := make([]models.Book, 0)
	cur, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error getting book: %w", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &bb); err != nil {
		return nil, fmt.Errorf("error getting books: %w", err)
	}

	return bb, nil
}
