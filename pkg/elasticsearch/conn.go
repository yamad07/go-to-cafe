package elastic

import (
	"fmt"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/yamad07/go-modular-monolith/pkg/config"
)

func Init() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			fmt.Sprintf("%s:%s", config.ElasticSearch.Host, config.ElasticSearch.Port),
		},
	}

	return elasticsearch.NewClient(cfg)
}
