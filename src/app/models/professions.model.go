package models

import (
	"time"

	"gorm.io/gorm"
)

type Profession struct {
	Id        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"size:150;" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
