// A Go implementation of "Acey Ducey" (In-Between) card game.
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// Game rules match original BCG listing. See README.md for limitations.

// trailing whitespace is intentionally added to about, heading.
var bookIntro = `
=============================================================================
| THIS IS A SIMULATION OF THE ACEY DUCEY CARD GAME.                         |
| IN THE GAME, THE DEALER (THE COMPUTER) DEALS TWO CARDS FACE UP.  YOU      |
| HAVE AN OPTION TO BET OR NOT TO BET DEPENDING ON WHETHER OR NOT YOU       |
| FEEL THE NEXT CARD DEALT WILL HAVE A VALUE BETWEEN THE FIRST TWO.         |
|                                                                           |
| YOUR INITIAL MONEY (BANKROLL) IS SET TO $100.  YOU MAY ALTER THIS VAR IF  |
| YOU WANT TO START WITH MORE OR LESS THAN $100.  THE GAME KEEPS GOING      |
| UNTIL YOU LOSE ALL YOUR MONEY OR INTERRUPT THE PROGRAM.                   |
|                                                                           |
| THE ORIGINAL PROGRAM AUTHOR WAS BILL PALMBY OF PRAIRIE VIEW, ILLINOIS.    |
=============================================================================

`

var heading = `
ACEY DUCEY CARD GAME
CREATIVE COMPUTING  MORRISTOWN, NEW JERSEY


`

var gameIntro = `
ACEY-DUCEY IS PLAYED IN THE FOLLOWING MANNER
THE DEALER (COMPUTER) DEALS TWO CARDS FACE UP
YOU HAVE AN OPTION TO BET OR NOT BET DEPENDING
ON WHETHER OR NOT YOU FEEL THE CARD WILL HAVE
A VALUE BETWEEN THE FIRST TWO.

IF YOU DO NOT WANT TO BET, INPUT A 0`

// var bankroll = 100
var currency = "DOLLARS"

// getCard returns an int of values 2 to 14, both inclusive. This is a flawed
// deal in that it does not model/track the deck. But is OK for a single player.
func getCard() int {
	return rand.Intn(13) + 2
}

// getDealerCards returns a 2 element array, with each value being 2-14 inclusive,
// and where elements are sorted increasing.  This follows the rules of the 1978
// edition of BCG (NOTE: not exactly official rules).
func getDealerCards() [2]int {
	//
	var dc [2]int
	dc[0] = getCard()
	// Prevent card A from exact matching card B (BCG rule)
	for {
		dc[1] = getCard()
		if dc[1] != dc[0] {
			break
		}
	}
	// sort low-high
	if dc[0] > dc[1] {
		dc[0], dc[1] = dc[1], dc[0]
	}

	return dc
}

func getCardName(c int) string {
	switch c {
	case 11:
		return "JACK"
	case 12:
		return "QUEEN"
	case 13:
		return "KING"
	case 14:
		return "ACE"
	default:
		return strconv.Itoa(c)
	}
}

func end() {
	fmt.Printf("\nOK HOPE YOU HAD FUN\n")
	os.Exit(0)
}

func prompt() bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("SHOW BOOK TEXT? (Y or YES): ")
		// ReadString will block until the delimiter arg is entered
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("AN ERROR OCCURED WHILE READING INPUT. PLEASE TRY AGAIN", err)
		}

		// remove delimeter
		response = strings.ToLower(strings.TrimSuffix(response, "\n"))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" || response == "" {
			return false
		}
	}
}

// getBet returns an int representing a valid bet amount >= 0, after validation loop.
func getBet(bankroll int) (bet int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("\nWHAT IS YOUR BET? ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("AN ERROR OCCURED WHILE READING INPUT. PLEASE TRY AGAIN", err)
		}

		// remove delimeter
		response = strings.ToLower(strings.TrimSuffix(response, "\n"))

		// Handle "0" while string, first, since Atoi() will treat non-numeric characters as 0.
		// Because 0 is valid, the "Chicken" logic is handled outside this loop.
		if response == "0" {
			return 0
		}

		// Cast input to string, and trap input errors
		bet, err := strconv.Atoi(response)
		if err != nil {
			fmt.Println("AN ERROR OCCURED WHILE PARSING INPUT. PLEASE TRY AGAIN")
			continue
		}

		// Trap negative value bets
		if bet < 0 {
			fmt.Println("PLEASE ENTER A POSITIVE NUMBER")
			continue
		}

		// Trap too large value bets
		if bet > bankroll {
			fmt.Println("SORRY, MY FRIEND BUT YOU BET TOO MUCH")
			fmt.Printf("YOU HAVE ONLY %d %s TO BET\n", bankroll, currency)
			continue
		}

		// By this point we checked for all non-normal clauses)
		//if bet > 0 && bet <= bankroll {
		if bet > 0 {
			// The bet is a valid amount not exceeding our wad.
			return bet
		}
	}
}

func main() {
	var bankroll = 100
	fmt.Println(heading)
	if prompt() {
		fmt.Println(bookIntro)
	}
	fmt.Println(gameIntro)

	for {
		var dealerCards [2]int = getDealerCards()

		fmt.Printf("YOU NOW HAVE %d %s.\n\n", bankroll, currency)
		fmt.Printf("HERE ARE YOUR NEXT TWO CARDS:\n %s\n %s\n", getCardName(dealerCards[0]), getCardName(dealerCards[1]))

		bet := getBet(bankroll)
		//fmt.Println("*** DEBUG Your bet =", bet)

		// 0 is a valid bet, but we end the current round.
		if bet == 0 {
			fmt.Printf("CHICKEN!!\n\n")
			continue
		}

		playerCard := getCard()
		fmt.Printf(" %s\n", getCardName(playerCard))

		if playerCard > dealerCards[0] && playerCard <= dealerCards[1] {
			fmt.Println("YOU WIN!!!")
			bankroll = bankroll + bet
		} else {
			fmt.Println("SORRY, YOU LOSE")
			bankroll = bankroll - bet
		}

		if bankroll <= 0 {
			fmt.Println("SORRY, FRIEND, BUT YOU BLEW YOUR WAD")
			end()
		}
	}
}
