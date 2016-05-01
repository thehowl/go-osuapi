package osuapi

import "strconv"

// Genres
const (
	GenreAny Genre = iota
	GenreUnspecified
	GenreVideoGame
	GenreAnime
	GenreRock
	GenrePop
	GenreOther
	GenreNovelty
	GenreHipHop Genre = iota + 1 // there's no 8, so we must manually increment it by one
	GenreElectronic
)

// Genre is the genre of a beatmap's song.
type Genre int

var genreString = [...]string{
	"any",
	"unspecified",
	"video game",
	"anime",
	"rock",
	"pop",
	"other",
	"novelty",
	"8",
	"hip hop",
	"electronic",
}

func (g Genre) String() string {
	if g >= 0 && int(g) < len(genreString) {
		return genreString[g]
	}
	return strconv.Itoa(int(g))
}
