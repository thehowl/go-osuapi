package osuapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Match is a multiplayer match.
type Match struct {
	Info  MatchInfo   `json:"match"`
	Games []MatchGame `json:"games"`
}

// MatchInfo contains useful information about a Match.
type MatchInfo struct {
	MatchID   int        `json:"match_id,string"`
	Name      string     `json:"name"`
	StartTime MySQLDate  `json:"start_time"`
	EndTime   *MySQLDate `json:"end_time"`
}

// MatchGame is a single beatmap played in the Match.
type MatchGame struct {
	GameID    int        `json:"game_id,string"`
	StartTime MySQLDate  `json:"start_time"`
	EndTime   *MySQLDate `json:"end_time"`
	BeatmapID int        `json:"beatmap_id,string"`
	PlayMode  Mode       `json:"play_mode,string"`
	// Refer to the wiki for information about what the three things below
	// are. I personally think that this information wouldn't be that
	// necessary to most people, and what is written on the wiki could be
	// outdated, thus I deemed useless making an appropriate "enum" (like
	// I did for Genre, Language and that stuff.)
	// You really can say I love writing long comments.
	MatchType   int              `json:"match_type,string"`
	ScoringType int              `json:"scoring_type,string"`
	TeamType    int              `json:"team_type,string"`
	Mods        Mods             `json:"mods,string"`
	Scores      []MatchGameScore `json:"scores"`
}

// MatchGameScore is a single score done by an user in a specific Game of a
// Match. I agree, these descriptions are quite confusing.
type MatchGameScore struct {
	Slot     int   `json:"slot,string"`
	Team     int   `json:"team,string"`
	UserID   int   `json:"user_id,string"`
	Score    int64 `json:"score,string"`
	MaxCombo int   `json:"maxcombo,string"`
	// There should be Rank here, but Rank is not actually used. (always 0)
	Count50   int `json:"count50,string"`
	Count100  int `json:"count100,string"`
	Count300  int `json:"count300,string"`
	CountMiss int `json:"countmiss,string"`
	CountGeki int `json:"countgeki,string"`
	CountKatu int `json:"countkatu,string"`
	// There should also be Perfect here, but that seems to also not be used. (always 0)
	Pass OsuBool `json:"pass"`
}

// GetMatch makes a get_match request to the osu! API.
func (c Client) GetMatch(matchID int) (*Match, error) {
	vals := url.Values{
		"mp": []string{strconv.Itoa(matchID)},
	}
	rawData, err := c.makerq("get_match", vals)
	if err != nil {
		return nil, err
	}
	match := Match{}
	err = json.Unmarshal(rawData, &match)
	if err != nil {
		return nil, err
	}
	return &match, nil
}
