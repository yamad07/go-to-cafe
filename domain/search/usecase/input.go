package search

type SearchInput struct {
	MaxLatitude  float64
	MinLatitude  float64
	MaxLongitude float64
	MinLongitude float64
}

type CreateInput struct {
	ID        int64
	Name      string
	Latitude  float64
	Longitude float64
}
