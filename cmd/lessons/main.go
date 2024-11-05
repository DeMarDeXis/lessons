package main

import (
	"context"
	"courses/internal/config"
	"courses/internal/httphandler"
	"courses/internal/lib/logger/handler/slogpretty"
	"courses/internal/service"
	"courses/internal/storage"
	"courses/internal/storage/postgres"
	"errors"
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

	logg.Info("starting lessons service", slog.String("env", cfg.Env))

	db, err := postgres.New(postgres.StorageConfig{
		Host:     cfg.Host,
		Port:     cfg.StorageConfig.Port,
		Username: cfg.Username,
		Password: cfg.Password,
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
	}, logg)
	if err != nil {
		logg.Error("failed to init db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	storageInit := storage.NewStorage(db, logg)
	logg.Info("storage init", slog.String("storage", "postgres"))

	srvce := service.NewService(storageInit, logg)
	logg.Info("service init", slog.String("service", "postgres"))

	handlers := httphandler.NewHandler(srvce, logg)
	logg.Info("handler init", slog.String("handler", "postgres"))

	srv := http.Server{
		Addr:         cfg.Address + ":" + strconv.Itoa(cfg.HTTPServer.Port),
		Handler:      handlers.InitRoutes(logg),
		ReadTimeout:  cfg.HTTPServer.ReadTimeout,
		WriteTimeout: cfg.HTTPServer.WriteTimeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("failed to start server", err)
		}
	}()
	logg.Info("server started", slog.String("address", cfg.Address+":"+strconv.Itoa(cfg.HTTPServer.Port)))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logg.Error("failed to stop server", err)
	}

	logg.Info("server stopped by graceful shutdown")

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

//TODO: check unused handlers
//TODO: check unused services methods
//TODO: check TODO in lessons