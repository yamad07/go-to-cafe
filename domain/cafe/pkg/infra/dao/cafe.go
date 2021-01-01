package dao

import (
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/model"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/infra/entity"
	"github.com/yamad07/go-modular-monolith/pkg/apperror"
	"gorm.io/gorm"
)

type Cafe struct {
	db *gorm.DB
}

func NewCafe(db *gorm.DB) Cafe {
	return Cafe{db}
}

func (c Cafe) List(ids []int64) (cafes []model.Cafe, aerr apperror.Error) {
	var cfs []entity.Cafe
	if err := c.db.Where("id in (?)", ids).Find(&cfs).Error; err != nil {
		return cafes, apperror.NewGormFind(err, "Cafe")
	}
	for _, c := range cfs {
		cafes = append(cafes, c.ToModel())
	}

	return cafes, nil
}

func (c Cafe) Create(p value.CreateCafeParam) (aerr apperror.Error) {
	cf := entity.Cafe{
		Code:      p.Code,
		Name:      p.Name,
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
	}
	if err := c.db.Create(&cf).Error; err != nil {
		return apperror.NewGormCreate(err, "Cafe")
	}

	return nil
}

func (c Cafe) GetByCode(code string) (cafe model.Cafe, aerr apperror.Error) {
	var cf entity.Cafe
	if err := c.db.Where("code = ?", code).Find(&cf).Error; err != nil {
		return cafe, apperror.NewGormFind(err, "Cafe")
	}

	return cf.ToModel(), nil
}

func (c Cafe) Update(cafe model.Cafe) (aerr apperror.Error) {
	cf := entity.Cafe{
		ID:        cafe.ID,
		Code:      cafe.Code,
		Name:      cafe.Name,
		Latitude:  cafe.Latitude,
		Longitude: cafe.Longitude,
	}
	if err := c.db.Save(&cf).Error; err != nil {
		return apperror.NewGormSave(err, "Cafe")
	}

	return nil
}
