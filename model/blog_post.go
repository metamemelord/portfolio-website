package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogPost struct {
	Author        string             `bson:"author" json:"author,omitempty"`
	AuthorContact string             `bson:"author_contact" json:"author_contact,omitempty"`
	Content       string             `bson:"content" json:"content" binding:"required"`
	Date          string             `bson:"date" json:"date" binding:"required"`
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	Subtitle      string             `bson:"subtitle" json:"subtitle,omitempty"`
	Tags          []string           `bson:"tags" json:"tags,omitempty"`
	Title         string             `bson:"title" json:"title" binding:"required"`
	Visible       bool               `bson:"visible" json:"visible" binding:"required"`
}
