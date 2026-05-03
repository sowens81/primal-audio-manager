package discogs

import (
	"fmt"
	"net/http"
)

func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Discogs token="+c.Token)
	req.Header.Set("User-Agent", "primal-audio-manager/1.0")

	return req, nil
}
