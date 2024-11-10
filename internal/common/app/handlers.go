package app

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/indigowar/not_amazing_amazon/internal/health"
	"github.com/indigowar/not_amazing_amazon/internal/users"
	usershandlers "github.com/indigowar/not_amazing_amazon/internal/users/handlers"
)

func setupHandlers(
	mux *http.ServeMux,

	// INFRA

	sessionManager *scs.SessionManager,

	// SERVICES

	healthService *health.Service,
	userService *users.Service,

	// TODO: add here anything that should be passed to handlers
) {

	health.SetupHandlers(mux, healthService)

	usershandlers.Setup(
		mux,
		&usershandlers.SetupConfig{Service: userService, SessionManager: sessionManager},
	)
}
