package search

type request struct {
	MaxLatitude  float64 `schema:"max_latitude"`
	MinLatitude  float64 `schema:"min_latitude"`
	MaxLongitude float64 `schema:"max_longitude"`
	MinLongitude float64 `schema:"min_longitude"`
}
