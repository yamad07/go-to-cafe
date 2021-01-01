package repository

import (
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/domain/model"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

type Search interface {
	Create(value.CreateCafeIndex) apperror.Error
	Search(value.RangeQuery) ([]model.CafeCandidate, apperror.Error)
}
