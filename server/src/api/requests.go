package api

type CreateAccountRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginToAccountRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateLearningPathRequestBody struct {
	Title string `json:"title"`
}

type CreateLearningPathStopRequestBody struct {
	Title          string `json:"title"`
	LearningPathID uint   `json:"learningPathId"`
	MarkdownBody   string `json:"body"`
	Stop           uint   `json:"stopNumber"`
}
