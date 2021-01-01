package entity

import "github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/model"

type Cafe struct {
	ID        int64   `gorm:"column:id"`
	Code      string  `gorm:"column:code"`
	Name      string  `gorm:"column:name"`
	Latitude  float64 `gorm:"column:latitude"`
	Longitude float64 `gorm:"column:longitude"`
}

func (Cafe) TableName() string {
	return "cafes"
}

func (c Cafe) ToModel() model.Cafe {
	return model.Cafe{
		ID:        c.ID,
		Code:      c.Code,
		Name:      c.Name,
		Latitude:  c.Latitude,
		Longitude: c.Longitude,
	}
}
