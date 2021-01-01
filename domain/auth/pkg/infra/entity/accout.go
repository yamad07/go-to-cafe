package entity

import "github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/model"

type Account struct {
	ID                int64  `gorm:"column:id"`
	Email             string `gorm:"column:email"`
	UID               string `gorm:"column:uid"`
	EncryptedPassword string `gorm:"column:encrypted_password"`
}

func (a Account) ToModel() model.Account {
	return model.Account{
		ID:                a.ID,
		Email:             a.Email,
		UID:               a.UID,
		EncryptedPassword: a.EncryptedPassword,
	}
}
