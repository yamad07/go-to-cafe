package registry

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	auth "github.com/yamad07/go-modular-monolith/domain/auth/pkg/registry"
	cafe "github.com/yamad07/go-modular-monolith/domain/cafe/pkg/registry"
	search "github.com/yamad07/go-modular-monolith/domain/search/pkg/registry"
	user "github.com/yamad07/go-modular-monolith/domain/user/pkg/registry"
	elastic "github.com/yamad07/go-modular-monolith/pkg/elasticsearch"
	"github.com/yamad07/go-modular-monolith/pkg/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	es *elasticsearch.Client
}

func NewRepository() (Repository, func() error) {
	db, cleanup, err := mysql.Init()
	if err != nil {
		panic(err)
	}

	es, err := elastic.Init()
	if err != nil {
		panic(err)
	}

	return Repository{
		db: db,
		es: es,
	}, cleanup
}

func (r Repository) NewAuth() auth.AuthRepository {
	return auth.NewAuthRepository(r.db)
}

func (r Repository) NewUser() user.Repository {
	return user.NewRepository(r.db)
}

func (r Repository) NewCafe() cafe.Repository {
	return cafe.NewRepository(r.db)
}

func (r Repository) NewSearch() search.Repository {
	return search.NewRepository(r.es)
}
