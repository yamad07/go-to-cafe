package cafe

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	h := NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/{code}", h.Find)
	r.Get("/", h.Fetch)
	return r
}

func NewAdminRouter(repo registry.Repository) http.Handler {
	h := NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/", h.Fetch)
	r.Get("/{code}", h.Find)
	r.Post("/", h.Create)
	r.Put("/", h.Update)
	return r
}
