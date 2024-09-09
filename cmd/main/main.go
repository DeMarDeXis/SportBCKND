package main

import (
	"context"
	"errors"
	"github.com/DeMarDeXis/VProj/internal/config"
	"github.com/DeMarDeXis/VProj/internal/httpHandler/handler"
	"github.com/DeMarDeXis/VProj/internal/httpHandler/service"
	"github.com/DeMarDeXis/VProj/internal/lib/logger/handler/slogpretty"
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
	//TODO: init cfg
	cfg := config.InitConfig()

	//TODO: init logger
	logg := setupPrettySlogLocal()

	logg.Info("starting VProj", slog.String("env", cfg.Env))

	//TODO: init db
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

	storageInit := storage.NewStorage(db, logg)
	logg.Info("db connected")

	//TODO: init service
	services := service.NewService(storageInit)
	logg.Info("service created")

	//TODO: init handlers
	handlers := handler.NewHandler(services, logg)
	logg.Info("handler created")

	//TODO: init server
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

	// TODO: Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logg.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logg.Error("server forced to shutdown", err)
	}

	// TODO: close db connection
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

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
