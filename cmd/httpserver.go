package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/handlers"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/repository"
)

func main() {
	log.Printf("Server started")

	PG_DSN := os.Getenv("PG_DSN")
	if PG_DSN == "" {
		log.Fatal("PG_DSN environment variable not set")
	}
	db, err := sql.Open("postgres", PG_DSN)
	defer db.Close()
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	repo := repository.New(db)
	DefaultAPIService := handlers.NewPotionAPIServer(repo)
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)

	router := openapi.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
