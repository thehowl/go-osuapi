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
}

func (a ApprovedStatus) String() string {
	if int(a) < len(approvedStatusesString) {
		return approvedStatusesString[a]
	}
	return strconv.Itoa(int(a))
}
