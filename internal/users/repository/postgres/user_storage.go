package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/indigowar/not_amazing_amazon/internal/users"
	"github.com/indigowar/not_amazing_amazon/internal/users/repository/postgres/gen"
)

//go:generate sqlc generate .

type UserStorage struct {
	conn *pgx.Conn
}

var _ users.UserStorage = &UserStorage{}

// Delete implements users.Repository.
func (s *UserStorage) Delete(ctx context.Context, id uuid.UUID) error {
	return gen.New(s.conn).DeleteUser(ctx, id)
}

// GetByID implements users.Repository.
func (s *UserStorage) GetByID(ctx context.Context, id uuid.UUID) (users.User, error) {
	user, err := gen.New(s.conn).SelectUserByID(ctx, id)
	if err != nil {
		return users.User{}, err
	}

	return s.userToModel(user), nil
}

// GetByPhoneNumber implements users.Repository.
func (s *UserStorage) GetByPhoneNumber(ctx context.Context, phone string) (users.User, error) {
	user, err := gen.New(s.conn).SelectUserByPhoneNumber(ctx, phone)
	if err != nil {
		return users.User{}, err
	}

	return s.userToModel(user), nil
}

// Insert implements users.Repository.
func (s *UserStorage) Insert(ctx context.Context, user users.User) error {
	return gen.New(s.conn).InsertUser(ctx, gen.InsertUserParams{
		ID:            user.ID,
		PhoneNumber:   user.PhoneNumber,
		Password:      []byte(user.Password),
		DisplayedName: user.DisplayedName,
	})
}

func (s *UserStorage) userToModel(user gen.User) users.User {
	return users.User{
		ID:               user.ID,
		PhoneNumber:      user.PhoneNumber,
		Password:         string(user.Password),
		DisplayedName:    user.DisplayedName,
		RegistrationDate: user.RegistrationDate.Time,
	}
}

func NewUserStorage(conn *pgx.Conn) *UserStorage {
	return &UserStorage{conn: conn}
}
