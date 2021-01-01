package main

import (
	user "github.com/yamad07/go-modular-monolith/domain/user/adapter/observer"
	"github.com/yamad07/go-modular-monolith/pkg/registry"
)

func NewObserver(repo registry.Repository) {
	user.NewObserver(repo.NewUser())
}
