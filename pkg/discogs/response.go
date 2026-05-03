package discogs

import (
	"encoding/json"
	"net/http"
)

func (c *Client) Execute(req *http.Request, v interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return parseError(resp)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
