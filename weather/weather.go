package weather

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"
)

func FetchWeather() (WeatherInfo, error) {
	weatherEndPoint := os.Getenv("WEATHER_END_POINT")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")

	queryParams := url.Values{}
	queryParams.Add("key", weatherApiKey)
	queryParams.Add("city", "510100")

	finalURL := weatherEndPoint + "?" + queryParams.Encode()

	resp, err := http.Get(finalURL)
	if err != nil {
		return WeatherInfo{}, errors.Wrap(err, "failed to get weather info")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherInfo{}, errors.Wrap(err, "failed to read weather info response body")
	}

	var data WeatherInfo
	if err := json.Unmarshal(body, &data); err != nil {
		return WeatherInfo{}, errors.Wrap(err, "failed to unmarshal weather info")
	}

	return data, nil
}
