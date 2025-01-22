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

	pgDSN := os.Getenv("PG_DSN")
	if pgDSN == "" {
		logger.Error("PG_DSN environment variable not set")
		os.Exit(1)
	}

	ctxDbOpen, cancelDbOpen := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelDbOpen()

	db, err := pgx.Connect(ctxDbOpen, pgDSN)
	if err != nil {
		logger.Error("error opening database: %v", err)
		os.Exit(1)
	}

	ctxDbClose, cancelDbClose := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelDbClose()
	defer db.Close(ctxDbClose)

	recipesRepositories := repository.NewRecipesStorage(db)
	recipesProcessors := processor.NewRecipesAPIServer(recipesRepositories)
	recipesService := openapi.NewRecipeAPIController(recipesProcessors)

	ingredientsRepositories := repository.NewIngredientsStorage(db)
	ingredientsProcessors := processor.NewIngredientAPIServer(ingredientsRepositories)
	ingredientAPIController := openapi.NewIngredientAPIController(ingredientsProcessors)

	hostPort := os.Getenv("HOST_PORT")
	if hostPort == "" {
		logger.Error("HOST_PORT environment variable not set")
		log.Fatal("HOST_PORT environment variable not set")
	}

	router := openapi.NewRouter(ingredientAPIController, recipesService)

	server := &http.Server{
		Addr:    hostPort,
		Handler: router,
	}

	//stop := make(chan os.Signal, 1)
	//signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	ctxEnd, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info("starting server on server", slog.String("port", hostPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("listenAndServer error :%v", err)
		}

	}()

	<-ctxEnd.Done()
	logger.Info("shutting down server")

	ctxServer, cancelServer := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelServer()

	if err := server.Shutdown(ctxServer); err != nil {
		logger.Error("server forced to shutdown: %v", err)
	}

	logger.Info("server exited gracefully")

	logger.Info("app exited gracefully")
}
