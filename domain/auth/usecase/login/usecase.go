package login

import (
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/repository"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

type Usecase struct {
	accountRepository repository.Account
}

func NewUsecase(
	accountRepository repository.Account,
) Usecase {
	return Usecase{
		accountRepository: accountRepository,
	}
}

func (u Usecase) Login(input Input) (output Output, aerr apperror.Error) {
	act, aerr := u.accountRepository.GetByEmail(input.Email)
	if aerr != nil {
		return output, aerr
	}
	return Output{
		Account: act,
	}, nil
}
