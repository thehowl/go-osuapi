package osuapi

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

/*
====================
     USER TESTS
====================
*/

func TestGetUser(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUser("peppy", OsuStandard)
	if err != nil {
		panic(err)
	}
	printUser(user)
}
func TestGetUserByID(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUserByID(140148, OsuMania)
	if err != nil {
		panic(err)
	}
	printUser(user)
}
func TestGetUserByUsername(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUserByUsername("Ikkun", Taiko)
	if err != nil {
		panic(err)
	}
	printUser(user)
}
func TestGetUserFull(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUserFull("803484", CatchTheBeat, "id", 20)
	if err != nil {
		panic(err)
	}
	printUser(user)
}

/*
====================
   BEATMAP TESTS
====================
*/

func TestGetBeatmapFull(t *testing.T) {
	c := testingGenClient()
	b, err := c.GetBeatmapFull(MySQLDate{}, 2, 0, "peppy", "", -1, 0, "", 0)
	if err != nil {
		panic(err)
	}
	for _, beatmap := range b {
		printBeatmap(beatmap)
	}
}
func TestGetBeatmapByDiffID(t *testing.T) {
	c := testingGenClient()
	b, err := c.GetBeatmapByDiffID(75)
	if err != nil {
		panic(err)
	}
	// We aren't iterating because in this case we are getting it by the primary key of the database, thus GetBeatmapByDiffID doesn't return an array but a single beatmap.
	printBeatmap(b)
}
func TestGetBeatmapBySetID(t *testing.T) {
	c := testingGenClient()
	b, err := c.GetBeatmapBySetID(240926)
	if err != nil {
		panic(err)
	}
	for _, beatmap := range b {
		printBeatmap(beatmap)
	}
}
func TestGetBeatmapByUser(t *testing.T) {
	c := testingGenClient()
	b, err := c.GetBeatmapByUser("Howl")
	if err != nil {
		panic(err)
	}
	for _, beatmap := range b {
		printBeatmap(beatmap)
	}
}
func TestGetBeatmapByUserID(t *testing.T) {
	c := testingGenClient()
	b, err := c.GetBeatmapByUserID(447818)
	if err != nil {
		panic(err)
	}
	for _, beatmap := range b {
		printBeatmap(beatmap)
	}
}
func TestGetBeatmapByUserWithType(t *testing.T) {
	c := testingGenClient()
	b, err := c.GetBeatmapByUserWithType("1273955", "id")
	if err != nil {
		panic(err)
	}
	for _, beatmap := range b {
		printBeatmap(beatmap)
	}
}

/*
====================
     HELPERS
====================
*/

func printUser(u User) {
	fmt.Printf(`Username: %v
	ID: %v
	Count300: %v
	Count100: %v
	Count50: %v
	PlayCount: %v
	RankedScore: %v
	TotalScore: %v
	Rank: %v
	Level: %v
	PP: %v
	Accuracy: %v
	CountRankSS: %v
	CountRankS: %v
	CountRankA: %v
	Country: %v
	CountryRank: %v
	Events: %v
`, u.Username, u.ID, u.Count300, u.Count100, u.Count50, u.PlayCount, u.RankedScore, u.TotalScore, u.Rank, u.Level, u.PP, u.Accuracy, u.CountRankSS, u.CountRankS, u.CountRankA, u.Country, u.CountryRank, u.Events)
}
func printBeatmap(b Beatmap) {
	fmt.Printf(`Beatmap: %s - %s [%s] (%s)
	Total length: %d
	Mode: %d
	Tags: %s
	Source: %s
`, b.Artist, b.Title, b.DiffName, b.Creator, b.TotalLength, b.Mode, b.Tags, b.Source)
}
func testingGenClient() *APIClient {
	data, err := ioutil.ReadFile("osukey.txt")
	if err != nil {
		panic(err)
	}
	k := strings.Trim(string(data), "\r\n\t ")
	return NewClient(k)
}
