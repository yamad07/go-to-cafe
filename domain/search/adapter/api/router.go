package search

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	h := NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/", h.Search)
	return r
}

func NewAdminRouter(repo registry.Repository) http.Handler {
	h := NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/", h.Search)
	return r
}
