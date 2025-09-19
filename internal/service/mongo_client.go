package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo wraps the official mongo.Client for easier use.
type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

// New creates a new Mongo wrapper and connects to the server.
func NewMongoClient(uri, dbName, user, pass string) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	if user != "" && pass != "" {
		clientOpts.SetAuth(options.Credential{
			Username: user,
			Password: pass,
		})
	}

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Mongo{
		client: client,
		db:     client.Database(dbName),
	}, nil
}

// Collection returns a handle to a specific collection.
func (m *Mongo) Collection(name string) *mongo.Collection {
	return m.db.Collection(name)
}

// Disconnect closes the underlying client connection.
func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
