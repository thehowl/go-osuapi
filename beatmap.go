package osuapi

// These are the statuses an osu! beatmap can have.
const (
	Graveyard = -2
	WIP       = -1
	Pending   = 0
	Ranked    = 1
	Approved  = 2
	Qualified = 3
)

// Any is shared between languages and genre, so that's why it's outside both blocks.
const Any = 0

// These are the Language IDs a beatmap can have.
const (
	LanguageOther = 1
	English       = 2
	Japanese      = 3
	Chinese       = 4
	Instrumental  = 5
	Korean        = 6
	French        = 7
	German        = 8
	Swedish       = 9
	Spanish       = 10
	Italian       = 11
)

// There are the various genres a beatmap can have as a GenreID.
const (
	Unspecified = 1
	Videogame   = 2
	Anime       = 3
	Rock        = 4
	Pop         = 5
	GenreOther  = 6
	Novelty     = 7
	HipHop      = 9
	Electronic  = 10
)

// Beatmap contains information about a beatmap difficulty.
type Beatmap struct {
	Approved          int       `json:"approved,string"`
	ApprovedDate      MySQLDate `json:"approved_date"`
	Artist            string    `json:"artist"`
	BeatmapID         int       `json:"beatmap_id,string"`
	BeatmapsetID      int       `json:"beatmapset_id,string"`
	Bpm               float64   `json:"bpm,string"`
	Creator           string    `json:"creator"`
	ApproachRate      int       `json:"diff_approach,string"`
	HPDrain           int       `json:"diff_drain,string"`
	OverallDifficulty int       `json:"diff_overall,string"`
	CircleSize        int       `json:"diff_size,string"`
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

// GetBeatmapFull allows you to lookup for a beatmap using all the search options provided by the API.
//
// If you want to use this, please make sure there's no other way to do what you are doing, as this function is
// extremely unreadable when written as is.
func GetBeatmapFull(
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
	// Limit of results to give in a page. Ignore value: 0.
	limit int,
) (b Beatmap, retErr error) {
	b = Beatmap{}
	retErr = nil
	return
}
