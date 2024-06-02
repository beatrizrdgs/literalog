package mongodb

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type single struct {
	*mongo.Client
}

var (
	once     sync.Once
	instance *single
)

func GetInstance() (*single, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if instance == nil {
		once.Do(
			func() {
				log.Println("Creating single mongo instance")
				uri := os.Getenv("MONGO_URI")
				client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
				if err != nil {
					log.Fatalf("Error connecting to MongoDB: %v", err)
					return
				}
				instance = &single{client}
			})
	}
	log.Println("Returning single instance")
	return instance, nil
}

func GetDatabase(name string) (*mongo.Database, error) {
	client, err := GetInstance()
	if err != nil {
		return nil, err
	}
	return client.Database(name), nil
}

func GetCollection(name string) (*mongo.Collection, error) {
	db, err := GetDatabase(os.Getenv("MONGO_DB"))
	if err != nil {
		return nil, err
	}
	return db.Collection(name), nil
}
