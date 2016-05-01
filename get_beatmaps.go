package osuapi

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

// GetBeatmapsOpts is a struct containing the GET query string parameters to an
// /api/get_beatmaps request.
type GetBeatmapsOpts struct {
	Since        *time.Time
	BeatmapSetID int
	BeatmapID    int
	// If both UserID and Username are set, UserID will be used.
	UserID   int
	Username string
	// Using a pointer because we MUST provide a way to make this parameter
	// optional, and it should be optional by default. This is because simply
	// doing != 0 won't work, because it wouldn't allow filtering with only
	// osu!std maps, and setting m=0 and not setting it it all makes the
	// difference between only allowing STD beatmaps and allowing beatmaps
	// from all modes.
	// Simply doing &osuapi.ModeOsuMania (or similiar) should do the trick,
	// should you need to use this field.
	// God was this comment long.
	Mode             *Mode
	IncludeConverted bool
	BeatmapHash      string
	Limit            int
}

// Beatmap is an osu! beatmap.
type Beatmap struct {
	BeatmapSetID      int            `json:"beatmapset_id,string"`
	BeatmapID         int            `json:"beatmap_id,string"`
	Approved          ApprovedStatus `json:"approved,string"`
	TotalLength       int            `json:"total_length,string"`
	HitLength         int            `json:"hit_length,string"`
	DiffName          string         `json:"version"`
	FileMD5           string         `json:"file_md5"`
	CircleSize        float64        `json:"diff_size,string"`
	OverallDifficulty float64        `json:"diff_overall,string"`
	ApproachRate      float64        `json:"diff_approach,string"`
	HPDrain           float64        `json:"diff_drain,string"`
	Mode              Mode           `json:"mode,string"`
	ApprovedDate      MySQLDate      `json:"approved_date"`
	LastUpdate        MySQLDate      `json:"last_update"`
	Artist            string         `json:"artist"`
	Title             string         `json:"title"`
	Creator           string         `json:"creator"`
	BPM               float64        `json:"bpm,string"`
	Source            string         `json:"source"`
	Tags              string         `json:"tags"`
	Genre             Genre          `json:"genre_id,string"`
	Language          Language       `json:"language_id,string"`
	FavouriteCount    int            `json:"favourite_count,string"`
	Playcount         int            `json:"playcount,string"`
	Passcount         int            `json:"passcount,string"`
	MaxCombo          int            `json:"max_combo,string"`
	DifficultyRating  float64        `json:"difficultyrating,string"`
}

// GetBeatmaps makes a get_beatmaps request to the osu! API.
func (c Client) GetBeatmaps(opts GetBeatmapsOpts) ([]Beatmap, error) {
	// setup of querystring values
	vals := url.Values{}
	switch {
	case opts.UserID != 0:
		vals.Add("u", strconv.Itoa(opts.UserID))
		vals.Add("type", "id")
	case opts.Username != "":
		vals.Add("u", opts.Username)
		vals.Add("type", "string")
	}
	if opts.Mode != nil {
		vals.Add("m", strconv.Itoa(int(*opts.Mode)))
	}
	if opts.BeatmapHash != "" {
		vals.Add("h", opts.BeatmapHash)
	}
	if opts.BeatmapID != 0 {
		vals.Add("b", strconv.Itoa(opts.BeatmapID))
	}
	if opts.BeatmapSetID != 0 {
		vals.Add("s", strconv.Itoa(opts.BeatmapSetID))
	}
	if opts.IncludeConverted {
		vals.Add("a", "1")
	}
	if opts.Since != nil {
		vals.Add("since", MySQLDate(*opts.Since).String())
	}
	if opts.Limit != 0 {
		vals.Add("limit", strconv.Itoa(opts.Limit))
	}

	// actual request
	rawData, err := c.makerq("get_beatmaps", vals)
	if err != nil {
		return nil, err
	}
	beatmaps := []Beatmap{}
	err = json.Unmarshal(rawData, &beatmaps)
	if err != nil {
		return nil, err
	}
	return beatmaps, nil
}
