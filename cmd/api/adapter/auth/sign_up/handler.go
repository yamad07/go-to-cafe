package sign_up

import (
	"encoding/json"
	"net/http"

	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/view"
	"github.com/yamad07/go-modular-monolith/domain/auth/usecase/sign_up"
	"github.com/yamad07/go-modular-monolith/pkg/presenter"
)

type handler struct {
	usecase sign_up.Usecase
}

func NewHandler(repo registry.AuthRepository) handler {
	usecase := sign_up.NewUsecase(repo)
	return handler{usecase}
}

func (h handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var body signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	ipt := sign_up.Input{
		Email:    body.Email,
		Password: body.Password,
	}
	opt, aerr := h.usecase.SignUp(ipt)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	v := view.NewAccount(opt.Account)
	presenter.Encode(w, v)
}
