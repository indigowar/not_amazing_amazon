package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID

	Passport string
	Password string

	DisplayedName string
	PhoneNumber   string

	RegistrationDate time.Time
}
