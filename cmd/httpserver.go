package main

import (
	"database/sql"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/consumer"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/handlers"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/repository"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	pg_dsn := os.Getenv("PG_DSN")
	if pg_dsn == "" {
		logger.Error("PG_DSN environment variable not set")
	}
	db, err := sql.Open("postgres", pg_dsn)
	defer db.Close()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	repo := repository.New(db)
	defaultAPIService := handlers.NewPotionAPIServer(repo)
	defaultAPIController := openapi.NewDefaultAPIController(defaultAPIService)

	router := openapi.NewRouter(defaultAPIController)

	host_port := os.Getenv("HOST_PORT")
	if host_port == "" {
		logger.Error("HOST_PORT environment variable not set")
	}

	// consumer start
	//ConsumerRecipreStart(brokers, versionInit, group, topics, assignor string, api RecipesSaver)
	brokers := os.Getenv("brokers")
	versionInit := os.Getenv("version")
	group := os.Getenv("group")
	topics := os.Getenv("topics")
	assignor := os.Getenv("assignor")

	// consumer start
	go consumer.ConsumerRecipreStart(brokers, versionInit, group, topics, assignor, defaultAPIService)

	logger.Info("Starting server on server port", host_port)

	log.Fatal(http.ListenAndServe(host_port, router))
}
