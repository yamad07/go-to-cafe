package auth

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yamad07/go-modular-monolith/cmd/api/adapter/auth/login"
	"github.com/yamad07/go-modular-monolith/cmd/api/adapter/auth/sign_up"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {

	r := chi.NewRouter()
	r.Mount("/login", login.NewRouter(repo.NewAuth()))
	r.Mount("/sign_up", sign_up.NewRouter(repo.NewAuth()))
	return r
}
