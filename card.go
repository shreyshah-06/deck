//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents the four possible suits of a card deck.
type Suit uint8

// Constants for the four standard suits of a deck of cards.
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // Joker is a special case, not part of the standard deck
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

// Rank represents the rank of a card (Ace, 2, 3, ..., King).
type Rank uint8

// Constants representing card ranks from Ace to King.
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

// Constants for the minimum and maximum ranks in a deck.
const (
	minRank = Ace
	maxRank = King
)

// Card represents a playing card with a suit and rank.
type Card struct {
	Suit
	Rank
}

// String returns a string representation of the card. For example: "Ace of Spades"
func (c Card) String() string {
	// Special case for Joker card
	if c.Suit == Joker {
		return c.Suit.String()
	}
	// Format the card's rank and suit
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

// New creates a new deck of cards, optionally applying any provided functions to modify the deck.
func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	// Generate standard cards (4 suits x 13 ranks)
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	// Apply any additional modifications (e.g., Jokers, custom filters)
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

// DefaultSort returns a sorted deck of cards based on their rank and suit.
func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

// Sort allows for custom sorting of cards by passing a 'less' function.
func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

// Less returns a comparison function for sorting cards based on their absolute rank.
func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRanks(cards[i]) < absRanks(cards[j])
	}
}

// absRanks calculates a unique integer value for each card based on suit and rank.
func absRanks(c Card) int {
	// The absolute value considers the suit and rank in a unique manner
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

// shuffleRand is a global random source used for shuffling the deck.
var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Shuffle randomly shuffles a deck of cards using the Fisher-Yates algorithm.
func Shuffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	perm := shuffleRand.Perm(len(cards))
	for i, j := range perm {
		ret[i] = cards[j]
	}
	return ret
}

// Jokers returns a function that adds a specified number of Joker cards to the deck.
func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Rank: Rank(i),
				Suit: Joker,
			})
		}
		return cards
	}
}

// Filter returns a function that filters out cards from the deck based on a given condition.
func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			// If the card doesn't satisfy the condition, it gets added to the result
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

// Deck returns a function that duplicates the deck 'n' times.
func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
