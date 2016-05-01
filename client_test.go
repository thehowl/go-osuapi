package osuapi

import (
	"os"
	"testing"
)

var apiKey string

func TestSetKey(t *testing.T) {
	apiKey = os.Getenv("OSU_API_KEY")
	if apiKey == "" {
		t.Fatal("OSU_API_KEY was not set. All tests that require a connection to the osu! server will fail.")
	}
}

// ck checks that the apikey is set, if not it immediately fails the test.
func ck(t *testing.T) {
	if apiKey == "" {
		t.Fatal("no api key")
	}
}

func TestTestClient(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	err := c.Test()
	if err != nil {
		t.Fatal(err)
	}
}
