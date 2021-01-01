package cafe

import "github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/model"

type CafeOutput struct {
	Cafe model.Cafe
}

type CafesOutput struct {
	Cafes []model.Cafe
}
