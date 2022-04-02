package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Datetime    string             `json:"datetime"`
	Description string             `json:"description"`
	Content     string             `json:"content"`
}
