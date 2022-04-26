package entity

import (
	"time"
)


type AwsResource struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}

type AzureResource struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}

type GcpResource struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}