package entity

import (
)


type Resource struct {
	Name		string		`json:"description" bson:"description"`
	Description		string		`json:"name" bson:"name"`
}