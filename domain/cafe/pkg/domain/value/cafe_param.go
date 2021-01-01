package value

import "github.com/google/uuid"

type CreateCafeParam struct {
	Name      string
	Code      string
	Latitude  float64
	Longitude float64
}

func NewCreateParam(
	name string,
	latitude float64,
	longitude float64,
) CreateCafeParam {
	code := uuid.New().String()
	return CreateCafeParam{
		Name:      name,
		Code:      code,
		Latitude:  latitude,
		Longitude: longitude,
	}
}
