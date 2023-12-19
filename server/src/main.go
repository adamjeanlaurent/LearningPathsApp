package main

import "github.com/adamjeanlaurent/LearningPathsApp/api"

func main() {
	var server *api.ApiServer = api.NewApiServer()
	server.ConnectAndRun()
}
