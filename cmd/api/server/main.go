package main

import (
	"net/http"

	"github.com/yamad07/go-modular-monolith/pkg/config"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func main() {
	config.Init()
	repo, cleanup := registry.NewRepository()
	defer cleanup()

	NewObserver(repo)
	r := NewRouter(repo)
	http.ListenAndServe(":9000", r)
}
