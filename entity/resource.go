package entity

import (
	
	"labix.org/v2/mgo/bson"
)


type Resource struct {
	ID		bson.ObjectId		`json:"id" bson:"_id"`
	Name		string		`json:"description" bson:"description"`
	Description		string		`json:"name" bson:"name"`
}