package observer

import (
	"github.com/yamad07/go-modular-monolith/domain/user/adapter/observer/registration"
	"github.com/yamad07/go-modular-monolith/domain/user/pkg/registry"
)

func NewObserver(repo registry.Repository) {
	registration.NewObserver(repo)
}
