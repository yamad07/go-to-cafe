package view

import "github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/model"

type Account struct {
	ID     int64  `json:"id"`
	UID    string `json:"uid"`
	UserID int64  `json:"user_id"`
}

func NewAccount(m model.Account) Account {
	return Account{
		ID:     m.ID,
		UID:    m.UID,
		UserID: m.UserID,
	}
}
