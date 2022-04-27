package model

type GeocodingResponse []Location

type Location struct {
	Name    string
	Country string
	State   string
	Lat     float32
	Lon     float32
}

type WeatherResponse struct {
	Main Weather
}

type Weather struct {
	Temp      float32
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int
	Humidity  int
}
