package repository

import (
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/model"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

type Cafe interface {
	List(ids []int64) ([]model.Cafe, apperror.Error)
	Create(value.CreateCafeParam) apperror.Error
	Update(model.Cafe) apperror.Error
	GetByCode(code string) (model.Cafe, apperror.Error)
}
