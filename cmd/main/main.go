package main

import (
	"context"
	"errors"
	"github.com/DeMarDeXis/VProj/internal/config"
	"github.com/DeMarDeXis/VProj/internal/httpHandler/handler"
	"github.com/DeMarDeXis/VProj/internal/lib/logger/handler/slogpretty"
	"github.com/DeMarDeXis/VProj/internal/service"
	storage "github.com/DeMarDeXis/VProj/internal/storage"
	"github.com/DeMarDeXis/VProj/internal/storage/postgres"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	cfg := config.InitConfig()

	logg := setupPrettySlogLocal()

	logg.Info("starting Diplom", slog.String("env", cfg.Env))

	db, err := postgres.New(postgres.StorageConfig{
		Host:     cfg.StorageConfig.Host,
		Port:     cfg.StorageConfig.Port,
		Username: cfg.StorageConfig.Username,
		Password: cfg.StorageConfig.Password,
		DBName:   cfg.StorageConfig.DBName,
		SSLMode:  cfg.StorageConfig.SSLMode,
	}, logg)
	if err != nil {
		logg.Error("failed to init db", slog.Any("error", err))
		os.Exit(1)
	}

	storageInit := storage.NewStorage(db)
	logg.Info("db connected")

	services := service.NewService(storageInit)
	logg.Info("service created")

	handlers := handler.NewHandler(services, logg)
	logg.Info("handler created")

	srv := http.Server{
		Addr:         cfg.Address + ":" + strconv.Itoa(cfg.HTTPServer.Port),
		Handler:      handlers.InitRoutes(logg),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	go func() {
		logg.Info("server is listening", slog.String("address", cfg.Address+":"+strconv.Itoa(cfg.HTTPServer.Port)))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("failed to start server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logg.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logg.Error("server forced to shutdown", err)
	}

	if err := db.Close(); err != nil {
		logg.Error("failed to close db connection", err)
	} else {
		logg.Info("db connection closed")
	}

	logg.Info("server exiting")
}

func setupPrettySlogLocal() *slog.Logger {
	opts := slogpretty.PrettyHandlersOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handlerLog := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handlerLog)
}

// TODO: \internal\model\nhl.go (16.02.25)
// TODO: get information about teams -> method (11.05.25)
// TODO: refactor storage query for GetLastSchedule, it must be corrected. Must output the next 10 games after already played games. (11.05.25)
// TODO: make method for getting NBA teams (11.05.25)
// TODO: Necesito hacer un Tableau para equipos deportivos favoritos de los usuarios (11.05.25)
