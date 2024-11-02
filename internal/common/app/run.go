package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/not_amazing_amazon/internal/common/config"
	"github.com/indigowar/not_amazing_amazon/internal/health"
)

func Run(cfg *config.Config) {
	logger := createLogger(cfg)

	postgres, err := pgx.Connect(context.Background(), fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.Host, cfg.Postgres.Port,
		cfg.Postgres.Database,
	))
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL", "err", err)
		os.Exit(1)
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Protocol: 0,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})

	minio, err := minio.New(fmt.Sprintf("%s:%d", cfg.Minio.Host, cfg.Minio.Port), &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.RootUser, cfg.Minio.RootPassword, ""),
		Secure: false,
	})
	if err != nil {
		logger.Error("Failed to connect to MinIO", "err", err)
		os.Exit(1)
	}

	healthSvc := health.NewService(postgres, redis, minio)

	mux := http.NewServeMux()

	health.SetupHandlers(mux, healthSvc)

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
