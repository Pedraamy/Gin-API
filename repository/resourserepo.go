package repository

import (
	"context"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pedraamy/gin-api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ResourceRepo interface {
	AddAws(resource entity.AwsResource)
	AddAzure(resource entity.AzureResource)
	AddGcp(resource entity.GcpResource)
	DeleteAws(resource entity.AwsResource)
	DeleteAzure(resource entity.AzureResource)
	DeleteGcp(resource entity.GcpResource)
	GetAllAws() []entity.AwsResource
	GetAllAzure() []entity.AzureResource
	GetAllGcp() []entity.GcpResource
	CloseDB()
}

type database struct {
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
	ctx := context.Background()
	db := client.Database("Test-v1")
	aws := db.Collection("AWS")
	azure := db.Collection("Azure")
	gcp := db.Collection("GCP")
	return &database{ctx, db, aws, azure, gcp}
}


func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) AddAws(resource entity.AwsResource) {
	db.connection.Create(&resource)
}

func (db *database) AddAzure(resource entity.AzureResource) {
	db.connection.Create(&resource)
}

func (db *database) AddGcp(resource entity.GcpResource) {
	db.connection.Create(&resource)
}

func (db *database) DeleteAws(resource entity.AwsResource) {
	db.connection.Delete(&resource)
}
func (db *database) DeleteAzure(resource entity.AzureResource) {
	db.connection.Delete(&resource)
}
func (db *database) DeleteGcp(resource entity.GcpResource) {
	db.connection.Delete(&resource)
}

func (db *database) GetAllAws() []bson.D {
	cur, err := db.aws.Find(db.ctx, bson.D{})
	var results []bson.D
	err = cur.All(db.ctx, &results)
	return results
}

func (db *database) GetAllAzure() []bson.D {
	cur, err := db.aws.Find(db.ctx, bson.D{})
	var results []bson.D
	err = cur.All(db.ctx, &results)
	return results
}

func (db *database) GetAllGcp() []bson.D {
	cur, err := db.aws.Find(db.ctx, bson.D{})
	var results []bson.D
	err = cur.All(db.ctx, &results)
	return results
}
