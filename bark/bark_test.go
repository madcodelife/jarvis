package bark

import (
	"macodelife/jarvis/config"
	"testing"
)

func TestPush(t *testing.T) {
	config.Init()

	Push(&BarkParams{
		Title: "我是标题",
		Body:  "我是内容",
	})
}
