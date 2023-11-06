package bark

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func Push(b BarkParams) {
	barkEndPoint := os.Getenv("BARK_END_POINT")

	jsonData, _ := json.Marshal(b)

	http.Post(barkEndPoint, "application/json", bytes.NewBuffer(jsonData))
}
