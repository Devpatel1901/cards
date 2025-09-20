package cards

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := NewDeck()
	if len(cards) != 13*4 {
		t.Errorf("Expected 52 cards, but got %d", len(cards))
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := NewDeck(FilterOut(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Errorf("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestFromDecks(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	cards := FromDecks(deck1, deck2)

	if len(cards) != 52*2 {
		t.Errorf("not enough cards, expected 104 cards, but got %d", len(cards))
	}
}

func TestDefaultSort(t *testing.T) {
	cards := []Card{
		{Suit: Hearth, Rank: Two},
		{Suit: Spade, Rank: Queen},
		{Suit: Spade, Rank: Four},
		{Suit: Diamond, Rank: Ace},
		{Suit: Club, Rank: Ten},
	}
	given := make([]Card, len(cards))
	copy(given, cards)

	cards = DefaultSort(cards)

	expectedRanges := [...]int{2, 1, 3, 4, 0}
	for gotRange, expectedRange := range expectedRanges {
		if !cards[gotRange].Equals(given[expectedRange]) {
			t.Errorf("card number %d must be %s, got %s", gotRange, given[expectedRange].String(), cards[gotRange].String())
		}
	}
}

func TestShuffle(t *testing.T) {
	given := []Card{
		{Suit: Hearth, Rank: Two},
		{Suit: Spade, Rank: Queen},
		{Suit: Spade, Rank: Four},
		{Suit: Diamond, Rank: Ace},
		{Suit: Club, Rank: Ten},
	}
	cards := make([]Card, len(given))
	copy(cards, given)

	cards = Shuffle(cards)

	if len(cards) != len(given) {
		t.Errorf("the result must preserve the slice length, expected %d, got %d", len(given), len(cards))
	}

	isSame := true
	for i := 0; i < len(cards); i++ {
		if !cards[i].Equals(given[i]) {
			isSame = false
			break
		}
	}
	if isSame {
		t.Error("shuffle has failed miserably...")
	}
}

func TestComputeCoeff(t *testing.T) {
	var tests = map[string]struct {
		given    int
		expected int
	}{
		"13":    {13, 100},
		"99":    {99, 100},
		"1":     {1, 10},
		"81234": {81234, 100000},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := computeCoeff(tt.given)
			if actual != tt.expected {
				t.Errorf("(%d): expected %d, actual %d", tt.given, tt.expected, actual)
			}
		})
	}
}
