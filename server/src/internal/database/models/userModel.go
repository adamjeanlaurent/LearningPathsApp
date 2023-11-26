package models

import "github.com/jinzhu/gorm"

type User struct {
	BaseModel
	Email         string `gorm:"not null"`
	Hash          string `gorm:"not null"`
	LearningPaths []LearningPath
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	return beforeCreate(scope)
}

func (user *User) BeforeUpdated(scope *gorm.Scope) error {
	return beforeUpdated(scope)
}
