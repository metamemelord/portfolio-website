package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RedirectionItem struct {
	ID           primitive.ObjectID                `bson:"_id" json:"-"`
	RoutingKey   string                            `bson:"routing_key" json:"routing_key"`
	Target       string                            `bson:"target" json:"target"`
	ExpiryString string                            `bson:"-" json:"expiry"`
	Expiry       time.Time                         `bson:"expiry" json:"-"`
	Active       bool                              `bson:"active" json:"-"`
	Permanent    bool                              `bson:"permanent" json:"-"`
	Metadata     map[string]map[string]interface{} `bson:"metadata" json:"metadata"`
	ForwardPath  *bool                             `bson:"forward_path" json:"forward_path"`
	HitCount     uint64                            `bson:"hit_count"`
}
