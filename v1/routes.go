package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

func V1Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/", Index)
	return r
}
