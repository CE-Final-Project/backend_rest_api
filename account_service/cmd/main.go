package main

import (
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/core/services"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/handler"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/repositories"
	"github.com/ce-final-project/backend_rest_api/account_service/pkg/postgres"
	GRPCServices "github.com/ce-final-project/backend_rest_api/account_service/proto/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	_ "time/tzdata"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Initial config error: %v", err)
		return
	}

	var db *sqlx.DB
	db, err = postgres.NewPostgresDB(cfg.Postgresql)
	if err != nil {
		log.Fatalf("Initial PostgresDB error: %v", err)
		return
	}

	accRepo := repositories.NewAccountRepository(db)
	accSrv := services.NewAccountService(accRepo)
	grpcHandler := handler.NewGRPCHandler(accSrv)

	s := grpc.NewServer()

	var listener net.Listener
	listener, err = net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.GRPC.Development {
		log.Println("GRPC Development reflection active")
		reflection.Register(s)
	}

	GRPCServices.RegisterAccountServiceServer(s, grpcHandler)

	log.Printf("Starting Account Service Grpc server port :%s\n", cfg.GRPC.Port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
