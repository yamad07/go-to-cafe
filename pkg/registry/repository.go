package registry

import (
	auth "github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
	user "github.com/yamad07/go-modular-monolith/domain/user/pkg/registry"
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

func (r Repository) NewAuth() auth.AuthRepository {
	return auth.NewAuthRepository(r.db)
}

func (r Repository) NewUser() user.Repository {
	return user.NewRepository(r.db)
}
