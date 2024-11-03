package redis

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/not_amazing_amazon/internal/users"
)

// Those constants are the keys for fields of Session struct in Redis.
const (
	fieldID         = "ID"
	fieldToken      = "Token"
	fieldExpiration = "Expiration"
)

// This time format is used to store a timestamp as a string in Redis
const timeFormat = time.RFC3339

func addSessionData(session users.Session) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		err := p.HSet(
			ctx, makeIDKey(session.User),
			fieldID, session.User.String(),
			fieldToken, session.Token,
			fieldExpiration, session.ExpiresAt.Format(timeFormat),
		).Err()
		if err != nil {
			return err
		}

		if err := p.ExpireAt(ctx, makeIDKey(session.User), session.ExpiresAt).Err(); err != nil {
			return err
		}

		return nil
	}
}

func addTokenIndex(session users.Session) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		duration := time.Until(session.ExpiresAt)
		if err := p.Set(ctx, makeTokenKey(session.Token), makeIDKey(session.User), duration).Err(); err != nil {
			return err
		}

		return nil
	}
}

func removeSession(id uuid.UUID) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		err := p.Del(ctx, makeIDKey(id)).Err()
		if errors.Is(err, redis.Nil) {
			return users.ErrNotFound
		}
		return err
	}
}

func removeTokenIndex(token string) txAction {
	return func(ctx context.Context, p redis.Pipeliner) error {
		err := p.Del(ctx, makeTokenKey(token)).Err()
		if errors.Is(err, redis.Nil) {
			return users.ErrNotFound
		}
		return err
	}
}
