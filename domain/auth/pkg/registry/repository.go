package registry

import (
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/infra/dao"
	"github.com/yamad07/go-modular-monolith/pkg/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() (Repository, func() error) {

	db, cleanup, err := mysql.Init()
	if err != nil {
		panic(err)
	}

	return Repository{
		db: db,
	}, cleanup
}

func (r Repository) NewAccount() dao.Account {
	return dao.NewAccount(r.db)
}
