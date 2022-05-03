package repository

import (
	"context"
	"log"
	"time"

	"github.com/pedraamy/gin-api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ResourceRepo interface {
	AddAws(resource entity.Resource) error
	AddAzure(resource entity.Resource) error
	AddGcp(resource entity.Resource) error
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
	base *mongo.Database
	aws *mongo.Collection
	azure *mongo.Collection
	gcp *mongo.Collection
}

func NewResourceRepo() ResourceRepo {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mongo:mongo@database-v1.myknd.mongodb.net/test"))
	if err != nil {
		log.Fatal("hihi")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	/* db := client.Database("Test-v1")
	aws := db.Collection("AWS")
	azure := db.Collection("Azure")
	gcp := db.Collection("GCP") */
	return &database{client: client, ctx: ctx}
}


func (db *database) Close() {
	err := db.client.Disconnect(db.ctx)
	if err != nil {
		panic("Failed to disconnect from database!")
	}
}

func (db *database) AddAws(resource entity.Resource) error {
	_, err := db.client.Database("Test-v1").Collection("AWS").InsertOne(db.ctx, resource)
	if err != nil {
		return err
	}
	return nil
}

func (db *database) AddAzure(resource entity.Resource) error {
	_, err := db.azure.InsertOne(db.ctx, resource)
	if err != nil {
		return err
	}
	return nil
}

func (db *database) AddGcp(resource entity.Resource) error {
	_, err := db.gcp.InsertOne(db.ctx, resource)
	if err != nil {
		return err
	}
	return nil
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
	cur, err := db.aws.Find(db.ctx, bson.D{})
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
	cur, err := db.aws.Find(db.ctx, bson.D{})
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
	cur, err := db.aws.Find(db.ctx, bson.D{})
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
