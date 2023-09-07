package integ

type RsHourly struct {
	Cod     string                `json:"cod"`
	Message int                   `json:"message"`
	Cnt     int                   `json:"cnt"`
	List    []RsHourlyWeatherData `json:"list"`
	City    RsHourlyCityInfo      `json:"city"`
}

type RsHourlyWeatherData struct {
	Dt         int64              `json:"dt"`
	Main       RsMain             `json:"main"`
	Weather    []RsWeather        `json:"weather"`
	Clouds     RsHourlyCloudsInfo `json:"clouds"`
	Wind       RsHourlyWindInfo   `json:"wind"`
	Visibility int                `json:"visibility"`
	Pop        float64            `json:"pop"`
	Sys        RsHourlySysInfo    `json:"sys"`
	DtTxt      string             `json:"dt_txt"`
}

type RsHourlyCloudsInfo struct {
	All int `json:"all"`
}

type RsHourlyWindInfo struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type RsHourlySysInfo struct {
	Pod string `json:"pod"`
}

type RsHourlyCityInfo struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Coord      RsCoord `json:"coord"`
	Country    string  `json:"country"`
	Population int     `json:"population"`
	Timezone   int     `json:"timezone"`
	Sunrise    int64   `json:"sunrise"`
	Sunset     int64   `json:"sunset"`
}
