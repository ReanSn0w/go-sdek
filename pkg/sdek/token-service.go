package sdek

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"
)

func (c *Client) TokenRefresh() error {

	method := "oauth/token?parameters"

	vs := url.Values{}
	vs.Add("grant_type", "client_credentials")
	vs.Add("client_id", c.auth.clientId)
	vs.Add("client_secret", c.auth.clientSecret)

	req, err := http.NewRequest(http.MethodPost, c.endPoint+method, strings.NewReader(vs.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json;charset=utf-8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	reqDump, _ := httputil.DumpRequest(req, true)
	c.logger.Logf("[DEBUG] REQ:\n%s\n\n", string(reqDump))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()

	resDump, _ := httputil.DumpResponse(res, true)
	log.Printf("[DEBUG] RES:\n%s\n\n", string(resDump))

	body, _ := io.ReadAll(res.Body)
	var token TokenRes
	var tokenErr *TokenErr

	if res.StatusCode == 200 {
		err = json.Unmarshal(body, &token)
		if err != nil {
			return err
		}
		c.token = token.AccessToken
	} else {
		err = json.Unmarshal(body, &tokenErr)
		if err != nil {
			return err
		}
	}

	if !(reflect.ValueOf(tokenErr).Kind() == reflect.Ptr && reflect.ValueOf(tokenErr).IsNil()) {
		err = tokenErr
	}

	return err
}
