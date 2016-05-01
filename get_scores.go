package osuapi

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// GetScoresOpts is a struct containing the GET query string parameters to an
// /api/get_scores request.
type GetScoresOpts struct {
	BeatmapID int
	// As usual, if both UserID and Username are set, UserID will override Username.
	UserID   int
	Username string
	Mode     Mode
	Mods     *Mods // Pointer because must have the possibility to be 0 (nomod) but also nil (whatever is fine)
	Limit    int
}

// Score is a top 100 score from an osu! beatmap.
type Score struct {
	ScoreID   int64     `json:"score_id,string"`
	Score     int64     `json:"score,string"`
	Username  string    `json:"username"`
	MaxCombo  int       `json:"maxcombo,string"`
	Count50   int       `json:"count50,string"`
	Count100  int       `json:"count100,string"`
	Count300  int       `json:"count300,string"`
	CountMiss int       `json:"countmiss,string"`
	CountKatu int       `json:"countkatu,string"`
	CountGeki int       `json:"countgeki,string"`
	FullCombo OsuBool   `json:"perfect,string"`
	Mods      Mods      `json:"enabled_mods,string"`
	UserID    int       `json:"user_id,string"`
	Date      MySQLDate `json:"date"`
	Rank      string    `json:"rank"` // Rank = SSH, SS, SH, S, A, B, C, D
	PP        float64   `json:"pp,string"`
}

// GetScores makes a get_scores request to the osu! API.
func (c Client) GetScores(opts GetScoresOpts) ([]Score, error) {
	// setup of querystring values
	vals := url.Values{}
	if opts.BeatmapID == 0 {
		return nil, errors.New("osuapi: BeatmapID must be set in GetScoresOpts")
	}
	vals.Add("b", strconv.Itoa(opts.BeatmapID))
	switch {
	case opts.UserID != 0:
		vals.Add("u", strconv.Itoa(opts.UserID))
		vals.Add("type", "id")
	case opts.Username != "":
		vals.Add("u", opts.Username)
		vals.Add("type", "string")
	}
	vals.Add("m", strconv.Itoa(int(opts.Mode)))
	if opts.Mods != nil {
		vals.Add("mods", strconv.Itoa(int(*opts.Mods)))
	}
	if opts.Limit != 0 {
		vals.Add("limit", strconv.Itoa(opts.Limit))
	}

	// actual request
	rawData, err := c.makerq("get_scores", vals)
	if err != nil {
		return nil, err
	}
	scores := []Score{}
	err = json.Unmarshal(rawData, &scores)
	if err != nil {
		return nil, err
	}
	return scores, nil
}
