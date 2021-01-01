package adapter

import (
	"encoding/json"
	"net/http"

	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/view"
	"github.com/yamad07/go-modular-monolith/domain/auth/usecase/login"
	"github.com/yamad07/go-modular-monolith/pkg/presenter"
)

type handler struct {
	usecase login.Usecase
}

func NewHandler(repo registry.Repository) handler {
	usecase := login.NewUsecase(repo.NewAccount())
	return handler{usecase}
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	var body loginRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	ipt := login.Input{
		Email:    body.Email,
		Password: body.Password,
	}
	opt, aerr := h.usecase.Login(ipt)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	v := view.NewAccount(opt.Account)
	presenter.Encode(w, v)
}
