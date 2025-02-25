# Deck Package

This Go package provides functionality for creating, sorting, shuffling, and manipulating a deck of playing cards. It implements a flexible card deck system with options for customizing the deck's contents and behavior using functional options.

## Features

- **Card Representation**: The `Card` type is used to represent a playing card with a suit and rank.
- **Deck Creation**: The `New` function generates a standard deck of cards.
- **Sorting**: The package supports sorting cards by suit and rank, as well as custom sorting through functional options.
- **Shuffling**: The deck can be shuffled using random permutations.
- **Jokers**: Add a custom number of Joker cards to the deck.
- **Filtering**: A flexible filtering option to remove specific cards (e.g., remove all 2s and 3s).
- **Multiple Decks**: Combine multiple decks into a single deck.

## Structure

The package contains the following components:

### `Card` Type

The `Card` type represents a single playing card and includes two fields:

- `Suit` (type: `Suit`): The suit of the card (Spade, Diamond, Club, Heart, Joker).
- `Rank` (type: `Rank`): The rank of the card (Ace, 2, 3, ..., King).

### Functional Options

Functional options allow you to customize the behavior of the deck during creation. The available options include:

- **`Sort`**: Sort the deck based on a custom comparison function.
- **`Shuffle`**: Shuffle the deck using random permutations.
- **`Jokers(n int)`**: Add `n` Joker cards to the deck.
- **`Filter(f func(Card) bool)`**: Filter out specific cards based on the provided filter function.
- **`Deck(n int)`**: Create a deck composed of `n` copies of a standard deck.

## Installation

To use the `deck` package in your Go project, you can import it into your Go code:

```go
import "https://github.com/shreyshah-06/deck/deck"
```