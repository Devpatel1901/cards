package cards

import "github.com/logrusorgru/aurora"

type Suit struct {
	value         int
	name          string
	color         func(interface{}) aurora.Value
	hasRank       bool
	asciiTemplate []string
}

func (s Suit) String() string {
	return s.name
}

func (s Suit) Value() int {
	return s.value
}

func (s Suit) Color(arg interface{}) aurora.Value {
	return s.color(arg)
}

func (s Suit) ASCIITemplate() []string {
	return s.asciiTemplate
}

func (s Suit) Equals(to Suit) bool {
	return s.value == to.value && s.name == to.name
}

func (s Suit) HasRank() bool {
	return s.hasRank
}

var (
	Spade = Suit{
		value:   1,
		name:    "Spade",
		color:   aurora.White,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s .    │`,
			`│  / \   │`,
			`│ (_,_)  │`,
			`│   I  %s│`,
			`└────────┘`,
		},
	}

	Diamond = Suit{
		value:   2,
		name:    "Diamond",
		color:   aurora.BrightRed,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s /\   │`,
			`│  /  \  │`,
			`│  \  /  │`,
			`│   \/ %s│`,
			`└────────┘`,
		},
	}

	Club = Suit{
		value:   3,
		name:    "Club",
		color:   aurora.White,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s _    │`,
			`│  ( )   │`,
			`│ (_x_)  │`,
			`│   Y  %s│`,
			`└────────┘`,
		},
	}

	Hearth = Suit{
		value:   4,
		name:    "Heart",
		color:   aurora.BrightRed,
		hasRank: true,
		asciiTemplate: []string{
			`┌────────┐`,
			`│%s_  _  │`,
			`│ ( \/ ) │`,
			`│  \  /  │`,
			`│   \/ %s│`,
			`└────────┘`,
		},
	}

	BlackJoker = Suit{
		value:   5,
		name:    "BlackJoker",
		color:   aurora.White,
		hasRank: false,
		asciiTemplate: []string{
			`┌────────┐`,
			`│* \||/ K│`,
			`│J /~~\ O│`,
			`│O( o o)J│`,
			`│K \ v/ *│`,
			`└────────┘`,
		},
	}
	RedJoker = Suit{
		value:   6,
		name:    "RedJoker",
		color:   aurora.BrightRed,
		hasRank: false,
		asciiTemplate: []string{
			`┌────────┐`,
			`│+ \||/ K│`,
			`│J /~~\ O│`,
			`│O( o o)J│`,
			`│K \ v/ +│`,
			`└────────┘`,
		},
	}

	suits        = [...]Suit{Spade, Diamond, Club, Hearth}
	cardTemplate = []string{
		`┌────────┐`,
		`│████████│`,
		`│████████│`,
		`│████████│`,
		`│████████│`,
		`└────────┘`,
	}
)
