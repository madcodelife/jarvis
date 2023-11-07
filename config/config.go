package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

var (
	// base
	CI = os.Getenv("CI")

	// weather
	WeatherEndPoint = os.Getenv("WEATHER_END_POINT")
	WeatherApiKey   = os.Getenv("WEATHER_API_KEY")

	// bark
	BarkEndPoint = os.Getenv("BARK_END_POINT")
)

const projectDirName = "jarvis"

func Init() {
	if CI == "" {
		re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
		cwd, _ := os.Getwd()
		rootPath := re.Find([]byte(cwd))

		var e map[string]string
		e, err := godotenv.Read(string(rootPath) + `/.env`)
		if err != nil {
			log.Fatal("failed to read .env file", err)
		}

		WeatherEndPoint = e["WEATHER_END_POINT"]
		WeatherApiKey = e["WEATHER_API_KEY"]
		BarkEndPoint = e["BARK_END_POINT"]
	}

}
