package dao

import (
	"testing"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
	"github.com/yamad07/go-modular-monolith/domain/search/pkg/domain/value"
	"github.com/yamad07/go-modular-monolith/pkg/config"
	elastic "github.com/yamad07/go-modular-monolith/pkg/elasticsearch"
)

var ES *elasticsearch.Client

func TestCafe(t *testing.T) {
	config.Init()
	ES, err := elastic.Init()
	if err != nil {
		panic(err)
	}

	cafe := NewCafe(ES)

	i := value.CreateCafeIndex{
		ID:        10,
		Latitude:  30.1,
		Longitude: 134.1,
	}
	aerr := cafe.Create(i)
	assert.NoError(t, aerr)

	q := value.RangeQuery{
		MinLatitude:  30.0,
		MaxLatitude:  30.2,
		MinLongitude: 134.0,
		MaxLongitude: 134.2,
	}

	candidates, aerr := cafe.SearchFromRange(q)
	assert.NoError(t, aerr)
	assert.Exactly(t, 1, len(candidates))

	q = value.RangeQuery{
		MinLatitude:  20.0,
		MaxLatitude:  20.2,
		MinLongitude: 134.0,
		MaxLongitude: 134.2,
	}

	candidates, aerr = cafe.SearchFromRange(q)
	assert.NoError(t, aerr)
	assert.Exactly(t, 0, len(candidates))
}
