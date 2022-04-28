package model

// GeocodingResponse : a type alias to a collection of locations
type GeocodingResponse []Location

type Location struct {
	Name    string
	Country string
	State   string
	Lat     float32
	Lon     float32
}

// WeatherResponse : a struct that represents a response from the OpenWeather Current Weather API
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
