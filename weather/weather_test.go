package weather

import (
	"log"
	"macodelife/jarvis/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchWeather(t *testing.T) {
	config.Init()

	liveWeather, err := fetchWeather()
	if err != nil {
		log.Fatalln("failed to fetch weather info:", err)
	}

	assert.Equal(t, liveWeather.Adcode, "510100")
	assert.Equal(t, liveWeather.City, "成都市")
}
