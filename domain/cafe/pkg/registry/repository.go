package registry

import (
	"github.com/yamad07/go-modular-monolith/domain/cafe/pkg/infra/dao"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {

	return Repository{
		db: db,
	}
}

func (r Repository) NewCafe() dao.Cafe {
	return dao.NewCafe(r.db)
}
