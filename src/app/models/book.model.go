package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string         `gorm:"size:150;unique:true" json:"title" validate:"required"`
	Price       int            `json:"price" validate:"required"`
	Author      string         `gorm:"size:150" json:"author" validate:"required"`
	Publisher   string         `gorm:"size:150" json:"publisher" validate:"required"`
	PublishDate string         `gorm:"type:date" json:"publish_date" validate:"required"`
	OwnerId     int            `json:"owner_id" validate:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
