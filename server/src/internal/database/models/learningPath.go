package models

type LearningPath struct {
	BaseModel
	UserID uint
	Title  string
	Stops  []LearningPathStop
}
