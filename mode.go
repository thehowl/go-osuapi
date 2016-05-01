package osuapi

import "strconv"

// osu! game modes IDs.
const (
	ModeOsu Mode = iota
	ModeTaiko
	ModeCatchTheBeat
	ModeOsuMania
)

// Mode is an osu! game mode.
type Mode int

var modesString = [...]string{
	"osu!",
	"Taiko",
	"Catch the Beat",
	"osu!mania",
}

func (m Mode) String() string {
	if m >= 0 && m <= 3 {
		return modesString[m]
	}
	return strconv.Itoa(int(m))
}
