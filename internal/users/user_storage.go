package users

import (
	"context"

	"github.com/google/uuid"
)

type UserStorage interface {
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	GetByPassport(ctx context.Context, passport string) (User, error)
	GetByPhoneNumber(ctx context.Context, phone string) (User, error)
	Insert(ctx context.Context, user User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
