package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	host, err := os.Hostname()
	if err != nil {
		host = "undefined"
	}

	app := NewApp()
	assert.Equal(t, app.Hostname, host)
}

