package storage

import (
	"time"

	"github.com/google/uuid"
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

func (model *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	return beforeCreate(scope)
}

func (model *BaseModel) BeforeUpdated(scope *gorm.Scope) error {
	return beforeUpdated(scope)
}

func NewBaseModel() *BaseModel {
	return &BaseModel{
		StableId: uuid.New().String(),
	}
}

type User struct {
	BaseModel
	Email         string `gorm:"not null"`
	Hash          string `gorm:"not null"`
	LearningPaths []LearningPath
}

type LearningPathStop struct {
	BaseModel
	LearningPathID uint
	Title          string
	MarkdownBody   string
	StopNumber     uint
}

func NewLearningPathStop() *LearningPathStop {
	return &LearningPathStop{
		BaseModel: *NewBaseModel(),
	}
}

type LearningPath struct {
	BaseModel
	UserID uint
	Title  string
	Stops  []LearningPathStop
}
