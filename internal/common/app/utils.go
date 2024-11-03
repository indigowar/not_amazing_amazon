package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/not_amazing_amazon/internal/common/config"
)

func createLogger(cfg *config.Config) *slog.Logger {
	var logWriter io.Writer = os.Stdout
	if cfg.LogFile != "" {
		file, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		logWriter = io.MultiWriter(os.Stdout, file)
	}

	return slog.New(slog.NewJSONHandler(logWriter, nil))
}

func connectToPostgres(cfg *config.Config, logger *slog.Logger) *pgx.Conn {
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
	return postgres
}

func connectToRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Protocol: 0,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Db,
	})
}

func connectToMinIO(cfg *config.Config, logger *slog.Logger) *minio.Client {
	minio, err := minio.New(fmt.Sprintf("%s:%d", cfg.Minio.Host, cfg.Minio.Port), &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.RootUser, cfg.Minio.RootPassword, ""),
		Secure: false,
	})
	if err != nil {
		logger.Error("Failed to connect to MinIO", "err", err)
		os.Exit(1)
	}
	return minio
}
