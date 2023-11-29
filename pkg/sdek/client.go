package sdek

import (
	"time"

	"github.com/ReanSn0w/go-sdek/pkg/tools"
)

const (
	ENDPOINT      = "https://api.cdek.ru/v2/"
	ENDPOINT_TEST = "https://api.edu.cdek.ru/v2/"
)

type (
	Client struct {
		auth     *auth
		endPoint string
		token    string
		logger   Logger
		cycle    *tools.CycleTask
	}

	auth struct {
		clientId     string
		clientSecret string
	}

	Logger interface {
		Logf(format string, args ...interface{})
	}
)

// NewClient returns a new CDEK API client.
func NewClient(logger Logger, clientId, clientSecret string) (*Client, error) {
	return newClient(logger, ENDPOINT, clientId, clientSecret)
}

// NewClientTest returns a new CDEK API client for test.
func NewClientTest(logger Logger) (*Client, error) {
	clientId, clientSecret := "epT5FMOa7IwjjlwTc1gUjO1GZDH1M1rE", "cYxOu9iAMZYQ1suEqfEvsHld4YQzjY0X"
	return newClient(logger, ENDPOINT_TEST, clientId, clientSecret)
}

func newClient(logger Logger, ENDPOINT, clientId, clientSecret string) (*Client, error) {
	client := &Client{
		endPoint: ENDPOINT,
		logger:   logger,
		auth: &auth{
			clientId:     clientId,
			clientSecret: clientSecret,
		},
	}

	err := client.TokenRefresh()

	client.cycle = tools.NewCycleTask(func() {
		client.TokenRefresh()
	})

	if err != nil {
		client.cycle.Run(time.Minute * 30)
	}

	return client, err
}

func (c *Client) TokenRefresher() {
	for {
		time.Sleep(time.Minute & 50)
		c.TokenRefresh()
	}
}
