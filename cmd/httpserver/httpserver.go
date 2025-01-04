package main

import (
	"context"
	"github.com/ZhdanovichVlad/potion-making-service/branches/generated/openapi"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/processor"
	"github.com/ZhdanovichVlad/potion-making-service/branches/internal/repository"
	"github.com/jackc/pgx/v4"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	exitCodeError = 1
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	pg_dsn := os.Getenv("PG_DSN")
	if pg_dsn == "" {
		logger.Error("PG_DSN environment variable not set")
		log.Fatal("PG_DSN environment variable not set")
	}

	ctxDbOpen, cancelDbOpen := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelDbOpen()

	db, err := pgx.Connect(ctxDbOpen, pg_dsn)
	if err != nil {
		logger.Error("error opening database: %v", err)
		log.Fatal("error opening database: %v", err)
	}

	recipesRepository := repository.NewRecipesStorage(db)
	recipesProcessor := processor.NewRecipesAPIServer(recipesRepository)
	recipesService := openapi.NewRecipeAPIController(recipesProcessor)

	ingredientsRepository := repository.NewIngredientsStorage(db)
	ingredientsProcessor := processor.NewIngredientAPIServer(ingredientsRepository)
	ingredientAPIController := openapi.NewIngredientAPIController(ingredientsProcessor)

	host_port := os.Getenv("HOST_PORT")
	if host_port == "" {
		logger.Error("HOST_PORT environment variable not set")
		log.Fatal("HOST_PORT environment variable not set")
	}

	router := openapi.NewRouter(ingredientAPIController, recipesService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("starting server on server port", host_port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listenAndServer error :%v", err)
		}

	}()

	<-stop
	logger.Info("shutting down server")

	ctxServer, cancelServer := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelServer()

	if err := server.Shutdown(ctxServer); err != nil {
		logger.Error("server forced to shutdown: %v", err)
	}

	logger.Info("server exited gracefully")

	ctxDbClose, cancelDbClose := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelDbClose()

	logger.Info("closing database connection")
	db.Close(ctxDbClose)

	logger.Info("app exited gracefully")

}
