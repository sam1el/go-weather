// filepath: /workspaces/go-weather/cmd/weather/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"go-weather/internal/weather"
)

func main() {
	zipcode := flag.String("zipcode", "", "Zipcode for the weather information")
	flag.Parse()

	if *zipcode == "" {
		fmt.Println("Please provide a zipcode using the -zipcode flag.")
		os.Exit(1)
	}

	weatherInfo, err := weather.GetWeather(*zipcode)
	if err != nil {
		log.Fatalf("Error fetching weather: %v", err)
	}

	fmt.Println(weatherInfo)
}
