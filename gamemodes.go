package osuapi

// The following are the game modes of osu!, as well as a few shorthands (like Std, Standard, Ctb...)
const (
	OsuStandard = iota
	Taiko
	CatchTheBeat
	OsuMania
	Standard = OsuStandard
	Std      = OsuStandard
	Ctb      = CatchTheBeat
	Mania    = OsuMania
)
