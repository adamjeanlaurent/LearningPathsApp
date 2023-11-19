package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type mySQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func getMySQLConnectionString(config mySQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)
}

func ConnectAndSetup() (*gorm.DB, error) {
	mysqlConfig := mySQLConfig{
		Username: "root",
		Password: "root1234",
		Host:     "localhost",
		Port:     "3306",
		Database: "LearningPathsApp",
	}

	var db *gorm.DB
	var dbConnectionError error

	db, dbConnectionError = gorm.Open("mysql", getMySQLConnectionString(mysqlConfig))

	return db, dbConnectionError
}
