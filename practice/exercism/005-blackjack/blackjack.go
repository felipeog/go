package blackjack

const (
	hit   = "H"
	split = "P"
	stand = "S"
	win   = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardsSum := ParseCard(card1) + ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)

	switch {
	// If you have a pair of aces, you split.
	case cardsSum == 22:
		return split

	// If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten, you automatically win.
	case cardsSum == 21 && dealerCardValue < 10:
		return win

	// If you have a Blackjack (two cards that sum up to a value of 21), and the dealer has an ace, a figure or a ten, you stand.
	case cardsSum == 21 && dealerCardValue >= 10:
		return stand

	// If your cards sum up to a value within the range [17, 20], you stand.
	case cardsSum >= 17 && cardsSum <= 20:
		return stand

	// If your cards sum up to a value within the range [12, 16] and the dealer has a 6 or lower, you stand.
	case cardsSum >= 12 && cardsSum <= 16 && dealerCardValue < 7:
		return stand

	// If your cards sum up to a value within the range [12, 16] you the dealer has a 7 or higher, you hit.
	case cardsSum >= 12 && cardsSum <= 16 && dealerCardValue >= 7:
		return hit

	// If your cards sum up to 11 or lower you hit.
	default:
		return hit
	}
}
