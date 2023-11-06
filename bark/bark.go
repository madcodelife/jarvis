package bark

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type BarkParams struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Push(b BarkParams) {
	barkEndPoint := os.Getenv("BARK_END_POINT")

	jsonData, _ := json.Marshal(b)

	http.Post(barkEndPoint, "application/json", bytes.NewBuffer(jsonData))
}
