package models

type User struct {
	Id    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"size:150;unique:true" json:"name"`
	Email string `gorm:"size:150;unique:true" json:"email"`
}
