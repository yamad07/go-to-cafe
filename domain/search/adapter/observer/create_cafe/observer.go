package create_cafe

import (
	"log"

	"github.com/yamad07/go-modular-monolith/domain/search/pkg/registry"
	search "github.com/yamad07/go-modular-monolith/domain/search/usecase"
	"github.com/yamad07/go-modular-monolith/pkg/boundary/create_cafe"
)

type Observer struct {
	usecase search.Usecase
}

func NewObserver(repo registry.Repository) {
	observer := Observer{
		usecase: search.NewUsecase(repo),
	}
	create_cafe.NewNotifier().Register(observer)
}

func (o Observer) OnNotify(evt create_cafe.Event) {
	ipt := search.CreateInput{
		ID:        evt.ID(),
		Name:      evt.Name(),
		Latitude:  evt.Latitude(),
		Longitude: evt.Longitude(),
	}
	aerr := o.usecase.Create(ipt)
	//TODO rich logger or sentry
	log.Println(aerr)
}
