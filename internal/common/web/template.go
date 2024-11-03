package web

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(ctx context.Context, w http.ResponseWriter, code int, component templ.Component) {
	w.WriteHeader(code)
	if err := component.Render(ctx, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
