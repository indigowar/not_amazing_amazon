package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func ShowLoginPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func HandleLogin(
	sm *scs.SessionManager,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
