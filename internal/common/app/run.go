package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/indigowar/not_amazing_amazon/internal/common/config"
	"github.com/indigowar/not_amazing_amazon/internal/health"
	"github.com/indigowar/not_amazing_amazon/internal/users"
	userspostgres "github.com/indigowar/not_amazing_amazon/internal/users/repository/postgres"
)

func Run(cfg *config.Config) {
	logger := createLogger(cfg)
	postgres := connectToPostgres(cfg, logger)
	redis := connectToRedis(cfg)
	minio := connectToMinIO(cfg, logger)

	healthSvc := health.NewService(postgres, redis, minio)
	usersSvc := users.NewUserService(
		logger,
		userspostgres.NewUserStorage(postgres),
		[]byte(cfg.SecretKey),
	)

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	mux := http.NewServeMux()

	setupHandlers(
		mux,
		sessionManager,

		healthSvc,
		usersSvc,
	)

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	go func() {
		logger.Info("Starting the server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server.ListenAndServe failed", "err", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logger.Info("Graceful shutdown is initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Failed to gracefully shutdown", "err", err)
		os.Exit(1)
	}

	logger.Info("Graceful shutdown is completed")
}
