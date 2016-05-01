package osuapi

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// GetUserScoresOpts are the options that can be passed in GetUserBest or
// in GetUserRecent; they use the same parameters.
type GetUserScoresOpts struct {
	// You know it by now. UserID overrides Username.
	UserID   int
	Username string
	Mode     Mode
	Limit    int
}

// GUSScore is a score from get_user_best or get_user_recent. It differs from
// normal Score by having a BeatmapID field. Stands for Get User Scores Score.
// Yeah, I suck at choosing names. but that's what programming is all about,
// after all.
type GUSScore struct {
	BeatmapID int `json:"beatmap_id,string"`
	Score
}

func (o GetUserScoresOpts) toValues() url.Values {
	vals := url.Values{}
	switch {
	case o.UserID != 0:
		vals.Add("u", strconv.Itoa(o.UserID))
		vals.Add("type", "id")
	case o.Username != "":
		vals.Add("u", o.Username)
		vals.Add("type", "string")
	}
	vals.Add("m", strconv.Itoa(int(o.Mode)))
	if o.Limit != 0 {
		vals.Add("limit", strconv.Itoa(o.Limit))
	}
	return vals
}

// GetUserBest makes a get_user_best request to the osu! API.
func (c Client) GetUserBest(opts GetUserScoresOpts) ([]GUSScore, error) {
	if opts.UserID == 0 && opts.Username == "" {
		return nil, errors.New("osuapi: must have either UserID or Username in GetUserScoresOpts")
	}

	rawData, err := c.makerq("get_user_best", opts.toValues())
	if err != nil {
		return nil, err
	}
	scores := []GUSScore{}
	err = json.Unmarshal(rawData, &scores)
	if err != nil {
		return nil, err
	}
	return scores, nil
}
