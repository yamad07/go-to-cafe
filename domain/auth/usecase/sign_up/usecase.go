package sign_up

import (
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/repository"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
	"github.com/yamad07/go-modular-monolith/pkg/boundary/sign_up"
)

type Usecase struct {
	accountRepository repository.Account
}

func NewUsecase(
	accountRepository repository.Account,
) Usecase {
	return Usecase{
		accountRepository,
	}
}

func (u Usecase) SignUp(input Input) (output Output, aerr apperror.Error) {
	p := value.NewCreateAccountParam(input.Email, input.Password)
	aerr = u.accountRepository.Create(p)
	if aerr != nil {
		return output, aerr
	}
	act, aerr := u.accountRepository.GetByUID(p.UID)
	if aerr != nil {
		return output, aerr
	}

	evt := EventData{
		accountID: act.ID,
		name:      input.Name,
	}
	sign_up.NewNotifier().Notify(evt)

	output = Output{
		Account: act,
	}

	return output, nil
}
