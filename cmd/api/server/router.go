package main

import (
	"net/http"

	"github.com/go-chi/chi"
	auth "github.com/yamad07/go-modular-monolith/domain/auth/adapter"
	"github.com/yamad07/go-modular-monolith/pkg/presenter"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Mount("auth", auth.NewRouter())
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Success(w)
	})

	return r
}
