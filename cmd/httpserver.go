package main

import (
	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/handlers"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	DefaultAPIService := handlers.NewPotionAPIServer()
	DefaultAPIController := openapi.NewDefaultAPIController(DefaultAPIService)

	router := openapi.NewRouter(DefaultAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
