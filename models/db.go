package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Datastore is the interface which should be implemented by DB.
type Datastore interface {
	AllEvents() ([]EventDocument, error)
	AddEvent(*NewEvent) error
}

// DB is a general abstraction for your database client
type DB struct {
	*mongo.Client
}

// NewDB creates a new mongodb database from the provided dbURI
func NewDB(dbURI string) (*DB, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	return &DB{client}, nil
}
