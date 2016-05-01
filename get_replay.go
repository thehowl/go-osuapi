package osuapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strconv"
)

// GetReplayOpts are the options that MUST be used to fetch a replay.
// ALL of the fields are **REQUIRED**, with an exception for UserID/Username,
// of which only one is required.
type GetReplayOpts struct {
	UserID    int
	Username  string
	Mode      Mode
	BeatmapID int
}

type replayResponse struct {
	Content string `json:"content"`
}

// GetReplay makes a get_replay request to the osu! API. Returns a reader from
// which the replay can be retrieved.
func (c Client) GetReplay(opts GetReplayOpts) (io.Reader, error) {
	vals := url.Values{}
	if opts.BeatmapID == 0 {
		return nil, errors.New("osuapi: BeatmapID MUST be set in GetReplayOpts")
	}
	vals.Add("m", strconv.Itoa(int(opts.Mode)))
	vals.Add("b", strconv.Itoa(opts.BeatmapID))
	switch {
	case opts.UserID != 0:
		vals.Add("u", strconv.Itoa(opts.UserID))
		vals.Add("type", "id")
	case opts.Username != "":
		vals.Add("u", opts.Username)
		vals.Add("type", "string")
	default:
		return nil, errors.New("osuapi: either UserID or Username MUST be set in GetReplayOpts")
	}
	data, err := c.makerq("get_replay", vals)
	if err != nil {
		return nil, err
	}
	rr := replayResponse{}
	err = json.Unmarshal(data, &rr)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewBuffer([]byte(rr.Content))
	return base64.NewDecoder(base64.StdEncoding, reader), nil
}
