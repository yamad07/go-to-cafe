package cafe

type EventData struct {
	id        int64
	name      string
	latitude  float64
	longitude float64
}

func (e EventData) ID() int64 {
	return e.id
}

func (e EventData) Name() string {
	return e.name
}

func (e EventData) Latitude() float64 {
	return e.latitude
}

func (e EventData) Longitude() float64 {
	return e.longitude
}
