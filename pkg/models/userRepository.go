package models

import (
	"gorm.io/gorm"
)

type User struct {
	UserID   string `gorm:"column:user_id;primaryKey"`
	Password string
}

type Level struct {
	UserID string `gorm:"column:user_id;primaryKey"`
	Level  int
}

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

func GetUser(db *gorm.DB, userId string) (User, error) {
	var user User
	result := db.First(&user, "user_id = ?", userId)
	return user, result.Error
}
