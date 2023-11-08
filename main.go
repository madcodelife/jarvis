package main

import (
	"macodelife/weather-cli/config"
	"macodelife/weather-cli/days"
	"macodelife/weather-cli/weather"
)

func main() {
	config.Init()

	go days.Push()
	weather.Push()
}
