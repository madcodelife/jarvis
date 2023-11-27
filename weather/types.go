package weather

type WeatherInfo struct {
	Status    string             `json:"status"`
	Count     string             `json:"count"`
	Info      string             `json:"info"`
	Infocode  string             `json:"infocode"`
	Forecasts []WeatherForecasts `json:"forecasts"`
}

type WeatherForecasts struct {
	City       string         `json:"city"`
	Adcode     string         `json:"adcode"`
	Province   string         `json:"province"`
	Reporttime string         `json:"reporttime"`
	Casts      []WeatherCasts `json:"casts"`
}

type WeatherCasts struct {
	Date           string `json:"date"`
	Week           string `json:"week"`
	Dayweather     string `json:"dayweather"`
	Nightweather   string `json:"nightweather"`
	Daytemp        string `json:"daytemp"`
	Nighttemp      string `json:"nighttemp"`
	Daywind        string `json:"daywind"`
	Nightwind      string `json:"nightwind"`
	Daypower       string `json:"daypower"`
	Nightpower     string `json:"nightpower"`
	DaytempFloat   string `json:"daytemp_float"`
	NighttempFloat string `json:"nighttemp_float"`
}
