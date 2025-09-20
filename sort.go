package cards

import "sort"

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].absRank() < cards[j].absRank()
	})
	return cards
}
