package cafe

import (
	"errors"

	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/repository"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/registry"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
	"github.com/yamad07/go-modular-monolith/pkg/boundary/create_cafe"
)

type Usecase struct {
	cafeRepository repository.Cafe
}

func NewUsecase(
	repo registry.Repository,
) Usecase {
	return Usecase{
		cafeRepository: repo.NewCafe(),
	}
}

func (u Usecase) Get(input GetInput) (output CafeOutput, aerr apperror.Error) {
	cafe, aerr := u.cafeRepository.GetByCode(input.Code)
	if aerr != nil {
		return output, aerr
	}
	output = CafeOutput{
		Cafe: cafe,
	}
	return output, nil
}

func (u Usecase) Fetch(input FetchInput) (output CafesOutput, aerr apperror.Error) {
	cafes, aerr := u.cafeRepository.List(input.IDs)
	if aerr != nil {
		return output, aerr
	}
	output = CafesOutput{
		Cafes: cafes,
	}
	return output, nil
}

func (u Usecase) Create(input CreateInput) (output CafeOutput, aerr apperror.Error) {
	p := value.NewCreateParam(
		input.Name,
		input.Latitude,
		input.Longitude,
	)
	aerr = u.cafeRepository.Create(p)
	if aerr != nil {
		return output, aerr
	}

	cafe, aerr := u.cafeRepository.GetByCode(p.Code)
	if aerr != nil {
		return output, aerr
	}

	evt := EventData{
		id:        cafe.ID,
		name:      cafe.Name,
		longitude: cafe.Longitude,
		latitude:  cafe.Latitude,
	}
	create_cafe.NewNotifier().Notify(evt)

	output = CafeOutput{
		Cafe: cafe,
	}
	return output, nil
}

func (u Usecase) Update(input UpdateInput) (output CafeOutput, aerr apperror.Error) {
	cafe, aerr := u.cafeRepository.GetByCode(input.Code)
	if cafe.ID != input.ID {
		err := errors.New("cafe id and code is invalid.")
		return output, apperror.New(apperror.CodeInvalid, err)
	}
	if input.Name != "" {
		cafe.Name = input.Name
	}

	if input.Latitude != 0 {
		cafe.Latitude = input.Latitude
	}

	if input.Longitude != 0 {
		cafe.Longitude = input.Longitude
	}

	aerr = u.cafeRepository.Update(cafe)
	if aerr != nil {
		return output, aerr
	}

	cafe, aerr = u.cafeRepository.GetByCode(input.Code)
	if aerr != nil {
		return output, aerr
	}

	output = CafeOutput{
		Cafe: cafe,
	}
	return output, nil
}
