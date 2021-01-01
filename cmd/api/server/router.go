package main

import (
	"net/http"

	"github.com/go-chi/chi"
	auth "github.com/yamad07/go-modular-monolith/domain/auth/adapter/api"
	cafe "github.com/yamad07/go-modular-monolith/domain/cafe/adapter/api"
	search "github.com/yamad07/go-modular-monolith/domain/search/adapter/api"
	"github.com/yamad07/go-modular-monolith/pkg/presenter"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()
	r = CommonMiddleware(r)

	r.Mount("/admin", newAdminRouter(repo))
	r.Mount("/app", newAppRouter(repo))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Success(w)
	})

	return r
}

func newAppRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()
	r = CommonMiddleware(r)

	r.Mount("/auth", auth.NewRouter(repo.NewAuth()))
	r.Mount("/cafes", cafe.NewRouter(repo.NewCafe()))
	r.Mount("/search", search.NewRouter(repo.NewSearch()))

	return r
}

func newAdminRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()

	r.Mount("/cafes", cafe.NewAdminRouter(repo.NewCafe()))

	return r
}
