package integ

type RsCurrent struct {
	Coord      RsCoord             `json:"coord"`
	Weather    []RsWeather         `json:"weather"`
	Base       string              `json:"base"`
	Main       RsMain              `json:"main"`
	Visibility int                 `json:"visibility"`
	Wind       RsCurrentWind       `json:"wind"`
	Clouds     RsCurrentClouds     `json:"clouds"`
	Rain       RsCurrentRainOrSnow `json:"rain"`
	Snow       RsCurrentRainOrSnow `json:"snow"`
	Dt         int64               `json:"dt"`
	Sys        RsCurrentSys        `json:"sys"`
	Timezone   int                 `json:"timezone"`
	Id         int                 `json:"id"`
	Name       string              `json:"name"`
	Cod        int                 `json:"cod"`
}

type RsCurrentWind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type RsCurrentClouds struct {
	All int `json:"all"`
}

type RsCurrentRainOrSnow struct {
	OneH   float64 `json:"1h"`
	ThreeH float64 `json:"3h"`
}

type RsCurrentSys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}
