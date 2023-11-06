package weather

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
