package integ

type RsLocation struct {
	CityId    string `json:"city_id"`
	CountryId string `json:"country_id"`
	RegionId  string `json:"region_id"`
	Name      string `json:"name"`
}

type RsLocations struct {
	Location []RsLocation `json:"city"`
}
