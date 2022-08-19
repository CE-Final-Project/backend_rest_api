package main

import (
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountservice/dbclient"
	"github.com/ce-final-project/backend_rest_api/accountservice/service"
)

var appName = "accountService"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeMongoClient()
	service.StartWebServer("6767")
}

func initializeMongoClient() {
	service.DBClient = new(dbclient.Mongo)
	service.DBClient.Seed()
	service.DBClient.Connect()
}
