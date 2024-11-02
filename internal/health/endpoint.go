package health

import (
	"net/http"

	"github.com/indigowar/not_amazing_amazon/internal/common/rest"
)

func HealthEndpoint(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := svc.Health(r.Context())

		if health.Status != "ok" {
			rest.JSON(w, http.StatusServiceUnavailable, health)
			return
		}

		rest.JSON(w, http.StatusOK, health)
	}
}

func HealthDetailEndpoint(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := svc.HealthDetailed(r.Context())

		if health.Status != "ok" {
			rest.JSON(w, http.StatusServiceUnavailable, health)
			return
		}

		rest.JSON(w, http.StatusOK, health)
	}
}
