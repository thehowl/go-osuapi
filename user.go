package osuapi

import (
	"encoding/json"
	"fmt"
	"time"
)

// User contains information about a user. Ezpz.
type User struct {
	Accuracy    float64 `json:"accuracy,string"`
	Count100    int     `json:"count100,string"`
	Count300    int     `json:"count300,string"`
	Count50     int     `json:"count50,string"`
	CountRankA  int     `json:"count_rank_a,string"`
	CountRankS  int     `json:"count_rank_s,string"`
	CountRankSS int     `json:"count_rank_ss,string"`
	Country     string  `json:"country"`
	Events      []struct {
		BeatmapID    int       `json:"beatmap_id,string"`
		BeatmapsetID int       `json:"beatmapset_id,string"`
		DateRaw      string    `json:"date"`
		Date         time.Time `json:"-"`
		DisplayHTML  string    `json:"display_html"`
		Epicfactor   int       `json:"epicfactor,string"`
	} `json:"events"`
	Level       float64 `json:"level,string"`
	PlayCount   int     `json:"playcount,string"`
	CountryRank int     `json:"pp_country_rank,string"`
	PP          float64 `json:"pp_raw,string"`
	Rank        int     `json:"pp_rank,string"`
	RankedScore int64   `json:"ranked_score,string"`
	TotalScore  int64   `json:"total_score,string"`
	ID          int     `json:"user_id,string"`
	Username    string  `json:"username"`
}

// GetUser retrieves an user from the osu! API knowing their username or ID. If you need to explicitly use either "string" or "id", use GetUser{Full,ByUsername,ByID}.
func (a *APIClient) GetUser(username string, gamemode int) (User, error) {
	return a.GetUserFull(username, gamemode, "", 0)
}

// GetUserByID retrieves an user from the osu! API knowing their ID. Well that was pretty fucking obvious.
func (a *APIClient) GetUserByID(id int, gamemode int) (User, error) {
	return a.GetUserFull(itos(id), gamemode, "id", 0)
}

// GetUserByUsername retrieves an user from the osu! API knowing their **username** (not either username or ID).
func (a *APIClient) GetUserByUsername(username string, gamemode int) (User, error) {
	return a.GetUserFull(username, gamemode, "string", 0)
}

// GetUserFull allows you to grab an user from the osu! API with full control over the passed parameters.
func (a *APIClient) GetUserFull(username string, gamemode int, usernameType string, events int) (endUser User, retErr error) {
	retErr = nil
	endUser = User{}
	if gamemode > 3 || gamemode < 0 {
		retErr = fmt.Errorf("passed gamemode is invalid")
		return
	}
	if events > 31 || events < 0 {
		retErr = fmt.Errorf("event days to show is invalid (must be in range 1-31, or 0 to disable)")
		return
	}
	if usernameType != "" && usernameType != "id" && usernameType != "string" {
		retErr = fmt.Errorf(`username type is invalid (must be either "string", "id" or empty string)`)
		return
	}
	qs := map[string]string{
		"u": username,
		"m": itos(gamemode),
	}
	if events != 0 {
		qs["event_days"] = itos(events)
	}
	if usernameType != "" {
		qs["type"] = usernameType
	}
	backJSON, err := a.makeRequest("get_user", qs)
	if err != nil {
		retErr = err
		return
	}
	usersArray := []User{}
	err = json.Unmarshal(backJSON, &usersArray)
	if err != nil {
		retErr = fmt.Errorf("There was an error unmarshaling the returned JSON from the osu! API. %v", err)
		return
	}
	if len(usersArray) == 0 {
		// No user was found, so /shrugs
		return
	}
	endUser = usersArray[0]
	for index, value := range endUser.Events {
		endUser.Events[index].Date, err = time.Parse("2006-01-02 15:04:05", value.DateRaw)
		if err != nil {
			retErr = fmt.Errorf("There was an error parsing the date in the user events. %v", err)
			return
		}
	}
	// Fucking hell that was a long ride.
	return
}
