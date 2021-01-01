package observer

import (
	"github.com/yamad07/go-modular-monolith/domain/search/adapter/observer/create_cafe"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/registry"
)

func NewObserver(repo registry.Repository) {
	create_cafe.NewObserver(repo)
}
