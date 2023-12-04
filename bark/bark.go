package bark

import (
	"bytes"
	"encoding/json"
	"log"
	"macodelife/weather-cli/config"
	"net/http"
	"strings"
	"sync"
)

func Push(b *BarkParams) {
	log.Println("push bark message:", b)

	b.Level = LevelTimeSensitive
	b.Icon = "https://res.cloudinary.com/dspnhl2nc/image/upload/v1701326671/jarvis/77_ukxrzn.jpg"

	jsonData, _ := json.Marshal(b)

	barkEndpoints := strings.Split(config.BarkEndpoints, ",")

	var wg sync.WaitGroup

	for _, endpoint := range barkEndpoints {
		wg.Add(1)

		makeRequest(jsonData, endpoint, &wg)
	}

	wg.Wait()
}

func makeRequest(jsonData []byte, endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("failed to push bark message:", err)
	}

	defer resp.Body.Close()

	log.Printf("%s pushed", endpoint)
}
