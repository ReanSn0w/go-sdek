package sdek

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httputil"
)

func (c *Client) get(method string, res, resErr interface{}) (int, error) {
	code, body, err := c.http(http.MethodGet, method, nil)
	if err != nil {
		return 0, err
	}
	return c.res(code, body, res, resErr)
}

func (c *Client) post(method string, req, res, resErr interface{}) (int, error) {
	code, body, err := c.http(http.MethodPost, method, req)
	if err != nil {
		return 0, err
	}
	return c.res(code, body, res, resErr)
}

func (c *Client) res(code int, body []byte, res, resErr interface{}) (int, error) {

	var err error
	if code == 200 || code == 202 {
		// Ok
		err = json.Unmarshal(body, res)
		if err != nil {
			return 0, err
		}
		return code, nil
	} else {
		// ERR
		err = json.Unmarshal(body, resErr)
		if err != nil {
			return 0, err
		}
		c.log.Logf("[ERROR] client.POST", string(body))
		return code, nil
	}
}

func (c *Client) http(httpMethod string, method string, reqJson interface{}) (int, []byte, error) {

	var bodyReq []byte
	if reqJson != nil {
		bodyReq, _ = json.MarshalIndent(reqJson, "", "    ")
	}

	u := c.endPoint + method
	req, err := http.NewRequest(httpMethod, u, bytes.NewReader(bodyReq))
	if err != nil {
		return 0, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json;charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.token)

	reqDump, _ := httputil.DumpRequest(req, true)
	c.logger.Logf("[DEBUG] REQ:\n%s\n\n", string(reqDump))
	if err != nil {
		return 0, nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer func() { _ = res.Body.Close() }()

	resDump, _ := httputil.DumpResponse(res, true)
	c.logger.Logf("[DEBUG] RES:\n%s\n\n", string(resDump))
	if err != nil {
		return 0, nil, err
	}

	body, _ := io.ReadAll(res.Body)

	return res.StatusCode, body, err
}
