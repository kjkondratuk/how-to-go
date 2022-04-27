# how-to-go
A companion repository to practice the basics of golang and get familiarized with the go language.

## Table of Contents
[Slides](#Slides)
[How to use this repository](#How to use this repository)

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
This will compile and execute the specified program.

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
help you immensely in implementing this tool to have a basic understanding of the syntax Go uses.  If you have, pick a
section below that you'd like to start with.  There will be a couple aspects to this solution:
1. An HTTP client to send requests to the OpenWeatherAPI
2. A command-line argument parser to get the user's location

#### 1 - Creating an HTTP Client

Fortunately with Go it is very easy to create an HTTP client and make calls with it.  The official documentation on the
topic can be found [here](https://pkg.go.dev/net/http).  

#### 2 - Parsing command-line arguments



## Upcoming Topics

Topics which were considered for this workshop, but will probably wind up in a future iteration of this, or added later:
* Concurrency
* Mocking
* Common standard library APIs
