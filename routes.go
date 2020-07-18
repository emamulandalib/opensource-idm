package main

import (
	"net/http"

	v1 "github.com/emamulandalib/opensource-idm/v1"
	"github.com/go-chi/chi"
)

func MainRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/v1", v1.V1Router())
	return r
}
