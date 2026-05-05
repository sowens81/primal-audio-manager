package discogs

import (
	"net/http"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/services"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL string
	Token   string

	HTTPClient HTTPClient

	Collection *services.CollectionService
	Release    *services.ReleaseService
}

func NewClient(token string) *Client {
	c := &Client{
		BaseURL:    "https://api.discogs.com",
		Token:      token,
		HTTPClient: http.DefaultClient, // ✅ now valid
	}

	c.Collection = services.NewCollectionService(c)
	c.Release = services.NewReleaseService(c)
	return c
}
