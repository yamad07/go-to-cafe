package registration

import (
	"github.com/yamad07/go-modular-monolith/domain/user/pkg/domain/repository"
	"github.com/yamad07/go-modular-monolith/domain/user/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

type Usecase struct {
	userRepository repository.User
}

func NewUsecase(
	userRepository repository.User,
) Usecase {
	return Usecase{
		userRepository: userRepository,
	}
}

func (u Usecase) Registration(input Input) apperror.Error {
	p := value.CreateUserParam{
		AccountID: input.AccountID,
		Name:      input.Name,
	}
	return u.userRepository.Create(p)
}
