package osuapi

import (
	"fmt"
	"github.com/franela/goreq"
	"io/ioutil"
	"strings"
)

const (
	// APIEndpoint is the base URL of all the osu! API requests.
	APIEndpoint = "https://osu.ppy.sh/api/"
	// DefaultLimiting is the number of requests you are allowed to do in a minute (this is the recommended limit).
	DefaultLimiting = 60
	// DefaultLimitingReplays is the number of requests you are allowed to do in 10 minutes for get_replay.
	DefaultLimitingReplays = 10
)

// APIClient allows you to make requests to the osu! API.
type APIClient struct {
	key string
}

// NewClient generates an osu! API client through which you can make requests to the osu! API (see struct APIClient).
func NewClient(apiKey string) *APIClient {
	return &APIClient{
		key: apiKey,
	}
}

func (a *APIClient) makeRequest(req string, params map[string]string) ([]byte, error) {
	if a.key == "" {
		return []byte{}, fmt.Errorf("The API key has not been set. Perhaps you created the client manually? If so, please use the constructor NewClient.")
	}
	params["k"] = a.key
	resp, err := goreq.Request{
		Uri:         APIEndpoint + req,
		QueryString: toFuckedUp(params),
	}.Do()
	if err != nil {
		return []byte{}, fmt.Errorf("failed to do request: %v", err)
	}
	finalResp, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("What the fuck? Why did this happen? It wasn't supposed to go like this, you know... %v", err)
	}
	if strings.Trim(string(finalResp), "\r\n\t ") == `{"error":"Please provide a valid API key."}` {
		return []byte{}, fmt.Errorf("invalid API key")
	}
	return finalResp, nil
}
