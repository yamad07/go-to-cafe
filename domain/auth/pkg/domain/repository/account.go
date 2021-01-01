package repository

import (
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/model"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

type Account interface {
	GetByEmail(string) (model.Account, apperror.Error)
}
