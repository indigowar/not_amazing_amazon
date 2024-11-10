package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/indigowar/not_amazing_amazon/internal/common/web"
	"github.com/indigowar/not_amazing_amazon/internal/users"
)

type SetupConfig struct {
	GlobalMiddleware web.Middleware
	Service          *users.Service
	SessionManager   *scs.SessionManager
}

func Setup(m *http.ServeMux, cfg *SetupConfig) {
	m.HandleFunc("GET /signin", ShowSigninPage("/signin"))
	m.HandleFunc("POST /signin", HandleSignin(cfg.SessionManager, cfg.Service))

	m.HandleFunc("GET /login", ShowLoginPage())
	m.HandleFunc("POST /login", HandleLogin(cfg.SessionManager))

	m.Handle("POST /logout", HandleLogout(cfg.SessionManager, "/login"))
}
