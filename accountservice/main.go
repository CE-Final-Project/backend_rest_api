package main

import (
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountservice/dbclient"
	"github.com/ce-final-project/backend_rest_api/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}

func InitializeMongoClient() {
	service.DBClient = &dbclient.Mongo{}
	service.DBClient.Connect()
}
