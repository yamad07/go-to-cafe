package cafe

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/registry"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/view"
	cafe "github.com/yamad07/go-modular-monolith/domain/cafe/usecase"
	"github.com/yamad07/go-modular-monolith/pkg/presenter"
)

type handler struct {
	usecase cafe.Usecase
}

func NewHandler(repo registry.Repository) handler {
	usecase := cafe.NewUsecase(repo.NewCafe())
	return handler{usecase}
}

func (h handler) Fetch(w http.ResponseWriter, r *http.Request) {
	var p fetchRequest
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&p, r.URL.Query()); err != nil {
		presenter.BadRequestError(w, fmt.Errorf("invalid query params: %v", err))
		return
	}

	ipt := cafe.FetchInput{
		IDs: p.IDs,
	}
	opt, aerr := h.usecase.Fetch(ipt)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	v := view.NewCafes(opt.Cafes)
	presenter.Encode(w, v)
}

func (h handler) Find(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	ipt := cafe.GetInput{
		Code: code,
	}
	opt, aerr := h.usecase.Get(ipt)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	v := view.NewCafe(opt.Cafe)
	presenter.Encode(w, v)
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	var body createRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	ipt := cafe.CreateInput{
		Name:      body.Name,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
	}
	opt, aerr := h.usecase.Create(ipt)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	v := view.NewCafe(opt.Cafe)
	presenter.Encode(w, v)
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	var body updateRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	code := chi.URLParam(r, "code")

	ipt := cafe.UpdateInput{
		Name:      body.Name,
		Code:      code,
		Latitude:  body.Latitude,
		Longitude: body.Longitude,
	}
	opt, aerr := h.usecase.Update(ipt)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	v := view.NewCafe(opt.Cafe)
	presenter.Encode(w, v)
}
