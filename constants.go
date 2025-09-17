package cards

// Suit constants represent the four standard suits
// in a deck of playing cards.
//
// Example:
//
//	card := Card{Suit: SuitHearts, Face: FaceAce}
//	fmt.Println(card) // Output: A of ♥
const (
	SuitHearts   = "Hearts"
	SuitDiamonds = "Diamonds"
	SuitClubs    = "Clubs"
	SuitSpades   = "Spades"
)

// Face constants represent the possible face values
// of a standard playing card, including numbers,
// face cards, and the Joker.
//
// Example:
//
//	// A standard face card
//	card := Card{Suit: SuitSpades, Face: FaceKing}
//
//	// A joker
//	card := Card{Suit: "", Face: FaceJoker}
const (
	FaceAce   = "A"
	FaceTwo   = "2"
	FaceThree = "3"
	FaceFour  = "4"
	FaceFive  = "5"
	FaceSix   = "6"
	FaceSeven = "7"
	FaceEight = "8"
	FaceNine  = "9"
	FaceTen   = "10"
	FaceJack  = "J"
	FaceQueen = "Q"
	FaceKing  = "K"
	FaceJoker = "Joker"
)

// Rank constants represent the numeric ranking of
// face values, typically used for sorting or comparisons.
// Higher numbers correspond to stronger ranks.
//
// Note: Ace is ranked highest by default, and Joker
// has its own rank value.
//
// Example:
//
//	if faceToRank["K"] > faceToRank["Q"] {
//		fmt.Println("King outranks Queen")
//	}
const (
	RankAce = iota + 14
	RankTwo = iota + 1
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankJoker = iota + 2
)

var faceToRank = map[string]int{
	FaceAce:   RankAce,
	FaceTwo:   RankTwo,
	FaceThree: RankThree,
	FaceFour:  RankFour,
	FaceFive:  RankFive,
	FaceSix:   RankSix,
	FaceSeven: RankSeven,
	FaceEight: RankEight,
	FaceNine:  RankNine,
	FaceTen:   RankTen,
	FaceJack:  RankJack,
	FaceQueen: RankQueen,
	FaceKing:  RankKing,
	FaceJoker: RankJoker,
}

var defaultSuitSeq = []string{SuitHearts, SuitDiamonds, SuitClubs, SuitSpades}
var defaultFaceSeq = []string{FaceAce, FaceTwo, FaceThree, FaceFour, FaceFive, FaceSix, FaceSeven, FaceEight, FaceNine, FaceTen, FaceJack, FaceQueen, FaceKing}
