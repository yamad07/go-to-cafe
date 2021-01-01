package value

import (
	"github.com/google/uuid"
)

type CreateAccountParam struct {
	Email    string
	Password Password
	UID      string
}

func NewCreateAccountParam(
	email string,
	password string,
) CreateAccountParam {
	uuid := uuid.New()
	return CreateAccountParam{
		Email:    email,
		Password: Password(password),
		UID:      uuid.String(),
	}
}
