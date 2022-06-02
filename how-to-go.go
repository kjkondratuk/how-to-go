package main

import (
	"encoding/json"
	"fmt"
	"github.com/kjkondratuk/how-to-go/model"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	usageText          = "Please specify input in the format: <city name> <state code> <country code>"
	geoApiTemplate     = "http://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&appid=%s"
	currentApiTemplate = "https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&units=imperial&appid=%s"
)

func main() {
	b, _ := os.ReadFile(".env")
	key := string(b)

	// ignore the program invocation, but read all the other arguments
	// TODO : determine what code here should be part of the "skeleton"
	if len(os.Args[1:]) == 3 {
		response := findCity(os.Args[1], os.Args[2], os.Args[3], key)

		// assume we got a single city match
		w := getWeatherForCity(key, response[0])

		fmt.Printf("--- Today's Weather ---\n - Low: %f\n - High %f\n - Current: %f\n - Feels Like: %f\n - Humidity: %d%%\n - Pressure: %dhPa",
			w.Main.TempMin, w.Main.TempMax, w.Main.Temp, w.Main.FeelsLike, w.Main.Humidity, w.Main.Pressure)
	} else {
		usage()
	}
	os.Exit(0)
}

func findCity(city string, state string, country string, apiKey string) model.GeocodingResponse {
	url := fmt.Sprintf(geoApiTemplate, city, state, country, apiKey)
	resp, err := http.Get(url)
	// TODO : add notes about cleaning up connections/resources using defer
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		exitWithError(err)
	}

	// Read the geocodingResponse body
	var geocodingResponse model.GeocodingResponse
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		exitWithError(err)
	}
	err = json.Unmarshal(all, &geocodingResponse)
	if err != nil {
		exitWithError(err)
	}

	//fmt.Println(geocodingResponse)
	return geocodingResponse
}

func getWeatherForCity(apiKey string, loc model.Location) model.WeatherResponse {
	url := fmt.Sprintf(currentApiTemplate, fmt.Sprintf("%f", loc.Lat), fmt.Sprintf("%f", loc.Lon), apiKey)
	resp, err := http.Get(url)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		exitWithError(err)
	}

	var weatherResponse model.WeatherResponse
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		exitWithError(err)
	}

	//fmt.Println(string(all))

	err = json.Unmarshal(all, &weatherResponse)
	if err != nil {
		exitWithError(err)
	}

	//fmt.Println(weatherResponse)
	return weatherResponse
}

// usage : a function that will print usage text and exit with an error code
func usage() {
	fmt.Println(usageText)
	os.Exit(-1)
}

// exit : a function that will print an error and exit the application with an error code
func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(-1)
}
