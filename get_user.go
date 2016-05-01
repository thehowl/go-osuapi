package osuapi

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// ErrNoSuchUser is returned when the requested user could not be found.
var ErrNoSuchUser = errors.New("osuapi: no such user could be found")

// GetUserOpts is a struct containing the GET query string parameters to an
// /api/get_user request.
type GetUserOpts struct {
	// If both UserID and Username are set, UserID will be used.
	UserID    int
	Username  string
	Mode      Mode
	EventDays int
}

// User is an osu! user.
type User struct {
	UserID      int     `json:"user_id,string"`
	Username    string  `json:"username"`
	Count300    int     `json:"count300,string"`
	Count100    int     `json:"count100,string"`
	Count50     int     `json:"count50,string"`
	Playcount   int     `json:"playcount,string"`
	RankedScore int64   `json:"ranked_score,string"`
	TotalScore  int64   `json:"total_score,string"`
	Rank        int     `json:"pp_rank,string"`
	Level       float64 `json:"level,string"`
	PP          float64 `json:"pp_raw,string"`
	Accuracy    float64 `json:"accuracy,string"`
	CountSS     int     `json:"count_rank_ss,string"`
	CountS      int     `json:"count_rank_s,string"`
	CountA      int     `json:"count_rank_a,string"`
	Country     string  `json:"country"`
	CountryRank int     `json:"pp_country_rank,string"`
	Events      []Event `json:"events"`
}

// Event is a notorious action an user has done recently.
type Event struct {
	DisplayHTML  string    `json:"display_html"`
	BeatmapID    int       `json:"beatmap_id,string"`
	BeatmapsetID int       `json:"beatmapset_id,string"`
	Date         MySQLDate `json:"date"`
	Epicfactor   int       `json:"epicfactor,string"`
}

// GetUser makes a get_user request to the osu! API.
func (c Client) GetUser(opts GetUserOpts) (*User, error) {
	// setup of querystring values
	vals := url.Values{}
	switch {
	case opts.UserID != 0:
		vals.Add("u", strconv.Itoa(opts.UserID))
		vals.Add("type", "id")
	case opts.Username != "":
		vals.Add("u", opts.Username)
		vals.Add("type", "string")
	default:
		return nil, errors.New("osuapi: either UserID or Username must be set in GetUserOpts")
	}
	vals.Add("m", strconv.Itoa(int(opts.Mode)))
	if opts.EventDays != 0 {
		vals.Add("event_days", strconv.Itoa(opts.EventDays))
	}

	// actual request
	rawData, err := c.makerq("get_user", vals)
	if err != nil {
		return nil, err
	}
	users := []User{}
	err = json.Unmarshal(rawData, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, ErrNoSuchUser
	}
	return &users[0], nil
}

// ToGetUserOpts converts an user to a GetUserOpts, so that it can be used
// with GetUser. Note that this does not work very well. It won't auto-detect
// the game mode, because the bloody osu! API does not return that in a
// get_user response. So it will just assume you want the osu! standard data
// and return that.
func (u User) ToGetUserOpts() GetUserOpts {
	return GetUserOpts{
		UserID: u.UserID,
	}
}
