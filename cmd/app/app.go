package app

import (
	"context"
	"entain-golang-task/cmd/api/routes"
	"entain-golang-task/database"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type App struct {
	Router *http.ServeMux
	Server *http.Server
	Logger zerolog.Logger
}

func NewApp() *App {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	router := http.NewServeMux()

	database.ConnectToDB(logger)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	routes.DefineRoutes(logger, router)

	return &App{
		Logger: logger,
		Router: router,
		Server: server,
	}
}

func (a *App) Run() {
	go func() {
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Logger.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	a.Logger.Info().Msgf("Server is up and running on %s", a.Server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Logger.Info().Msg("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database.DB.Close()

	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	a.Logger.Info().Msg("Server exiting")
}
