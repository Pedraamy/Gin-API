package repository

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pedraamy/gin-api/entity"
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
	connection *gorm.DB
}

func NewResourceRepo() ResourceRepo {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.AwsResource{}, &entity.AzureResource{}, &entity.GcpResource{})
	return &database{
		connection: db,
	}
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

func (db *database) GetAllAws() []entity.AwsResource {
	var resources []entity.AwsResource
	db.connection.Set("gorm:auto_preload", true).Find(&resources)
	return resources
}

func (db *database) GetAllAzure() []entity.AzureResource {
	var resources []entity.AzureResource
	db.connection.Set("gorm:auto_preload", true).Find(&resources)
	return resources
}

func (db *database) GetAllGcp() []entity.GcpResource {
	var resources []entity.GcpResource
	db.connection.Set("gorm:auto_preload", true).Find(&resources)
	return resources
}
