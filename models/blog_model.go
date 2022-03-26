package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	Id          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Datetime    string             `json:"datetime"`
	Description string             `json:"description"`
	Content     string             `json:"content"`
}
