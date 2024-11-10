package handlers

import (
	"net/http"

	"github.com/alexedwards/scs/v2"

	"github.com/indigowar/not_amazing_amazon/internal/common/web"
)

func HandleLogout(sm *scs.SessionManager, loginPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := web.GetUserID(sm, r)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sm.Remove(r.Context(), web.UserIDInSession)

		http.Redirect(w, r, loginPath, http.StatusPermanentRedirect)
		return
	}
}
