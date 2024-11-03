package users

import (
	"context"

	"github.com/google/uuid"
)

type SessionStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (Session, error)
	GetBytToken(ctx context.Context, token string) (Session, error)
	Insert(ctx context.Context, session Session) error
	Delete(ctx context.Context, token string) error
}
