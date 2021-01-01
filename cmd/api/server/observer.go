package main

import (
	search "github.com/yamad07/go-modular-monolith/domain/search/adapter/observer"
	user "github.com/yamad07/go-modular-monolith/domain/user/adapter/observer"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func NewObserver(repo registry.Repository) {
	user.NewObserver(repo.NewUser())
	search.NewObserver(repo.NewSearch())
}
