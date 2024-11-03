package users

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	User      uuid.UUID
	Token     string
	ExpiresAt time.Time
}
