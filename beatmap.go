package osuapi

import (
	"encoding/json"
	"fmt"
)

// These are the statuses an osu! beatmap can have.
const (
	Graveyard = iota - 2
	WIP
	Pending
	Ranked
	Approved
	Qualified
)

// Any is shared between languages and genre, so that's why it's outside both blocks.
const Any = 0

// These are the Language IDs a beatmap can have.
const (
	LanguageOther = iota + 1
	English
	Japanese
	Chinese
	Instrumental
	Korean
	French
	German
	Swedish
	Spanish
	Italian
)

// There are the various genres a beatmap can have as a GenreID.
const (
	Unspecified = iota + 1
	Videogame
	Anime
	Rock
	Pop
	GenreOther
	Novelty
	HipHop
	Electronic
)

// Beatmap contains information about a beatmap difficulty.
type Beatmap struct {
	Approved     int       `json:"approved,string"`
	ApprovedDate MySQLDate `json:"approved_date"`
	Artist       string    `json:"artist"`
	BeatmapID    int       `json:"beatmap_id,string"`
	BeatmapsetID int       `json:"beatmapset_id,string"`
	Bpm          float64   `json:"bpm,string"`
	Creator      string    `json:"creator"`
	// We're using float32 here because it's just one decimal digit who cares
	ApproachRate      float32   `json:"diff_approach,string"`
	HPDrain           float32   `json:"diff_drain,string"`
	OverallDifficulty float32   `json:"diff_overall,string"`
	CircleSize        float32   `json:"diff_size,string"`
	Difficulty        float64   `json:"difficultyrating,string"`
	FavouriteCount    int       `json:"favourite_count,string"`
	FileMD5           string    `json:"file_md5"`
	GenreID           int       `json:"genre_id,string"`
	HitLength         int       `json:"hit_length,string"`
	LanguageID        int       `json:"language_id,string"`
	LastUpdate        MySQLDate `json:"last_update,string"`
	MaxCombo          int       `json:"max_combo,string"`
	Mode              int       `json:"mode,string"`
	Passcount         int       `json:"passcount,string"`
	Playcount         int       `json:"playcount,string"`
	Source            string    `json:"source"`
	Tags              string    `json:"tags"`
	Title             string    `json:"title"`
	TotalLength       int       `json:"total_length,string"`
	DiffName          string    `json:"version"`
}

// Welcome to hell.

// GetBeatmapByDiffID returns a beatmap by searching its ID (or Diff ID), which is located in https://osu.ppy.sh/b/<ID>.
func (a *APIClient) GetBeatmapByDiffID(id int) (Beatmap, error) {
	beatmaps, err := a.GetBeatmapFull(MySQLDate{}, 0, id, "", "", -1, 0, "", 0)
	if err != nil {
		return Beatmap{}, err
	}
	if len(beatmaps) == 0 {
		return Beatmap{}, nil
	}
	return beatmaps[0], nil
}

// GetBeatmapFull allows you to lookup for a beatmap using all the search options provided by the API.
//
// If you want to use this, please make sure there's no other way to do what you are doing, as this function is
// extremely unreadable when written as is.
func (a *APIClient) GetBeatmapFull(
	// since is a MySQLDate containing the last time the beatmap was updated. Ignore value: MySQLDate{}.
	since MySQLDate,
	// set is the ID of the beatmap set (http://osu.ppy.sh/s/<id>). Ignore value: 0.
	set int,
	// diff is the ID of a beatmap (http://osu.ppy.sh/b/<id>). Ignore value: 0.
	beatmapID int,
	// user is the username or user ID of the beatmap(s) creator. Ignore value: "".
	username string,
	// usernameType is either "id" or "string", and it specifies accurately whether username is, in fact, a username
	// or it is an user ID. Ignore value: "".
	usernameType string,
	// Beatmap mode. Ignore value: -1.
	mode int,
	// Specifies whether beatmaps converted from osu! standard should be included, if the specified mode
	// isn't Standard. Ignore value: 0.
	includeConverted int,
	// If you're looking for a specific beatmap with a known hash, you should write it here. Ignore value: "".
	md5hash string,
	// Limit of results to give in a page, range 1-500. Ignore value: 0.
	limit int,
) (b []Beatmap, retErr error) {
	b = []Beatmap{}
	retErr = nil

	// SANITY CHECKS
	if err := checkUsernameType(usernameType); err != nil {
		retErr = err
		return
	}
	if mode < -1 || mode > 3 {
		retErr = fmt.Errorf("invalid gamemode %d", mode)
		return
	}
	if includeConverted != 0 && includeConverted != 1 {
		retErr = fmt.Errorf("includeConverted must be either 0 or 1")
		return
	}
	if limit < 0 || limit > 500 {
		retErr = fmt.Errorf("limit must be in range 1-500 or 0 to disable")
		return
	}

	// QUERY STRING PREPARATION
	qs := map[string]string{
		"a": itos(includeConverted),
	}
	if !since.GetTime().IsZero() {
		qs["since"] = since.String()
	}
	if set != 0 {
		qs["s"] = itos(set)
	}
	if beatmapID != 0 {
		qs["b"] = itos(beatmapID)
	}
	if username != "" {
		qs["u"] = username
	}
	if usernameType != "" {
		qs["type"] = usernameType
	}
	if mode != -1 {
		qs["m"] = itos(mode)
	}
	if md5hash != "" {
		qs["hash"] = md5hash
	}
	if limit != 0 {
		qs["limit"] = itos(limit)
	}

	// REQUEST SENDING
	backJSON, err := a.makeRequest("get_beatmaps", qs)
	if err != nil {
		retErr = err
		return
	}
	err = json.Unmarshal(backJSON, &b)
	if err != nil {
		retErr = fmt.Errorf("There was an error unmarshaling the returned JSON from the osu! API. %v", err)
		return
	}
	return
}
