# how-to-go
A companion repository to practice the basics of golang and get familiarized with the go language.

## Table of Contents

* [Slides](#Slides)
* [How to use this repository](#How%20to%20use%20this%20repository)
* [Running examples](#Running%20examples)
* [So, what's Next?](#So%2C%20what%27s%20Next%3F)
* [Challenge](#Challenge)
  * [How?](#How%3F)
    * [1 - Create an API Key](#1%20-%20Create%20an%20API%20Key)
    * [2 - Check out the API Documentation](#2%20-%20Check%20out%20the%20API%20Documentation)
      * [Geocoding API](#Geocoding%20API)
      * [Current Weather API](#Current%20Weather%20API)
    * [3 - Making the API Calls](#3%20-%20Making%20the%20API%20Calls)
* [The Solution](#The%20Solution)
* [Upcoming Topics](#Upcoming%20Topics)

## Slides

Slides for this workshop are available on Google Drive [here](https://docs.google.com/presentation/d/1CZf_PU0QUsokGeG1LApA2PQcDt8zIrkGZTHvHvpHH1g/edit?usp=sharing).

## How to use this repository

First, install go from [the official website](https://go.dev/dl/).

Then, check out the repo to the following directory:
```text
~/go/src/github.com/kjkondratuk/
```

**Note** : if you forked the repo, you may have to replace `kjkondratuk` with your username.

## Running examples

All example code is located under `./examples` and organized into folders to facilitate simple building.
To execute any of the example code simply execute the following in a directory with a source file:
```text
go run <go file>
```
This will compile and execute the specified program.  Alternatively, you can leverage the included `Makefile`.  To
execute various build targets, or use the included IntelliJ/Goland run configurations (if you're using one of those IDEs)
included in the `.idea` directory.

## So, what's next?

First, let's review the basics of the language by browsing the [examples folder](./examples) to better
understand the basics of the language, then let's move on to the challenge below and build
a simple command-line application.

## Challenge

Create a command line application that does the following:
1. Accepts a user's location as a parameter (or set of parameters)
2. Makes a call to the [OpenWeather API](https://openweathermap.org/api) to get weather information for the relevant
location
3. Parses the response and prints out current weather information in a readable
format to the command line

### How?

If you've reached this section and haven't yet reviewed the contents of the examples folder, please do so.  It will
help you considerably in implementing this tool to have a basic understanding of the syntax Go uses. If you'd like to start
by creating the HTTP client, we can simply hard-code parameter values for now, and add argument parsing later.  
Likewise, if you want to start with argument parsing, you can do that and add the HTTP calls later.

For the sake of time, I have included response models that can be used to unmarshal responses from the API, included 
some examples of how to use the `http` standard library, and included a skeleton application that captures several
arguments from the supplied program arguments.

#### 1 - Create an API Key

Before we get started, we should probably create an API key by registering for an account that will allow us to make
requests to the OpenWeather API.  For our purposes, it will probably be easiest to use the Geocoding API to acquire 
geographic coordinates and then use those coordinates to fetch the weather information.

Once you are registered, locate your API key and paste it into a file called `.env`.  Our application will source the API
key from this file.

#### 2 - Check out the API Documentation

##### Geocoding API
Full documentation for the Geocoding API is available [here](https://openweathermap.org/api/geocoding-api).

Minimally, we will have to supply the following values to the API to get a response:
1. `q` - the query - formatted as `city,state_code,country_code`
2. `appid` - your API key

You should be able to unmarshal this response as follows into the provided [GeocodingResponse model](./model/response.go):
```go
    var geocodingResponse model.GeocodingResponse
	all, err := ioutil.ReadAll(resp.Body)
	// ... handle error ...
	err = json.Unmarshal(all, &geocodingResponse)
    // ... handle error ...
```

##### Current Weather API
Full documentation for the Current Weather API is available [here](https://openweathermap.org/current).

Minimally, we will have to supply the following values to the API to get a response:
1. `lat` - the latitude for the corresponding weather information
2. `lon` - the longitude for the correpsonding weather information
3. `appid` - your API key

You should be able to unmarshal this response as follows into the provided [WeatherResponse model](./model/response.go),
similarly to how you did above with the `GeocodingResponse`.

#### 3 - Making the API Calls

Fortunately with Go it is very easy to create an HTTP client and make calls with it.  The official documentation on the
topic can be found [here](https://pkg.go.dev/net/http).  But a simple example looks like this:
```go
    response, err := http.Get("http://my-application.com")
	// close the body reader on completion, to free resources
    if client != nil && client.Body != nil {
        defer client.Body.Close()
    }
	// handle an error in issuing or receiving the request
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
```

Making a request happens via the following general procedure:
1. Issue the request to the specified URL
2. `defer` the closure of the body, if applicable (see above)
3. Handle any errors that arise during the request
4. Unmarshal the response JSON to a go struct (refer to above details about the APIs for more info on this)
5. Handle any errors that arise from the unmarshalling
6. Use the returned struct!

### What should my output look like?

You can format your output however you'd like, but something like the following might be appropriate:
```text
--- Today's Weather ---
 - Low: 34.049999
 - High 46.619999
 - Current: 38.799999
 - Feels Like: 34.340000
 - Humidity: 62%
 - Pressure: 1031hPa
```

## The Solution

If at any point you find yourself irreversibly stuck, check out the `solution` branch of this repository for a working
example.

## Upcoming Topics

Topics which were considered for this workshop, but will probably wind up in a future iteration of this, or added later:
* Concurrency
* Mocking
* Common standard library APIs
