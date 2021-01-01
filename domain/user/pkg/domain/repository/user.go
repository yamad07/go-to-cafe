package repository

import (
	"github.com/yamad07/go-modular-monolith/domain/user/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

type User interface {
	Create(value.CreateUserParam) apperror.Error
}
