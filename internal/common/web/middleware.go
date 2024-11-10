package web

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type Middleware = func(next http.Handler) http.Handler

func LoggingMiddleware(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Incoming request",
				"s method", r.Method,
				"s url", r.URL.String(),
				"s user-agent", r.UserAgent(),
			)

			lrw := &loggingResponseWriter{ResponseWriter: w}

			next.ServeHTTP(lrw, r)

			logger.Info("Request handled",
				"s status", lrw.statusCode,
				"s method", r.Method,
				"s url", r.URL.String(),
			)
		})

	}
}

// Custom ResponseWriter to capture data for logging
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func UserLoggedInMiddleware(sm *scs.SessionManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := GetUserID(sm, r)
			if err != nil {
				if errors.Is(err, ErrSessionObjectIsEmpty) {
					// TODO: Send here 403
				}

				// TODO: Send here 400
			}

			next.ServeHTTP(w, r)
		})

	}
}
