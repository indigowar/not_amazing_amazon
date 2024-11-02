package users

import (
	"net/http"

	"github.com/indigowar/not_amazing_amazon/internal/users/views"
)

type ViewConfig struct {
	Service *UserService
}

func SetupHandlers(mux *http.ServeMux, cfg ViewConfig) {
	mux.HandleFunc("GET /auth/signin", func(w http.ResponseWriter, r *http.Request) {
		_ = views.SignInPage().Render(r.Context(), w)
	})
}
