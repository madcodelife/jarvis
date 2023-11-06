package weather

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/pkg/errors"
)

type WeatherInfo struct {
	Status   string        `json:"status"`
	Count    string        `json:"count"`
	Info     string        `json:"info"`
	Infocode string        `json:"infocode"`
	Lives    []WeatherLive `json:"lives"`
}

type WeatherLive struct {
	Province         string `json:"province"`
	City             string `json:"city"`
	Adcode           string `json:"adcode"`
	Weather          string `json:"weather"`
	Temperature      string `json:"temperature"`
	WindDirection    string `json:"winddirection"`
	WindPower        string `json:"windpower"`
	Humidity         string `json:"humidity"`
	ReportTime       string `json:"reporttime"`
	TemperatureFloat string `json:"temperature_float"`
	HumidityFloat    string `json:"humidity_float"`
}

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
