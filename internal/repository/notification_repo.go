package repository

import (
	"context"
	"time"

	"github.com/quochao170402/notification-service/internal/core"
	"github.com/quochao170402/notification-service/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationRepository interface {
	Create(ctx context.Context, n *core.Notification) (string, error)
	GetByID(ctx context.Context, id string) (*core.Notification, error)
	GetAll(ctx context.Context) ([]core.Notification, error)
	UpdateStatus(ctx context.Context, id string, status core.Status) error
	Delete(ctx context.Context, id string) error
}

type notificationRepo struct {
	mongo *service.Mongo
}

const collection = "notifications"

func (r *notificationRepo) Create(ctx context.Context, n *core.Notification) (string, error) {
	// ensure ID and timestamps
	n.ID = primitive.NewObjectID()
	n.CreatedAt = time.Now()
	if _, err := r.mongo.Collection(collection).InsertOne(ctx, n); err != nil {
		return "", err
	}
	return n.ID.Hex(), nil
}

func (r *notificationRepo) GetByID(ctx context.Context, id string) (*core.Notification, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var notif core.Notification
	if err := r.mongo.Collection(collection).
		FindOne(ctx, bson.M{"_id": oid}).Decode(&notif); err != nil {
		return nil, err
	}
	return &notif, nil
}

func (r *notificationRepo) GetAll(ctx context.Context) ([]core.Notification, error) {
	cur, err := r.mongo.Collection(collection).
		Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var list []core.Notification
	if err := cur.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *notificationRepo) UpdateStatus(ctx context.Context, id string, status core.Status) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"status":  status,
			"sent_at": time.Now(),
		},
	}
	_, err = r.mongo.Collection(collection).UpdateOne(ctx, bson.M{"_id": oid}, update)
	return err
}

func (r *notificationRepo) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = r.mongo.Collection(collection).DeleteOne(ctx, bson.M{"_id": oid})
	return err
}

// NewNotificationRepository creates a new repo with the shared Mongo wrapper.
func NewNotificationRepository(m *service.Mongo) NotificationRepository {
	return &notificationRepo{mongo: m}
}
