package routes

import (
	"log"
	"net/http"

	"github.com/adamjeanlaurent/LearningPathsApp/internal/database"

	"github.com/jinzhu/gorm"
)

type RequestHandlerClient struct {
	DB *gorm.DB
}

func configureRoutes(handler *RequestHandlerClient) {
	http.HandleFunc("/createAccount", handler.CreateAccountRequestHandler)
}

func RunServer() {
	// connect to db
	var db *gorm.DB
	db, err := database.ConnectAndSetup()

	// check for db connection errors
	if err != nil {
		log.Fatal(err)
	}

	if db == nil {
		log.Fatal("Nil database connection")
	}

	handler := &RequestHandlerClient{DB: db}
	configureRoutes(handler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
