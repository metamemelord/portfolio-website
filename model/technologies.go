package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Technology struct {
	ID        primitive.ObjectID     `bson:"_id" json:"id"`
	Name      string                 `bson:"name" json:"name"`
	Order     uint                   `bson:"order" json:"order"`
	Type      string                 `bson:"type" json:"type"`
	MediaType string                 `bson:"media_type" json:"media_type"`
	URL       string                 `bson:"url" json:"url"`
	Src       string                 `bson:"src" json:"src,omitempty"`
	CSSClass  string                 `bson:"css_class" json:"css_class,omitempty"`
	Metadata  map[string]interface{} `bson:"metadata" json:"metadata,omitempty"`
}
