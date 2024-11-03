package health

import (
	"net/http"

	"github.com/indigowar/not_amazing_amazon/internal/common/web"
)

func HealthEndpoint(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := svc.Health(r.Context())

		if health.Status != "ok" {
			web.JSON(w, http.StatusServiceUnavailable, health)
			return
		}

		web.JSON(w, http.StatusOK, health)
	}
}

func HealthDetailEndpoint(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		health := svc.HealthDetailed(r.Context())

		if health.Status != "ok" {
			web.JSON(w, http.StatusServiceUnavailable, health)
			return
		}

		web.JSON(w, http.StatusOK, health)
	}
}

func SetupHandlers(mux *http.ServeMux, svc *Service) {
	mux.Handle("GET /health", HealthEndpoint(svc))
	mux.Handle("GET /health/detailed", HealthDetailEndpoint(svc))
}
