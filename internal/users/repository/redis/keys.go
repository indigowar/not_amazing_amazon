package redis

import (
	"fmt"

	"github.com/google/uuid"
)

func makeIDKey(id uuid.UUID) string {
	return fmt.Sprintf("sessions:%s", id.String())
}

func makeTokenKey(token string) string {
	return fmt.Sprintf("sessions:token:%s", token)
}

