package cards

import (
	"math/rand/v2"
	"slices"
	"sort"
)

// Deck represents configuration options for generating a new deck of cards.
// It is not used directly; instead, it is configured through option functions
// (such as Shuffle, Joker, AdditionalDecks, etc.) and passed into New().
//
// Fields inside Deck are set internally by these option functions and should
// not be modified directly.
//
// Example:
//
//	// Create a basic deck
//	deck := New()
//
//	// Create a shuffled deck with 2 jokers
//	deck := New(Shuffle(true), Joker(2))
//
//	// Create a deck with no Hearts and no Aces
//	deck := New(FilterCard([]string{"Ace"}, []string{"Hearts"}))
//
//	// Create 2 combined decks, sorted in default order
//	deck := New(AdditionalDecks(1), DefaultSort(true))
//
// Typical usage:
//
//	options := []func(*Deck){
//		Shuffle(true),
//		Joker(3),
//		FilterCard([]string{"2"}, nil), // remove all 2s
//	}
//	deck := New(options...)
type Deck struct {
	applyDefaultSort        bool
	applyCustomSort         bool
	isShuffle               bool
	numberOfJokers          int
	numberOfAdditionalDecks int
	excludeFaceValue        []string
	excludeSuitValue        []string
}

// DefaultSort sets whether the deck should use the default comparison function for sorting.
// When enabled, cards are sorted based on their rank in ascending order.
//
// Example:
//
//	// Enable default sorting
//	deck := New(DefaultSort(true))
//
//	// Disable default sorting (cards remain unsorted unless another option applies)
//	deck := New(DefaultSort(false))
func DefaultSort(val bool) func(*Deck) {
	return func(d *Deck) {
		d.applyDefaultSort = val
	}
}

// CustomSort sets whether the deck should use a custom sorting function.
// This option applies a user-defined sort order instead of the default.
//
// Example:
//
//	// Apply custom sorting
//	deck := New(CustomSort(true))
//
//	// Skip custom sorting (deck order determined by other options such as shuffle)
//	deck := New(CustomSort(false))
func CustomSort(val bool) func(*Deck) {
	return func(d *Deck) {
		d.applyCustomSort = val
	}
}

// Shuffle sets whether the deck should be shuffled randomly.
// When enabled, the order of the cards is randomized.
//
// Example:
//
//	// Shuffle the deck
//	deck := New(Shuffle(true))
//
//	// Keep cards in order (based on default/custom sort)
//	deck := New(Shuffle(false))
func Shuffle(val bool) func(*Deck) {
	return func(d *Deck) {
		d.isShuffle = val
	}
}

// Joker adds the given number of jokers into the deck.
// Jokers are represented by the face value "Joker".
//
// Example:
//
//	// Add 2 jokers
//	deck := New(Joker(2))
//
//	// Add no jokers
//	deck := New(Joker(0))
//
//	// Add 5 jokers for custom game rules
//	deck := New(Joker(5))
func Joker(num int) func(*Deck) {
	return func(d *Deck) {
		d.numberOfJokers = num
	}
}

// AdditionalDecks configures the number of extra standard decks to include.
// Each additional deck duplicates the standard set of cards.
//
// Example:
//
//	// Use a single standard deck (no extra decks)
//	deck := New(AdditionalDecks(0))
//
//	// Use two decks combined (1 original + 1 additional)
//	deck := New(AdditionalDecks(1))
//
//	// Use three decks combined (1 original + 2 additional)
//	deck := New(AdditionalDecks(2))
func AdditionalDecks(num int) func(*Deck) {
	return func(d *Deck) {
		d.numberOfAdditionalDecks = num
	}
}

// FilterCard removes specific cards by face value, suit, or both.
// If only face values are provided, all cards with those faces are excluded.
// If only suits are provided, all cards with those suits are excluded.
// If both are provided, only matching face+suit pairs are excluded.
//
// Example:
//
//	// Remove all "2" cards (across every suit)
//	deck := New(FilterCard([]string{"2"}, nil))
//
//	// Remove all "Hearts"
//	deck := New(FilterCard(nil, []string{"Hearts"}))
//
//	// Remove only the "Ace of Spades"
//	deck := New(FilterCard([]string{"Ace"}, []string{"Spades"}))
//
//	// Remove no cards (pass empty slices)
//	deck := New(FilterCard(nil, nil))
func FilterCard(face []string, suit []string) func(*Deck) {
	return func(d *Deck) {
		d.excludeFaceValue = face
		d.excludeSuitValue = suit
	}
}

func initializeDecks(sSeq []string, fSeq []string, faceToRank map[string]int) []Card {
	var cards []Card

	for _, sVal := range sSeq {
		for _, fVal := range fSeq {
			newCard := Card{sVal, fVal, faceToRank[fVal], faceToRank[fVal]}
			cards = append(cards, newCard)

		}
	}

	return cards
}

func excludeCardFromDeck(cards []Card, face []string, suit []string) []Card {
	var filteredCards []Card

	if len(face) == 0 && len(suit) != 0 {
		for _, c := range cards {
			if !slices.Contains(suit, c.Suit) {
				filteredCards = append(filteredCards, c)
			}
		}
	} else if len(face) != 0 && len(suit) == 0 {
		for _, c := range cards {
			if !slices.Contains(face, c.Face) {
				filteredCards = append(filteredCards, c)
			}
		}
	} else {
		for _, c := range cards {
			if !(c.Face == FaceJoker || (slices.Contains(face, c.Face) && slices.Contains(suit, c.Suit))) {
				filteredCards = append(filteredCards, c)
			}
		}
	}

	return filteredCards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	}
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func New(options ...func(*Deck)) []Card {
	deck := &Deck{}
	var cards []Card

	cards = initializeDecks(defaultSuitSeq, defaultFaceSeq, faceToRank)

	for _, o := range options {
		o(deck)
	}

	if deck.numberOfAdditionalDecks > 0 {
		for range deck.numberOfAdditionalDecks {
			inti := cards
			cards = append(cards, inti...)
		}
	}

	if deck.numberOfJokers > 0 {
		for range deck.numberOfJokers {
			jokerCard := Card{"", FaceJoker, faceToRank[FaceJoker], faceToRank[FaceJoker]}
			cards = append(cards, jokerCard)
		}
	}

	if deck.applyDefaultSort {
		sort.Slice(cards, Less(cards))
	}

	if deck.applyCustomSort {
		cards = Sort(Less)(cards)
	}

	if deck.isShuffle {
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	}

	if len(deck.excludeFaceValue) != 0 || len(deck.excludeSuitValue) != 0 {
		cards = excludeCardFromDeck(cards, deck.excludeFaceValue, deck.excludeSuitValue)
	}

	return cards
}
