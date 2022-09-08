package main

import (
	"flag"
	"github.com/ce-final-project/backend_rest_api/account_service/config"
	"github.com/ce-final-project/backend_rest_api/account_service/internal/server"
	"github.com/ce-final-project/backend_rest_api/pkg/logger"
	"log"
	_ "time/tzdata"
)

func main() {

	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("AccountService")

	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run())
}
