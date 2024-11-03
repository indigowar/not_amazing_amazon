package users

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger      *slog.Logger
	storage     UserStorage
	securityKey []byte
}

func (svc *Service) SignIn(
	ctx context.Context,
	passport string,
	password string,
	displayedName string,
	phoneNumber string,
) (uuid.UUID, error) {
	hashedPassport, err := bcrypt.GenerateFromPassword([]byte(passport), bcrypt.DefaultCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	hashedPhoneNumber, err := bcrypt.GenerateFromPassword([]byte(phoneNumber), bcrypt.DefaultCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := User{
		ID:               uuid.New(),
		Passport:         string(hashedPassport),
		Password:         string(hashedPassword),
		DisplayedName:    displayedName,
		PhoneNumber:      string(hashedPhoneNumber),
		RegistrationDate: time.Now(),
	}

	if err := svc.storage.Insert(ctx, user); err != nil {
		return uuid.UUID{}, err
	}

	return user.ID, nil
}

func NewUserService(
	logger *slog.Logger,
	storage UserStorage,
	securityKey []byte,
) *Service {
	return &Service{
		logger:      logger,
		storage:     storage,
		securityKey: securityKey,
	}
}
