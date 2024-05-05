package main

import "testing"

func TestGetCard(t *testing.T) {
	t.Run("validate card is not less than 2, inclusive.", func(t *testing.T) {
		got := getCard()
		want := 2

		if got < want {
			t.Errorf("got %q, want must at least %q", got, want)
		}
	})
	t.Run("validate card is not more than 14, inclusive.", func(t *testing.T) {
		got := getCard()
		want := 14

		if got > want {
			t.Errorf("got %q, want must at most be %q", got, want)
		}
	})
}

func TestGetDealerCards(t *testing.T) {
	var dealerCards [2]int = getDealerCards()

	if dealerCards[0] > dealerCards[1] {
		t.Errorf("got %q and %q. Ordering is not LOW-HIGH", dealerCards[0], dealerCards[1])
	}
}
