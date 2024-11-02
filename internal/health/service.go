package health

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	dbCon       *pgx.Conn
	redisClient *redis.Client
	minioCon    *minio.Client
}

func (svc *Service) Health(ctx context.Context) Health {
	dependencies := []Dependency{
		svc.checkDependency("database", func() error { return svc.dbCon.Ping(ctx) }),
		svc.checkDependency("session-storage", func() error { return svc.redisClient.Ping(ctx).Err() }),
		svc.checkDependency("image-storage", func() error {
			_, err := svc.minioCon.ListBuckets(ctx)
			return err
		}),
	}

	var err error = nil
	for _, dep := range dependencies {
		if dep.Error != nil {
			err = errors.Join(err, fmt.Errorf("%s is down", dep.Name))
		}
	}

	if err != nil {
		return Health{
			Status: "downgraded",
			Error:  err,
		}
	}

	return Health{Status: "ok"}
}

func (svc *Service) HealthDetailed(ctx context.Context) Health {
	dependencies := []Dependency{
		svc.checkDependency("database", func() error { return svc.dbCon.Ping(ctx) }),
		svc.checkDependency("session-storage", func() error { return svc.redisClient.Ping(ctx).Err() }),
		svc.checkDependency("image-storage", func() error {
			_, err := svc.minioCon.ListBuckets(ctx)
			return err
		}),
	}

	var err error = nil
	for _, dep := range dependencies {
		if dep.Error != nil {
			err = errors.Join(err, fmt.Errorf("%s is down", dep.Name))
		}
	}

	if err != nil {
		return Health{
			Status:       "downgraded",
			Error:        err,
			Dependencies: dependencies,
		}
	}

	return Health{Status: "ok", Dependencies: dependencies}
}

func (svc *Service) checkDependency(name string, action func() error) Dependency {
	start := time.Now()
	if err := action(); err != nil {
		return Dependency{
			Name:   name,
			Status: "down",
			Error:  err,
		}
	}

	return Dependency{
		Name:           name,
		Status:         "ok",
		ResponseTimeMs: time.Since(start).Milliseconds(),
	}
}

func NewService(db *pgx.Conn, redis *redis.Client, minio *minio.Client) *Service {
	return &Service{
		dbCon:       db,
		redisClient: redis,
		minioCon:    minio,
	}
}
