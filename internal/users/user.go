package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID
	PhoneNumber      string
	Password         string
	DisplayedName    string
	RegistrationDate time.Time
}
