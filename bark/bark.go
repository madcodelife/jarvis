package bark

import (
	"bytes"
	"encoding/json"
	"log"
	"macodelife/weather-cli/config"
	"net/http"
	"strings"
)

func Push(b *BarkParams) {
	log.Println("push bark message:", b)

	level := LevelTimeSensitive
	b.Level = &level

	jsonData, _ := json.Marshal(b)

	barkEndPoints := strings.Split(config.BarkEndPoint, ",")

	for _, endpoint := range barkEndPoints {
		send(jsonData, endpoint)
	}
}

func send(jsonData []byte, endpoint string) {
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("failed to push bark message:", err)
	}
	defer resp.Body.Close()
}
