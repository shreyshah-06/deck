package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
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

func TestNew(t *testing.T){

	cards := New()

	// 13 ranks * 4 suits
	if len(cards) != 13*4{
		t.Error("Wrong number of cards in a new deck")
	}

	// check if all suits and ranks are present and the order of suits is (Spade, Diamond, Club, Heart)

	for i, suit := range suits{
		for j := minRank; j <= maxRank; j++{
			if cards[i*13 + int(j) - 1].Suit != suit || cards[i*13 + int(j) - 1].Rank != Rank(j){
				t.Error("Expected", suit, Rank(j), "but got", cards[i*13 + int(j) - 1])
			}
		}
	}

}

func TestDefaultSort(t *testing.T){
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp{
		t.Error("Expected", exp, "but got", cards[0])
	}
}
func TestSort(t *testing.T){
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp{
		t.Error("Expected", exp, "but got", cards[0])
	}
}

func TestJokers(t *testing.T){
	cards := New(Jokers(3))
	count := 0
	for _, card := range cards{
		if card.Suit == Joker{
			count++
		}
	}
	if count != 3{
		t.Error("Expected 3 Jokers but got", count)
	}
}

func TestFilter(t *testing.T){
	filter := func(card Card) bool{
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, card := range cards{
		if card.Rank == Two || card.Rank == Three{
			t.Error("Expected all 2s and 3s to be filtered out")
		}
	}
}