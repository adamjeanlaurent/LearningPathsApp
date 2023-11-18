package routes

import (
	"net/http"
)

func (handler *RequestHandlerClient) CreateAccountRequestHandler(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
}
