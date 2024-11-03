package users

import (
	"net/http"

	"github.com/indigowar/not_amazing_amazon/internal/users/handlers"
)

type SetupConfig struct {
	Service *Service
}

func Setup(m *http.ServeMux, cfg SetupConfig) {
	m.HandleFunc("GET /signin", handlers.ShowSigninPage("/signin"))
	m.HandleFunc("POST /signin", handlers.HandleSignin(cfg.Service.logger))
}
