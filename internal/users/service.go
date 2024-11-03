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
	logger         *slog.Logger
	userStorage    UserStorage
	sessionStorage SessionStorage
	securityKey    []byte
}

func (svc *Service) SignIn(
	ctx context.Context,
	phoneNumber string,
	password string,
	displayedName string,
) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := User{
		ID:               uuid.New(),
		PhoneNumber:      phoneNumber,
		Password:         string(hashedPassword),
		DisplayedName:    displayedName,
		RegistrationDate: time.Now(),
	}

	if err := svc.userStorage.Insert(ctx, user); err != nil {
		return "", err
	}

	session := Session{
		User:      user.ID,
		Token:     generateToken(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := svc.sessionStorage.Insert(ctx, session); err != nil {
		return "", err
	}

	return session.Token, nil
}

func (svc *Service) Login(ctx context.Context, phone string, password string) (string, error) {
	panic("unimplemented")
}

func NewUserService(
	logger *slog.Logger,
	userStorage UserStorage,
	sessionStorage SessionStorage,
	securityKey []byte,
) *Service {
	return &Service{
		logger:         logger,
		userStorage:    userStorage,
		sessionStorage: sessionStorage,
		securityKey:    securityKey,
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
