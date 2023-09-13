package integ

type RsClimate struct {
	City    RsClimateCityInfo      `json:"city"`
	Code    string                 `json:"code"`
	Message float64                `json:"message"`
	Cnt     int                    `json:"cnt"`
	List    []RsClimateWeatherData `json:"list"`
}

type RsClimateCityInfo struct {
	Id         int                `json:"id"`
	Name       string             `json:"name"`
	Coord      RsClimateCoordInfo `json:"coord"`
	Country    string             `json:"country"`
	Population int                `json:"population"`
	Timezone   int                `json:"timezone"`
}

type RsClimateCoordInfo struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type RsClimateWeatherData struct {
	Dt        int64                    `json:"dt"`
	Sunrise   int64                    `json:"sunrise"`
	Sunset    int64                    `json:"sunset"`
	Temp      RsClimateTemperatureInfo `json:"temp"`
	FeelsLike RsClimateFeelsLikeInfo   `json:"feels_like"`
	Pressure  int                      `json:"pressure"`
	Humidity  int                      `json:"humidity"`
	Weather   []RsWeather              `json:"weather"`
	Speed     float64                  `json:"speed"`
	Deg       int                      `json:"deg"`
	Clouds    int                      `json:"clouds"`
	Rain      float64                  `json:"rain"`
}

type RsClimateTemperatureInfo struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type RsClimateFeelsLikeInfo struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}
