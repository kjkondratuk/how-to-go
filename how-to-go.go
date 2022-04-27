package main

import (
	"fmt"
	"os"
)

const (
	usageText = "Please specify input in the format: <city name> <state code> <country code>"

	// These are some format strings that can be used in your code to make calls to the OpenWeather API
	geoApiTemplate     = "http://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&appid=%s"
	currentApiTemplate = "https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&units=imperial&appid=%s"
)

func main() {
	// Read our API key from the .env file
	b, _ := os.ReadFile(".env")
	apiKey := string(b)

	// ignore the program invocation, but read all the other arguments
	appArgs := os.Args[1:]
	if len(appArgs) == 3 {
		// TODO : Add your application code here!
	} else {
		// print a helpful message if there are an incorrect number of parameters
		usage()
	}
	os.Exit(0)
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
