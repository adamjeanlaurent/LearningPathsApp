package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

func beforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedDate", time.Now())
	scope.SetColumn("UpdatedDate", time.Now())
	return nil
}

func beforeUpdated(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedDate", time.Now())
	return nil
}

type BaseModel struct {
	ID          uint      `gorm:"primary_key;auto_increment"`
	StableId    string    `gorm:"not null"`
	CreatedDate time.Time `gorm:"not null"`
	UpdatedDate time.Time ``
}
