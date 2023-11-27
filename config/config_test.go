package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	Init()

	assert.NotNil(t, WeatherEndpoint)
	assert.NotNil(t, WeatherApiKey)
	assert.NotNil(t, BarkEndpoints)
}
