package adapter

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
)

func NewRouter() http.Handler {
	repo, cleanup := registry.NewRepository()
	defer cleanup()

	h := NewHandler(repo)
	r := chi.NewRouter()
	r.Post("/", h.Login)
	return r
}
