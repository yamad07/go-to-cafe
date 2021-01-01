package registry

import elasticsearch "github.com/elastic/go-elasticsearch/v7"

type Repository struct {
	es *elasticsearch.Client
}

func NewRepository(es *elasticsearch.Client) Repository {

	return Repository{
		es: es,
	}
}
