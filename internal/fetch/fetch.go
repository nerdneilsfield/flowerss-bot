package fetch

import (
	"net/http"

	"github.com/SlyMarbo/rss"
	"github.com/nerdneilsfield/flowerss-bot/pkg/client"
)

// FetchFunc rss fetch func
func FetchFunc(client *client.HttpClient) rss.FetchFunc {
	return func(url string) (resp *http.Response, err error) {
		resp, err = client.Get(url)
		if err != nil {
			return nil, err
		}
		return
	}
}
