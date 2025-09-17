package cards

// Card represents a single playing card in the deck.
// Each card has a suit (Hearts, Diamonds, Clubs, Spades),
// a face value (e.g., "Ace", "2", ..., "King"), a numeric
// rank used for ordering, and an optional value field for
// game-specific scoring.
//
// Example:
//
//	// Standard card
//	card := Card{Suit: SuitHearts, Face: "Ace", Rank: 14, Value: 1}
//
//	// Joker card
//	card := Card{Suit: "", Face: FaceJoker, Rank: 0, Value: 0}
type Card struct {
	Suit  string
	Face  string
	Rank  int
	Value int
}

// String returns a human-readable string representation of the card.
// For standard suits, it includes the face and the Unicode symbol of the suit.
// For jokers or custom cards without a suit, only the face is returned.
//
// Example:
//
//	card := Card{Suit: SuitSpades, Face: "King", Rank: 13, Value: 10}
//	fmt.Println(card.String()) // Output: King of ♠
//
//	card = Card{Suit: "", Face: FaceJoker, Rank: 0, Value: 0}
//	fmt.Println(card.String()) // Output: Joker
func (c Card) String() string {
	if c.Suit == SuitHearts {
		return c.Face + " of " + " ♥ "
	} else if c.Suit == SuitDiamonds {
		return c.Face + " of " + " ♦ "
	} else if c.Suit == SuitClubs {
		return c.Face + " of " + " ♣ "
	} else if c.Suit == SuitSpades {
		return c.Face + " of " + " ♠ "
	}
	return c.Face
}
