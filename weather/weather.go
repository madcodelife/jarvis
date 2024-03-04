package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"macodelife/jarvis/bark"
	"macodelife/jarvis/config"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

func fetchWeather() (WeatherForecasts, error) {
	queryParams := url.Values{}
	queryParams.Add("key", config.WeatherApiKey)
	queryParams.Add("city", "510100")
	queryParams.Add("extensions", "all")

	finalURL := config.WeatherEndpoint + "?" + queryParams.Encode()
	resp, err := http.Get(finalURL)
	if err != nil {
		return WeatherForecasts{}, errors.Wrap(err, "failed to get weather info")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherForecasts{}, errors.Wrap(err, "failed to read weather info response body")
	}

	var data WeatherInfo
	if err := json.Unmarshal(body, &data); err != nil {
		return WeatherForecasts{}, errors.Wrap(err, "failed to unmarshal weather info")
	}

	return data.Forecasts[0], nil
}

func Push() {
	weatherForecasts, err := fetchWeather()
	if err != nil {
		log.Fatalln("failed to fetch weather info:", err)
	}

	weatherCasts := weatherForecasts.Casts[0]

	var title string
	if weatherCasts.Dayweather == weatherCasts.Nightweather {
		title = weatherCasts.Dayweather
	} else {
		title = fmt.Sprintf("%s 转 %s", weatherCasts.Dayweather, weatherCasts.Nightweather)
	}

	var body string
	dayTemp, _ := strconv.Atoi(weatherCasts.Daytemp)
	nightTemp, _ := strconv.Atoi(weatherCasts.Nighttemp)
	if dayTemp > nightTemp {
		body = fmt.Sprintf("%s°C - %s°C", weatherCasts.Nighttemp, weatherCasts.Daytemp)
	} else {
		body = fmt.Sprintf("%s°C - %s°C", weatherCasts.Daytemp, weatherCasts.Nighttemp)
	}

	d := bark.BarkParams{
		Title: fmt.Sprintf("☁️ 今日天气「%s」", title),
		Body:  fmt.Sprintf("温度 %s", body),
	}
	bark.Push(&d)
}
