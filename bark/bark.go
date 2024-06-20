package bark

import (
	"bytes"
	"encoding/json"
	"log"
	"macodelife/jarvis/config"
	"net/http"
	"strings"
	"sync"
)

func Push(b *BarkParams) {
	log.Println("push bark message:", b)

	b.Level = LevelTimeSensitive
	b.Icon = "https://p.madcodelife.com/blog/2024/06/2a020e187c1252a3c59c70c6ff29f8f8.jpeg"

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
