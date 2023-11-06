package main

import (
	"fmt"
	"log"
	"macodelife/weather-cli/bark"
	"macodelife/weather-cli/weather"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	CI := os.Getenv("CI")

	if CI == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to load .env file", err)
		}
	}

	weatherData, weatherErr := weather.FetchWeather()
	if weatherErr != nil {
		log.Fatalln("failed to fetch weather info:", weatherErr)
	}

	liveWeather := weatherData.Lives[0]

	d := bark.BarkParams{
		Title: "今日天气",
		Body:  fmt.Sprintf("今日天气「%s」，温度 %s°C", liveWeather.Weather, liveWeather.Temperature),
	}

	bark.Push(d)
}
