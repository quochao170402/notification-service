package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Notification represents a generic notification document in MongoDB.
type Notification struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Channel    Channel            `bson:"channel"      json:"channel"`
	Recipients []string           `bson:"recipients"   json:"recipients"`
	Title      string             `bson:"title,omitempty"   json:"title,omitempty"`
	Message    string             `bson:"message"     json:"message"`
	Data       map[string]string  `bson:"data,omitempty"     json:"data,omitempty"`
	Metadata   map[string]any     `bson:"metadata,omitempty" json:"metadata,omitempty"`
	Status     Status             `bson:"status"     json:"status"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	SentAt     *time.Time         `bson:"sent_at,omitempty" json:"sent_at,omitempty"`
}
