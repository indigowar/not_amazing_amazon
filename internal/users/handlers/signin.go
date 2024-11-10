package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/indigowar/not_amazing_amazon/internal/common/web"
	"github.com/indigowar/not_amazing_amazon/internal/users"
	"github.com/indigowar/not_amazing_amazon/internal/users/handlers/templates"
)

func ShowSigninPage(handlerPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.Render(r.Context(), w, http.StatusOK, templates.SignIn(handlerPath))
	}
}

func HandleSignin(sm *scs.SessionManager, svc *users.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := web.GetUserID(sm, r); err == nil {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}

		displayedName := r.FormValue("display_name")
		phoneNumber := r.FormValue("phone_number")
		password := r.FormValue("password")

		id, err := svc.SignIn(r.Context(), phoneNumber, password, displayedName)
		if err != nil {
			// Handle the error
			w.Write([]byte(err.Error()))
			return
		}

		sm.Put(r.Context(), web.UserIDInSession, id.String())

		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}
