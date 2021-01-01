package registry

import (
	"github.com/yamad07/go-modular-monolith/domain/user/pkg/infra/dao"
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

func (r Repository) NewUser() dao.User {
	return dao.NewUser(r.db)
}
