package view

import "github.com/yamad07/go-modular-monolith/domain/search/pkg/domain/model"

type CafeCandidate struct {
	ID int64 `json:"id"`
}

func NewCafeCandidate(m model.CafeCandidate) CafeCandidate {
	return CafeCandidate{
		ID: m.ID,
	}
}

func NewCafeCandidates(cfs []model.CafeCandidate) (v []CafeCandidate) {
	for _, cf := range cfs {
		v = append(v, NewCafeCandidate(cf))
	}
	return v
}
