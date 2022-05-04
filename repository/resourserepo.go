package repository

import (
	"context"
	"log"

	"github.com/pedraamy/gin-api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ResourceRepo interface {
	AddAws(resource entity.Resource) (*mongo.InsertOneResult, error)
	AddAzure(resource entity.Resource) (*mongo.InsertOneResult, error)
	AddGcp(resource entity.Resource) (*mongo.InsertOneResult, error)
	/* DeleteAws(resource entity.Resource) error
	DeleteAzure(resource entity.Resource) error
	DeleteGcp(resource entity.Resource) error */
	GetAllAws() ([]bson.D, error)
	GetAllAzure() ([]bson.D, error)
	GetAllGcp() ([]bson.D, error)
	Close()
}

type database struct {
	client *mongo.Client
	ctx context.Context
	name string
}

func NewResourceRepo(name string) ResourceRepo {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mongo:mongo@database-v1.myknd.mongodb.net/test"))
	if err != nil {
		log.Fatal("lol")
	}
	ctx := context.Background()
	client.Connect(ctx)
	/* db := client.Database("Test-v1")
	aws := db.Collection("AWS")
	azure := db.Collection("Azure")
	gcp := db.Collection("GCP") */
	return &database{client: client, ctx: ctx, name: name}
}


func (db *database) Close() {
	err := db.client.Disconnect(db.ctx)
	if err != nil {
		panic("Failed to disconnect from database!")
	}
}

func (db *database) AddAws(resource entity.Resource) (*mongo.InsertOneResult, error) {
	res, err := db.AddCollection("AWS", resource)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (db *database) AddAzure(resource entity.Resource) (*mongo.InsertOneResult, error) {
	res, err := db.AddCollection("Azure", resource)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (db *database) AddGcp(resource entity.Resource) (*mongo.InsertOneResult, error) {
	res, err := db.AddCollection("GCP", resource)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/* func (db *database) DeleteAws(resource entity.Resource) {
	db.connection.Delete(&resource)
}
func (db *database) DeleteAzure(resource entity.AzureResource) {
	db.connection.Delete(&resource)
}
func (db *database) DeleteGcp(resource entity.GcpResource) {
	db.connection.Delete(&resource)
} */

func (db *database) GetAllAws() ([]bson.D, error) {
	cur, err := db.GetAllCollection("AWS")
	if err != nil {
		return nil, err
	}
	var results []bson.D
	err = cur.All(db.ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *database) GetAllAzure() ([]bson.D, error) {
	cur, err := db.GetAllCollection("Azure")
	if err != nil {
		return nil, err
	}
	var results []bson.D
	err = cur.All(db.ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *database) GetAllGcp() ([]bson.D, error) {
	cur, err := db.GetAllCollection("GCP")
	if err != nil {
		return nil, err
	}
	var results []bson.D
	err = cur.All(db.ctx, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (db *database) GetAllCollection(collection string) (*mongo.Cursor, error) {
	cur, err := db.client.Database(db.name).Collection(collection).Find(db.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	return cur, err
}

func (db *database) AddCollection(collection string, resource entity.Resource) (*mongo.InsertOneResult, error) {
	res, err := db.client.Database(db.name).Collection(collection).InsertOne(db.ctx, resource)
	if err != nil {
		return nil, err
	}
	return res, err
} 