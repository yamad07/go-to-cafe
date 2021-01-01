package registry

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/infra/dao"
)

type Repository struct {
	es *elasticsearch.Client
}

func NewRepository(es *elasticsearch.Client) Repository {

	return Repository{
		es: es,
	}
}

func (repo Repository) NewSearch() dao.Search {
	return dao.NewSearch(repo.es)
}
