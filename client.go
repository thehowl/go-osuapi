package osuapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is an osu! API client that is able to make API requests.
type Client struct {
	client *http.Client
	key    string
}

// NewClient generates a new Client based on an API key.
func NewClient(key string) *Client {
	c := &Client{&http.Client{}, key}
	return c
}

func (c Client) makerq(endpoint string, queryString url.Values) (*http.Response, error) {
	queryString.Set("k", c.key)
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/"+endpoint+"?"+queryString.Encode(), nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

type testResponse struct {
	Error string `json:"error"`
}

// Test makes sure the API is working (and the API key is valid).
func (c Client) Test() error {
	resp, err := c.makerq("get_user", url.Values{
		"u": []string{
			"2",
		},
	})
	if err != nil {
		return err
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var tr testResponse
	err = json.Unmarshal(respData, &tr)
	// Ignore cannot unmarshal stuff
	if err != nil && err.Error() != "json: cannot unmarshal array into Go value of type osuapi.testResponse" {
		return err
	}
	if tr.Error != "" {
		return errors.New("osu! API response: " + tr.Error)
	}
	return nil
}
