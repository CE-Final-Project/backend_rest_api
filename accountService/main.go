package main

import (
	"context"
	"fmt"
	"github.com/ce-final-project/backend_rest_api/accountService/api"
	"github.com/ce-final-project/backend_rest_api/accountService/core"
	"github.com/ce-final-project/backend_rest_api/accountService/repository/mongo"
	"github.com/ce-final-project/backend_rest_api/accountService/repository/redis"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
)

var appName = "accountService"

func main() {

	l := hclog.Default()

	repo := chooseRepo()
	service := core.NewAccountService(repo)
	handler := api.NewHandler(service)

	sm := mux.NewRouter()
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/accounts", handler.GetAllAccount)

	getR.HandleFunc("/accounts/{id}", handler.GetSingleAccount)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         httpPort(),                                       // configure the bind address
		Handler:      ch(sm),                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info(appName + ": Starting server on port " + httpPort())

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := s.Shutdown(ctx)
	if err != nil {
		l.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() core.AccountRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := redis.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mongo.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	log.Fatal("Config Database not found!")
	return nil
}
