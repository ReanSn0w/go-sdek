package sdek

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
	return newClient(ENDPOINT, clientId, clientSecret)
}

// NewClientTest returns a new CDEK API client for test.
func NewClientTest(logger Logger) (*Client, error) {
	clientId, clientSecret := "epT5FMOa7IwjjlwTc1gUjO1GZDH1M1rE", "cYxOu9iAMZYQ1suEqfEvsHld4YQzjY0X"
	return newClient(ENDPOINT_TEST, clientId, clientSecret)
}

func newClient(ENDPOINT, clientId, clientSecret string) (*Client, error) {
	client := &Client{
		endPoint: ENDPOINT,
		auth: &auth{
			clientId:     clientId,
			clientSecret: clientSecret,
		},
	}

	err := client.TokenRefresh()
	return client, err
}
