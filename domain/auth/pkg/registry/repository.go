package registry

import (
	"github.com/yamad07/go-modular-monolith/domain/auth/pkg/infra/dao"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {

	return AuthRepository{
		db: db,
	}
}

func (r AuthRepository) NewAccount() dao.Account {
	return dao.NewAccount(r.db)
}
