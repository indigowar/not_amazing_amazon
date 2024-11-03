package redis

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/indigowar/not_amazing_amazon/internal/users"
)

type SessionStorage struct {
	client *redis.Client
}

var _ users.SessionStorage = &SessionStorage{}

// Delete implements users.SessionStorage.
func (s *SessionStorage) Delete(ctx context.Context, token string) error {
	session, err := s.GetBytToken(ctx, token)
	if err != nil {
		return err
	}

	err = makeTx(
		s.client,
		removeSession(session.User),
		removeTokenIndex(session.Token),
	)(ctx)

	if err != nil {
		return err
	}

	return nil
}

// GetByID implements users.SessionStorage.
func (s *SessionStorage) GetByID(ctx context.Context, id uuid.UUID) (users.Session, error) {
	data, err := s.client.HGetAll(ctx, makeIDKey(id)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return users.Session{}, users.ErrNotFound
		}

		return users.Session{}, err
	}

	return sessionFromData(data)
}

// GetBytToken implements users.SessionStorage.
func (s *SessionStorage) GetBytToken(ctx context.Context, token string) (users.Session, error) {
	sessionKey, err := s.client.Get(ctx, makeTokenKey(token)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return users.Session{}, users.ErrNotFound
		}
		return users.Session{}, err
	}

	data, err := s.client.HGetAll(ctx, sessionKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return users.Session{}, users.ErrNotFound
		}
		return users.Session{}, err
	}

	return sessionFromData(data)
}

// Insert implements users.SessionStorage.
func (s *SessionStorage) Insert(ctx context.Context, session users.Session) error {
	return makeTx(
		s.client,
		addSessionData(session),
		addTokenIndex(session),
	)(ctx)
}

func NewSessionStorage(client *redis.Client) *SessionStorage {
	return &SessionStorage{client}
}

func sessionFromData(data map[string]string) (users.Session, error) {
	var id uuid.UUID
	var token string
	var expDate time.Time

	var err error
	for key, value := range data {
		switch key {
		case fieldID:
			id, err = uuid.Parse(value)
			if err != nil {
				return users.Session{}, err
			}
		case fieldToken:
			token = value
		case fieldExpiration:
			expDate, err = time.Parse(timeFormat, value)
			if err != nil {
				return users.Session{}, err
			}
		default:
			return users.Session{}, errors.New("invalid key")
		}
	}

	return users.Session{
		User:      id,
		Token:     token,
		ExpiresAt: expDate,
	}, nil
}
