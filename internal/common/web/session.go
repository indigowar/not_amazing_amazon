package web

import (
	"errors"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"
)

var (
	ErrSessionObjectIsEmpty = errors.New("object with given session name is not found")
)

const UserIDInSession = "user-id"

func GetUserID(sm *scs.SessionManager, r *http.Request) (uuid.UUID, error) {
	value := sm.GetString(r.Context(), UserIDInSession)
	if value == "" {
		return uuid.UUID{}, ErrSessionObjectIsEmpty
	}

	id, err := uuid.Parse(value)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
