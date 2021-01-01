package dao

import (
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/model"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/infra/entity"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) Account {
	return Account{db}
}

func (d Account) GetByEmail(email string) (act model.Account, aerr apperror.Error) {
	var a entity.Account
	if err := d.db.Where("email = ?", email).First(&a).Error; err != nil {
		return act, apperror.NewGormFind(err, "")
	}

	return a.ToModel(), nil
}

func (d Account) Create(p value.CreateAccountParam) (aerr apperror.Error) {
	enc, err := p.Password.Encrypt()
	if err != nil {
		return apperror.New(apperror.CodeError, err)
	}
	a := entity.Account{
		Email:             p.Email,
		UID:               p.UID,
		EncryptedPassword: enc,
	}

	if err := d.db.Create(&a).Error; err != nil {
		return apperror.NewGormCreate(err, "Account")
	}

	return nil
}

func (d Account) GetByUID(uid string) (act model.Account, aerr apperror.Error) {
	var a entity.Account
	if err := d.db.Where("uid = ?", uid).First(&a).Error; err != nil {
		return act, apperror.NewGormFind(err, "")
	}

	return a.ToModel(), nil
}
