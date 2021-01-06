package search

import (
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/domain/repository"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/registry"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

type Usecase struct {
	searchRepository repository.Search
}

func NewUsecase(
	repo registry.Repository,
) Usecase {
	return Usecase{
		searchRepository: repo.NewSearch(),
	}
}

func (u Usecase) Create(input CreateInput) apperror.Error {
	idx := value.CreateCafeIndex{
		ID:        input.ID,
		Name:      input.Name,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
	}
	return u.searchRepository.Create(idx)
}

func (u Usecase) Search(input SearchInput) (output Output, aerr apperror.Error) {
	q := value.NewRangeQuery(
		input.MinLatitude,
		input.MaxLatitude,
		input.MinLongitude,
		input.MaxLongitude,
	)
	cands, aerr := u.searchRepository.Search(q)
	if aerr != nil {
		return output, aerr
	}
	output = Output{
		cands,
	}
	return output, nil
}
