package handlers

import (
	"log/slog"
	"net/http"

	"github.com/indigowar/not_amazing_amazon/internal/common/web"
	"github.com/indigowar/not_amazing_amazon/internal/users/handlers/templates"
)

func ShowSigninPage(handlerPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.Render(r.Context(), w, http.StatusOK, templates.SignIn(handlerPath))
	}
}

func HandleSignin(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		passport := r.FormValue("passport")
		password := r.FormValue("passport")
		email := r.FormValue("email")
		phoneNumber := r.FormValue("phone_number")
		displayedName := r.FormValue("display_name")

		logger.Info(
			"received sign in handle",
			"passport", passport,
			"password", password,
			"email", email,
			"phoneNumber", phoneNumber,
			"displayedName", displayedName,
		)
	}
}
