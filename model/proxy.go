package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProxyItem struct {
	ID           primitive.ObjectID `bson:"_id" json:"-"`
	RoutingKey   string             `json:"routing_key" bson:"routing_key"`
	Target       string             `json:"target" bson:"target"`
	ExpiryString string             `json:"expiry" bson:"-"`
	Expiry       time.Time          `json:"-" bson:"expiry"`
	Active       bool               `json:"-" bson:"active"`
	Permanent    bool               `json:"permanent" bson:"permanent"`
}
