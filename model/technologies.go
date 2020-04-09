package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Technology struct {
	ID        primitive.ObjectID     `bson:"_id" json:"id"`
	Name      string                 `bson:"name" json:"name"`                     // Go
	Type      string                 `bson:"type" json:"type"`                     // Programming Language
	MediaType string                 `bson:"media_type" json:"media_type"`         // Image
	URL       string                 `bson:"url" json:"url"`                       // URL
	Src       string                 `bson:"src" json:"src,omitempty"`             // <URL>
	CSSClass  string                 `bson:"css_class" json:"css_class,omitempty"` // Empty
	Metadata  map[string]interface{} `bson:"metadata" json:"metadata,omitempty"`   // nil
}
