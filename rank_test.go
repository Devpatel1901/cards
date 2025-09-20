package cards

import "testing"

func TestRankToString(t *testing.T) {
	testCases := []struct {
		name     string
		rank     Rank
		expected string
	}{
		{name: "Rank Of Ace", rank: Ace, expected: "A"},
		{name: "Rank of Two", rank: Two, expected: "2"},
		{name: "Rank of Three", rank: Three, expected: "3"},
		{name: "Rank of Four", rank: Four, expected: "4"},
		{name: "Rank of Five", rank: Five, expected: "5"},
		{name: "Rank of Six", rank: Six, expected: "6"},
		{name: "Rank of Seven", rank: Seven, expected: "7"},
		{name: "Rank of Eight", rank: Eight, expected: "8"},
		{name: "Rank of Nine", rank: Nine, expected: "9"},
		{name: "Rank of Ten", rank: Ten, expected: "10"},
		{name: "Rank of Jack", rank: Jack, expected: "J"},
		{name: "Rank of Queen", rank: Queen, expected: "Q"},
		{name: "Rank of King", rank: King, expected: "K"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.rank.Single()
			want := tc.expected

			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", want, actual)
			}
		})
	}
}
