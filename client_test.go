package osuapi

import (
	"os"
	"strings"
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
		t.Skip("no api key")
	}
}

// fe returns err.Error(), with the API key removed
func fe(err error) string {
	return strings.Replace(err.Error(), apiKey, "xxxxxx", -1)
}

func TestTestClient(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	err := c.Test()
	if err != nil {
		t.Fatal(fe(err))
	}
}

func TestGetUser(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	_, err := c.GetUser(GetUserOpts{
		Username:  "Loctav",
		Mode:      ModeTaiko,
		EventDays: 4,
	})
	if err != nil && err != ErrNoSuchUser {
		t.Fatal(fe(err))
	}
}

func TestGetBeatmaps(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	_, err := c.GetBeatmaps(GetBeatmapsOpts{
		BeatmapSetID: 332532,
	})
	if err != nil {
		t.Fatal(fe(err))
	}
}

func TestGetScores(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	_, err := c.GetScores(GetScoresOpts{
		BeatmapID: 736213,
	})
	if err != nil {
		t.Fatal(fe(err))
	}
}

func TestGetUserBest(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	_, err := c.GetUserBest(GetUserScoresOpts{
		UserID: 2,
	})
	if err != nil {
		t.Fatal(fe(err))
	}
}

func TestGetUserRecent(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	_, err := c.GetUserRecent(GetUserScoresOpts{
		Username: "Cookiezi",
	})
	if err != nil {
		t.Fatal(fe(err))
	}
}

func TestGetMatch(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	_, err := c.GetMatch(20138460)
	if err != nil {
		t.Fatal(fe(err))
	}
}

func TestGetReplay(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	replayReader, err := c.GetReplay(GetReplayOpts{
		Username:  "rrtyui",
		BeatmapID: 131891,
	})
	if err != nil {
		t.Fatal(fe(err))
	}
	d := make([]byte, 16)
	_, err = replayReader.Read(d)
	if err != nil {
		t.Fatal(fe(err))
	}
	t.Logf("rrtyui on the big black: %x", d)
}
