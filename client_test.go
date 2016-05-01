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
		t.Fatal("no api key")
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
	user, err := c.GetUser(GetUserOpts{
		Username:  "Loctav",
		Mode:      ModeTaiko,
		EventDays: 4,
	})
	if err != nil && err != ErrNoSuchUser {
		t.Fatal(fe(err))
	}
	if user != nil {
		t.Logf("%+v", user)
	}
}

func TestGetBeatmaps(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	beatmaps, err := c.GetBeatmaps(GetBeatmapsOpts{
		BeatmapSetID: 332532,
	})
	if err != nil {
		t.Fatal(fe(err))
	}
	t.Logf("%+v", beatmaps)
}

func TestGetScores(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	scores, err := c.GetScores(GetScoresOpts{
		BeatmapID: 736213,
	})
	if err != nil {
		t.Fatal(fe(err))
	}
	t.Logf("%+v...", scores[:5])
}

func TestGetUserBest(t *testing.T) {
	ck(t)
	c := NewClient(apiKey)
	scores, err := c.GetUserBest(GetUserScoresOpts{
		UserID: 2,
	})
	if err != nil {
		t.Fatal(fe(err))
	}
	t.Logf("%+v...", scores[:3])
}
