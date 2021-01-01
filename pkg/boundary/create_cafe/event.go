package create_cafe

type Event interface {
	ID() int64
	Name() string
	Latitude() float64
	Longitude() float64
}
