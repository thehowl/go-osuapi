package osuapi

import "strconv"

// Approved statuses.
const (
	StatusGraveyard ApprovedStatus = iota - 2
	StatusWIP
	StatusPending
	StatusRanked
	StatusApproved
	StatusQualified
	StatusLoved
)

// ApprovedStatus - also known as ranked status - is the status of a beatmap.
// Yeah, no shit, I know. It tells whether the beatmap is ranked, qualified,
// graveyarded or other memes.
type ApprovedStatus int

var approvedStatusesString = [...]string{
	"graveyard",
	"WIP",
	"pending",
	"ranked",
	"approved",
	"qualified",
	"loved",
}

func (a ApprovedStatus) String() string {
	if a >= -2 && int(a)+2 < len(approvedStatusesString) {
		return approvedStatusesString[a+2]
	}
	return strconv.Itoa(int(a))
}
