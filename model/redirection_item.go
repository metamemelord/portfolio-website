package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RedirectionItem struct {
	ID           *primitive.ObjectID               `bson:"_id" json:"_id,omitempty"`
	RoutingKey   string                            `bson:"routing_key" json:"routing_key"`
	Target       string                            `bson:"target" json:"target"`
	ExpiryString string                            `bson:"-" json:"expiry"`
	Expiry       *time.Time                        `bson:"expiry" json:"-"`
	Active       *bool                             `bson:"active" json:"active,omitempty"`
	Permanent    *bool                             `bson:"permanent" json:"permanent,omitempty"`
	Metadata     map[string]map[string]interface{} `bson:"metadata,omitempty" json:"metadata,omitempty"`
	ForwardPath  *bool                             `bson:"forward_path" json:"forward_path"`
	HitCount     uint64                            `bson:"hit_count" json:"hit_count,omitempty"`
	CreatedAt    *time.Time                        `bson:"created_at,omitempty" json:"created_at,omitempty"`
}

type RedirectionItemSearchFilter struct {
	Key   string
	Value interface{}
}
