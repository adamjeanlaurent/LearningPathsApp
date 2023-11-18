package models

type LearningPathStop struct {
	BaseModel
	LearningPathID uint
	Title          string
	MarkdownBody   string
	StopNumber     string
}
