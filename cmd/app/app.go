package app

import (
	"context"
	"entain-golang-task/cmd/api/routes"
	"entain-golang-task/cmd/middleware"
	"entain-golang-task/database"
	"entain-golang-task/pkg/cfg"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type App struct {
	Router *httprouter.Router
	Server *http.Server
	Logger zerolog.Logger
}

func NewApp() *App {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	router := httprouter.New()

	config, err := cfg.LoadConfig()
	if err != nil {
		panic("cannot load env config")
	}

	database.ConnectToDB(logger, config)

	//will be enabled by migrations_enabled flag
	//migrations.MigrateDB()
	wrappedRouter := middleware.ErrorHandlingMiddleware(logger, router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedRouter,
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
		if err := a.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.Logger.Fatal().Err(err).Msg("server failed to start")
		}
	}()

	a.Logger.Info().Msgf("server is up and running on %s", a.Server.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Logger.Info().Msg("server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database.DB.Close()

	if err := a.Server.Shutdown(ctx); err != nil {
		a.Logger.Fatal().Err(err).Msg("server forced to shutdown")
	}

	a.Logger.Info().Msg("server exiting")
}
