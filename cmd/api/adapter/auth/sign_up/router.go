package sign_up

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
)

func NewRouter(repo registry.AuthRepository) http.Handler {

	r := chi.NewRouter()
	r.Post("/", NewHandler(repo).SignUp)
	return r
}
