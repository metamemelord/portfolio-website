package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id"`
	Company          string             `bson:"company" json:"company" binding:"required"`
	Title            string             `bson:"title" json:"title" binding:"required"`
	Responsibilities []string           `bson:"responsibilities,omitempty" json:"responsibilities,omitempty"`
	FromDate         string             `bson:"from_date" json:"from_date" binding:"required"`
	ToDate           string             `bson:"to_date" json:"to_date" binding:"required"`
}

type Occupation struct {
	Title   string `json:"title" bson:"title"`
	Company string `json:"company" bson:"company"`
	Since   string `json:"since" bson:"since"`
}
