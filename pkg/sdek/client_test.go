package sdek_test

import (
	"testing"

	"github.com/ReanSn0w/go-sdek/pkg/sdek"
)

var client = sdek.NewClientTest()

func TestNewClient(t *testing.T) {

	err := client.TokenRefresh()
	if err != nil {
		t.Fatal(err)
	}
}
