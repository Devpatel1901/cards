package cards

import "testing"

func TestAddJokers(t *testing.T) {
	var cards []Card

	cards = append(cards, Card{Suit: suits[0], Rank: Ace})
	cards = append(cards, Card{Suit: suits[1], Rank: Two})
	cards = append(cards, Card{Suit: suits[2], Rank: Three})
	cards = append(cards, Card{Suit: suits[3], Rank: Four})

	want := 6

	cards = AddJokers(cards)

	if len(cards) != want {
		t.Errorf("Expected number of cards = %d, but got = %d", want, len(cards))
	}

	if !cards[5].Suit.Equals(RedJoker) {
		t.Errorf("Expected Red Joker Card, but got = %v", cards[5].Suit.String())
	}

	if !cards[4].Suit.Equals(BlackJoker) {
		t.Errorf("Expected Black Joker Card, but got = %v", cards[4].Suit.String())
	}
}
