package adapter

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/domain/auth/adapter/api/login"
	"github.com/yamad07/go-modular-monolith/domain/auth/adapter/api/sign_up"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
)

func NewRouter(repo registry.AuthRepository) http.Handler {

	r := chi.NewRouter()
	r.Post("/login", login.NewHandler(repo).Login)
	r.Post("/sign_up", sign_up.NewHandler(repo).SignUp)
	return r
}
