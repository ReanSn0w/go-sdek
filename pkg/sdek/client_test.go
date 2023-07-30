package sdek_test

import (
	"log"
	"testing"

	"github.com/ReanSn0w/go-sdek/pkg/sdek"
)

var (
	client, _ = sdek.NewClientTest(&Log{})
)

func TestNewClient(t *testing.T) {
	err := client.TokenRefresh()
	if err != nil {
		t.Fatal(err)
	}
}

// MARK: - Implementation of logger for tests

type Log struct {
}

func (l *Log) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
