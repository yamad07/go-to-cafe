package dao

import (
	"github.com/yamad07/go-modular-monolith/domain/user/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/domain/user/pkg/infra/entity"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return User{db}
}

func (d User) Create(p value.CreateUserParam) apperror.Error {
	u := entity.User{
		Name:      p.Name,
		AccountID: p.AccountID,
	}
	if err := d.db.Create(&u).Error; err != nil {
		return apperror.NewGormCreate(err, "User")
	}

	return nil
}
