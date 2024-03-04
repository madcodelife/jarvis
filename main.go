package main

import (
	"macodelife/jarvis/config"
	"macodelife/jarvis/days"
	"macodelife/jarvis/weather"
)

func main() {
	config.Init()

	go days.Push()
	weather.Push()
}
