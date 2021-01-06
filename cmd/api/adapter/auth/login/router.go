package login

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/domain/auth/adapter/api/login"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
)

func NewRouter(repo registry.AuthRepository) http.Handler {

	r := chi.NewRouter()
	r.Post("/", login.NewHandler(repo).Login)
	return r
}
