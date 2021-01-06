package cafe

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	h := NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/", h.Fetch)
	r.Get("/{code}", h.Find)
	r.Post("/", h.Create)
	r.Put("/", h.Update)
	return r
}
