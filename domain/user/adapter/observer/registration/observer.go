package registration

import (
	"log"

	"github.com/yamad07/go-modular-monolith/domain/user/pkg/registry"
	"github.com/yamad07/go-modular-monolith/domain/user/usecase/registration"
	"github.com/yamad07/go-modular-monolith/pkg/boundary/sign_up"
)

type Observer struct {
	usecase registration.Usecase
}

func NewObserver(repo registry.Repository) {
	observer := Observer{
		usecase: registration.NewUsecase(
			repo.NewUser(),
		),
	}
	sign_up.NewNotifier().Register(observer)
}

func (o Observer) OnNotify(evt sign_up.Event) {
	ipt := registration.Input{
		Name:      evt.Name(),
		AccountID: evt.AccountID(),
	}
	aerr := o.usecase.Registration(ipt)
	//TODO rich logger or sentry
	log.Println(aerr)
}
