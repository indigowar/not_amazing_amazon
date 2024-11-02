package users

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type UserService struct {
	logger   *slog.Logger
	userRepo UserRepository
}

func (svc *UserService) SignIn(
	ctx context.Context,
	passport string,
	password string,
	displayedName string,
	phoneNumber string,
) (uuid.UUID, error) {
	panic("unimplemented")
}

func NewUserService(
	logger *slog.Logger,
	userRepo UserRepository,
) *UserService {
	return &UserService{
		logger:   logger,
		userRepo: userRepo,
	}
}
