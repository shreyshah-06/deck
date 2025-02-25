package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

// ExampleCard demonstrates how to print a card in the format "Rank of Suit".
// It also demonstrates printing a Joker card.
func ExampleCard() {
	// Print different cards using the Card struct
	fmt.Println(Card{Rank: Ace, Suit: Heart})  
	fmt.Println(Card{Rank: Two, Suit: Club})   
	fmt.Println(Card{Rank: Three, Suit: Spade}) 
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})              

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Three of Spades
	// Jack of Diamonds
	// Joker
}

// TestNew tests the creation of a new deck of cards and ensures it contains
// the correct number of cards, including the proper suit/rank combinations.
func TestNew(t *testing.T) {
	// Generate a new deck
	cards := New()

	// There should be 13 ranks per suit, and 4 suits, resulting in 52 cards
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}

	// Ensure all suits and ranks are present in the correct order
	for i, suit := range suits {
		for j := minRank; j <= maxRank; j++ {
			// Check if the card at the correct index has the expected suit and rank
			if cards[i*13+int(j)-1].Suit != suit || cards[i*13+int(j)-1].Rank != Rank(j) {
				t.Error("Expected", suit, Rank(j), "but got", cards[i*13+int(j)-1])
			}
		}
	}
}

// TestDefaultSort tests the DefaultSort function, ensuring that the cards
// are sorted as expected.
func TestDefaultSort(t *testing.T) {
	// Generate a deck sorted by default
	cards := New(DefaultSort)

	// The first card after sorting should be the Ace of Spades
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected", exp, "but got", cards[0])
	}
}

// TestSort tests the Sort function with a custom sorting condition.
func TestSort(t *testing.T) {
	// Generate a deck sorted by the custom sorting function 'Less'
	cards := New(Sort(Less))

	// The first card should be the Ace of Spades as per the sorting function
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected", exp, "but got", cards[0])
	}
}

// TestJokers tests the addition of Joker cards to the deck.
func TestJokers(t *testing.T) {
	// Generate a deck with 3 Joker cards
	cards := New(Jokers(3))
	count := 0

	// Count how many Joker cards are in the deck
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}

	// Ensure exactly 3 Joker cards are in the deck
	if count != 3 {
		t.Error("Expected 3 Jokers but got", count)
	}
}

// TestShuffle tests the shuffling functionality to ensure that cards
// are shuffled correctly, and the shuffle is deterministic for testing.
func TestShuffle(t *testing.T) {
	// Make shuffleRand deterministic to test a predictable shuffle
	shuffleRand = rand.New(rand.NewSource(0))

	// Generate the original deck
	originalCards := New()

	// Record the cards that should be in the first and second position
	first := originalCards[40]
	second := originalCards[35]

	// Shuffle the deck
	cards := New(Shuffle)

	// Ensure the shuffled deck has the same cards in the first and second position
	if cards[0] != first {
		t.Error("Expected first card to be", first, "but got", cards[0])
	}
	if cards[1] != second {
		t.Error("Expected second card to be", second, "but got", cards[1])
	}
}

// TestFilter tests the filtering functionality to ensure that certain cards
// are excluded from the deck based on the provided filter condition.
func TestFilter(t *testing.T) {
	// Define a filter function that excludes cards with Rank Two or Three
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	// Generate a deck with the filter applied
	cards := New(Filter(filter))

	// Check that no cards with Rank Two or Three remain in the deck
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all 2s and 3s to be filtered out")
		}
	}
}

// TestDeck tests the functionality to create multiple decks in one.
func TestDeck(t *testing.T) {
	// Generate 3 decks by repeating the base deck 3 times
	cards := New(Deck(3))

	// Ensure the total number of cards is 3 times the size of a single deck
	if len(cards) != 13*4*3 {
		t.Error("Expected 3 decks but got", len(cards)/13)
	}
}
