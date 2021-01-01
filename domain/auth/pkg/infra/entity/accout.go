package entity

import "github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/model"

type Account struct {
	ID                int64  `gorm:"column:id"`
	UserID            int64  `gorm:"column:user_id"`
	UID               string `gorm:"column:uid"`
	EncryptedPassword string `gorm:"column:encrypted_password"`
}

func (a Account) ToModel() model.Account {
	return model.Account{
		ID:                a.ID,
		UserID:            a.UserID,
		UID:               a.UID,
		EncryptedPassword: a.EncryptedPassword,
	}
}
