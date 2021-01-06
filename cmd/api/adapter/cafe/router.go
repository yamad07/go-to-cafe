package cafe

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/cmd/api/adapter/cafe/search"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()
	r.Mount("/search", search.NewRouter(repo.NewSearch()))
	return r
}
