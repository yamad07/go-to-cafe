package cafe

type GetInput struct {
	Code string
}

type CreateInput struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type UpdateInput struct {
	ID        int64
	Code      string
	Name      string
	Latitude  float64
	Longitude float64
}

type FetchInput struct {
	IDs []int64
}
