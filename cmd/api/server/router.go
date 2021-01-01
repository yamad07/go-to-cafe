package main

import (
	"net/http"

	"github.com/go-chi/chi"
	auth "github.com/yamad07/go-modular-monolith/domain/auth/adapter/api"
	"github.com/yamad07/go-modular-monolith/pkg/presenter"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()
	r = CommonMiddleware(r)

	r.Mount("/auth", auth.NewRouter(repo.NewAuth()))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Success(w)
	})

	return r
}
