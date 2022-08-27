package main

import (
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/services"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/handler"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/repositories"
	GRPCServices "github.com/ce-final-project/backend_rest_api/account_service/proto/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	_ "time/tzdata"
)

func main() {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=admin password=test dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	accRepo := repositories.NewAccountRepository(db)
	accSrv := services.NewAccountService(accRepo)
	grpcHandler := handler.NewGRPCHandler(accSrv)

	s := grpc.NewServer()

	var listener net.Listener
	listener, err = net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}

	GRPCServices.RegisterAccountServiceServer(s, grpcHandler)

	log.Println("Starting Account Service Grpc server port :5050")
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
