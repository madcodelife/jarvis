package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"macodelife/weather-cli/bark"
	"macodelife/weather-cli/config"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func fetchWeather() (WeatherLive, error) {
	queryParams := url.Values{}
	queryParams.Add("key", config.WeatherApiKey)
	queryParams.Add("city", "510100")

	finalURL := config.WeatherEndPoint + "?" + queryParams.Encode()
	resp, err := http.Get(finalURL)
	if err != nil {
		return WeatherLive{}, errors.Wrap(err, "failed to get weather info")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherLive{}, errors.Wrap(err, "failed to read weather info response body")
	}

	var data WeatherInfo
	if err := json.Unmarshal(body, &data); err != nil {
		return WeatherLive{}, errors.Wrap(err, "failed to unmarshal weather info")
	}

	return data.Lives[0], nil
}

func Push() {
	liveWeather, err := fetchWeather()
	if err != nil {
		log.Fatalln("failed to fetch weather info:", err)
	}

	d := bark.BarkParams{
		Title: fmt.Sprintf("☁️ 今日天气「%s」", liveWeather.Weather),
		Body:  fmt.Sprintf("温度 %s°C", liveWeather.Temperature),
	}
	bark.Push(&d)
}
