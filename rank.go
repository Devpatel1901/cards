//go:generate stringer -type=Rank
package cards

import "strconv"

type Rank int

const (
	_ Rank = iota

	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

func (r Rank) Single() string {
	if r > Ace && r < Jack {
		return strconv.Itoa(int(r))
	}

	if r == Ace {
		return "A"
	}

	return string(r.String()[0])
}
