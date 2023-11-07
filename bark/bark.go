package bark

import (
	"bytes"
	"encoding/json"
	"log"
	"macodelife/weather-cli/config"
	"net/http"
)

func Push(b *BarkParams) {
	log.Println("push bark message:", b)

	jsonData, _ := json.Marshal(b)

	resp, err := http.Post(config.BarkEndPoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("failed to push bark message:", err)
	}
	defer resp.Body.Close()
}
