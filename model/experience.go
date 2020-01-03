package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Experience struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Company  string             `bson:"company" json:"company" binding:"required"`
	Title    string             `bson:"title" json:"title" binding:"required"`
	FromDate string             `bson:"from_data" json:"from_date" binding:"required"`
	ToDate   string             `bson:"to_data" json:"to_date" binding:"required"`
}
