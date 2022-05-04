package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Response struct {
	Message string `json:"message"`
}

type Success struct {
	ID primitive.ObjectID `json:"ID"`
}
