package view

import "github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/model"

type Cafe struct {
	ID        int64   `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewCafe(m model.Cafe) Cafe {
	return Cafe{
		ID:        m.ID,
		Code:      m.Code,
		Name:      m.Name,
		Latitude:  m.Latitude,
		Longitude: m.Longitude,
	}
}

func NewCafes(cfs []model.Cafe) (v []Cafe) {
	for _, cf := range cfs {
		v = append(v, NewCafe(cf))
	}

	return v
}
