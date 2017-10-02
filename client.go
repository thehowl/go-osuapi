package osuapi

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Client is an osu! API client that is able to make API requests.
type Client struct {
	client *http.Client
	key    string
}

// NewClient generates a new Client based on an API key.
func NewClient(key string) *Client {
	t := &http.Transport{
		// values for http.DefaultTransport
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		// disable HTTP/2, as it does not seem to cope well with osu!
		TLSNextProto: make(map[string]func(authority string, c *tls.Conn) http.RoundTripper),
	}
	c := &Client{&http.Client{
		Transport: t,
	}, key}
	return c
}

func (c Client) makerq(endpoint string, queryString url.Values) ([]byte, error) {
	queryString.Set("k", c.key)
	req, err := http.NewRequest("GET", "https://osu.ppy.sh/api/"+endpoint+"?"+queryString.Encode(), nil)
	if err != nil {
		return nil, err
	}
	// if we are rate limiting requests, then wait before making request
	if requestsAvailable != nil {
		<-requestsAvailable
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, err
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
	var tr testResponse
	err = json.Unmarshal(resp, &tr)
	// Ignore cannot unmarshal stuff
	if err != nil && err.Error() != "json: cannot unmarshal array into Go value of type osuapi.testResponse" {
		return err
	}
	if tr.Error != "" {
		return errors.New("osuapi: " + tr.Error)
	}
	return nil
}
