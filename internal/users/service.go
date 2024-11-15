package users

import (
	"context"
	"log/slog"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger      *slog.Logger
	userStorage UserStorage
	securityKey []byte
}

func (svc *Service) SignIn(
	ctx context.Context,
	phoneNumber string,
	password string,
	displayedName string,
) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.UUID{}, err
	}

	user := User{
		ID:               uuid.New(),
		PhoneNumber:      phoneNumber,
		Password:         string(hashedPassword),
		DisplayedName:    displayedName,
		RegistrationDate: time.Now(),
	}

	if err := svc.userStorage.Insert(ctx, user); err != nil {
		return uuid.UUID{}, err
	}

	return user.ID, nil
}

func (svc *Service) CheckCredentials(ctx context.Context, phone string, password string) (uuid.UUID, error) {
	user, err := svc.userStorage.GetByPhoneNumber(ctx, phone)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return uuid.UUID{}, err
	}

	return user.ID, nil
}

func NewUserService(
	logger *slog.Logger,
	userStorage UserStorage,
	securityKey []byte,
) *Service {
	return &Service{
		logger:      logger,
		userStorage: userStorage,
		securityKey: securityKey,
	}
}

func generateToken() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	l := len(alphabet)

	var sb strings.Builder
	for i := 0; i < 64; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}

	return sb.String()
}
