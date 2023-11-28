package routes

type CreateAccountRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginToAccountRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
