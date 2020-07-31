package v1

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Router ...
func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/", Index)
	r.Post("/download-file", DownloadFile)
	return r
}
