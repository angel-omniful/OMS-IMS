package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Webhook struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TenantID  string             `bson:"tenant_id" json:"tenant_id"`
	URL       string             `bson:"url" json:"url"`
	Event     string             `bson:"event" json:"event"`       // e.g., "order.updated"
	Active    bool               `bson:"active" json:"active"`     // defaults to true
	Secret    string             `bson:"secret,omitempty" json:"secret,omitempty"` // optional HMAC secret
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

