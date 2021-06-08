package elasticsearch

type ElasticsearchEntity struct {
	Name     string
	Location []Coordinate
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}
