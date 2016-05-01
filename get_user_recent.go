package osuapi

import (
	"encoding/json"
	"errors"
)

// GetUserRecent makes a get_user_recent request to the osu! API.
func (c Client) GetUserRecent(opts GetUserScoresOpts) ([]GUSScore, error) {
	if opts.UserID == 0 && opts.Username == "" {
		return nil, errors.New("osuapi: must have either UserID or Username in GetUserScoresOpts")
	}

	rawData, err := c.makerq("get_user_recent", opts.toValues())
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
