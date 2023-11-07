package main

import (
	"fmt"
	"log"
	"macodelife/weather-cli/bark"
	"macodelife/weather-cli/config"
	"macodelife/weather-cli/weather"
)

func main() {
	config.Init()

	liveWeather, err := weather.FetchWeather()
	if err != nil {
		log.Fatalln("failed to fetch weather info:", err)
	}

	d := bark.BarkParams{
		Title: "今日天气",
		Body:  fmt.Sprintf("今日天气「%s」，温度 %s°C", liveWeather.Weather, liveWeather.Temperature),
	}

	bark.Push(&d)
}
