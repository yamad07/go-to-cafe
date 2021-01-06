package search

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/registry"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/view"
	search "github.com/yamad07/go-modular-monolith/domain/search/usecase"
	"github.com/yamad07/go-modular-monolith/pkg/presenter"
)

type handler struct {
	usecase search.Usecase
}

func NewHandler(repo registry.Repository) handler {
	usecase := search.NewUsecase(repo)
	return handler{usecase}
}

func (h handler) Search(w http.ResponseWriter, r *http.Request) {
	var p request
	decoder := schema.NewDecoder()
	if err := decoder.Decode(&p, r.URL.Query()); err != nil {
		presenter.BadRequestError(w, fmt.Errorf("invalid query params: %v", err))
		return
	}

	ipt := search.SearchInput{
		MaxLatitude:  p.MaxLatitude,
		MinLatitude:  p.MinLatitude,
		MaxLongitude: p.MaxLongitude,
		MinLongitude: p.MinLongitude,
	}
	opt, aerr := h.usecase.Search(ipt)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	v := view.NewCafeCandidates(opt.CafeCandidates)
	presenter.Encode(w, v)
}
