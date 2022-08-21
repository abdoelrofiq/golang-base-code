package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string         `gorm:"size:150;unique:true" json:"name"`
	Email        string         `gorm:"size:150;unique:true" json:"email"`
	ProfessionId int            `json:"profession_id"`
	Profession   Profession     `gorm:"foreignkey:Id;references:ProfessionId"`
	Books        []Book         `gorm:"foreignkey:OwnerId;references:Id"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
