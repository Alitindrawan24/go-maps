package reversemap

type ReverseMaps struct {
	PlaceID     int      `json:"place_id"`
	License     string   `json:"license"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Address     Address  `json:"address"`
	BoundingBox []string `json:"boundingbox"`
}

type Address struct {
	Road         string `json:"road"`
	Village      string `json:"village"`
	Municipality string `json:"municipality"`
	State        string `json:"state"`
	PostCode     string `json:"post_code"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
}
