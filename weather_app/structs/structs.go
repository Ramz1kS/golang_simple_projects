package structs

type City struct {
	Name string
	Lat float64
	Lon float64
	Country string
	State string
}

type WeatherInfo struct {
	Short []ShortInfo `json:"weather"`
	Main MainInfo
	Wind WindInfo
	Clouds CloudsInfo
}

type ShortInfo struct {
	Name string `json:"main"`
	Description string 
}

type WindInfo struct {
	Speed float64
	Direction float64 `json:"deg"`
	Gust float64
}

type CloudsInfo struct {
	Percantage float64 `json:"all"`
}

type MainInfo struct {
	Temp float64
	TempFeels float64 `json:"feels_like"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
	SeaLevel float64 `json:"sea_level"`
	GroundLevel float64 `json:"grnd_level"`
}