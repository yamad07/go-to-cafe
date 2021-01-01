package value

type RangeQuery struct {
	MinLatitude  float64
	MaxLatitude  float64
	MinLongitude float64
	MaxLongitude float64
}

func NewRangeQuery(
	minLatitude float64,
	maxLatitude float64,
	minLongitude float64,
	maxLongitude float64,
) RangeQuery {
	if minLatitude == 0 && maxLatitude == 0 {
		minLatitude = -90
		maxLatitude = 90
	}

	if minLongitude == 0 && maxLongitude == 0 {
		minLatitude = -180
		maxLatitude = 180
	}

	return RangeQuery{
		MaxLatitude:  maxLatitude,
		MinLatitude:  minLatitude,
		MaxLongitude: maxLongitude,
		MinLongitude: minLongitude,
	}
}
