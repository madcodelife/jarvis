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
	WeatherEndpoint = os.Getenv("WEATHER_ENDPOINT")
	WeatherApiKey   = os.Getenv("WEATHER_API_KEY")

	// bark
	BarkEndpoints = os.Getenv("BARK_ENDPOINTS")

	// supabase
	SupabaseUrl = os.Getenv("SUPABASE_URL")
	SupabaseKey = os.Getenv("SUPABASE_KEY")
)

func Init() {
	if CI == "" {
		const projectDirName = "jarvis"
		re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
		cwd, _ := os.Getwd()
		rootPath := re.Find([]byte(cwd))

		var e map[string]string
		e, err := godotenv.Read(string(rootPath) + `/.env`)
		if err != nil {
			log.Fatal("failed to read .env file", err)
		}

		WeatherEndpoint = e["WEATHER_ENDPOINT"]
		WeatherApiKey = e["WEATHER_API_KEY"]
		BarkEndpoints = e["BARK_ENDPOINTS_TEST"]
		SupabaseUrl = e["SUPABASE_URL"]
		SupabaseKey = e["SUPABASE_KEY"]
	}
}
