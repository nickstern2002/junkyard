// blackjack.go
package server

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// Define the Card struct
type Card struct {
	Suit  string
	Value string
	Score int
}

var (
	playerHand               []Card
	dealerHand               []Card
	shuffledDeck             []Card
	playerScore, dealerScore int
	suits                    = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values                   = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	deck                     []Card
	endGameState             = false
)

// Seed the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
	deck = createDeck()
}

// createDeck function to generate the deck of cards
func createDeck() []Card {
	var d []Card
	for _, suit := range suits {
		for _, value := range values {
			card := Card{
				Suit:  suit,
				Value: value,
				Score: calculateCardScore(value),
			}
			d = append(d, card)
		}
	}
	return d
}

// calculateCardScore assigns numeric values to each card for the game
func calculateCardScore(value string) int {
	switch value {
	case "J", "Q", "K", "10":
		return 10
	case "A":
		return 11
	default:
		return int(value[0] - '0') // Convert the string to a numeric value
	}
}

// shuffleDeck function to shuffle the cards in a random order
func shuffleDeck(deck []Card) []Card {
	shuffledDeck := append([]Card(nil), deck...) // create a copy of the deck
	rand.Shuffle(len(shuffledDeck), func(i, j int) {
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	})
	return shuffledDeck
}

// calculateHandScore function to sum up the total score for a hand of cards
func calculateHandScore(hand []Card) int {
	total := 0
	aces := 0
	for _, card := range hand {
		total += card.Score
		if card.Value == "A" {
			aces++
		}
	}
	for aces > 0 && total > 21 {
		total -= 10
		aces--
	}
	return total
}

// getBlackJack handles both the displaying of the game state and receiving player input
func getBlackJack(w http.ResponseWriter, r *http.Request) {
	endGameState = false
	// If it's a POST request (form submission), process the action
	if r.Method == http.MethodPost {
		// Check the input action (e.g., "hit" or "stand")
		action := r.FormValue("action")

		// If the player chooses to "hit", give them another card
		if action == "hit" {
			playerHand = append(playerHand, shuffledDeck[len(playerHand)+len(dealerHand)])
			playerScore = calculateHandScore(playerHand)
			displayGameState(w)
		}

		// If the player chooses to "stand", proceed with dealer's action
		if action == "stand" {
			for dealerScore < 17 {
				dealerHand = append(dealerHand, shuffledDeck[len(playerHand)+len(dealerHand)])
				dealerScore = calculateHandScore(dealerHand)
			}

			// Display results and winner
			endGameState = true
			displayGameState(w)
			return
		}

		if action == "reset" {
			resetGameState(w)
		}

	} else {
		// Initialize the game for the first time
		initializeGame()
		// Display the initial state and offer a "hit" or "stand" option
		displayGameState(w)
	}
}

// initializeGame sets up a new game (shuffle, deal cards)
func initializeGame() {
	shuffledDeck = shuffleDeck(deck)
	playerHand = []Card{shuffledDeck[0], shuffledDeck[1]}
	dealerHand = []Card{shuffledDeck[2], shuffledDeck[3]}
	playerScore = calculateHandScore(playerHand)
	dealerScore = calculateHandScore(dealerHand)
}

// logGameState logs the current state of the game
func logGameState() {
	log.Print(
		"\n\nSTATE BREAK\n\n",
		"endGameState: ",
		endGameState,
		"\nLogging Player Stats...\n",
		playerHand,
		"\nPlayer Score: ",
		playerScore,
		"\n\nLogging Dealer Stats...\n",
		dealerHand,
		"\nDealer Score: ",
		dealerScore,
	)
}

// displayGameState shows the current game status
func displayGameState(w http.ResponseWriter) {
	logGameState()
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")

	// Display the player's hand and score
	playerHandOutput := ""
	for _, card := range playerHand {
		playerHandOutput += card.Value + " " + card.Suit + ", "
	}
	io.WriteString(w, fmt.Sprintf("<p>Your hand: %s </p>", playerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Your score: %d</p>", playerScore))

	// Display the dealer's hand (only one card visible initially)
	io.WriteString(w, fmt.Sprintf("<p>Dealer shows: %s %s</p>", dealerHand[0].Value, dealerHand[0].Suit))

	// Display the "hit" and "stand" options in a form
	if playerScore <= 21 {
		if dealerScore >= 17 && endGameState {

			// Get Dealers Hand and Output it
			dealerHandOutput := ""
			for _, card := range dealerHand {
				dealerHandOutput += card.Value + " " + card.Suit + ", "
			}

			io.WriteString(w, fmt.Sprintf("<p>Dealer's final hand: %s </p>", dealerHandOutput))
			io.WriteString(w, fmt.Sprintf("<p>Dealer's score: %d</p>", dealerScore))

			if dealerScore > 21 {
				io.WriteString(w, "<h1>Dealer Busted: You Win!</h1>")
			} else if dealerScore > playerScore {
				io.WriteString(w, "<h1>You Lose :(</h1>")
			} else if dealerScore < playerScore {
				io.WriteString(w, "<h1>You Win!</h1>")
			} else {
				io.WriteString(w, "<h1>You Draw :| </h1>")
			}
			io.WriteString(w, `
			<form method="post">
				<button name="action" value="reset">Reset Game</button>
			</form>
			`)
		} else {
			io.WriteString(w, `
			<form method="post">
				<button name="action" value="hit">Hit</button>
				<button name="action" value="stand">Stand</button>
			</form>
		`)
		}
	} else {
		io.WriteString(w, "<p>Busted! You lose!</p>")
		io.WriteString(w, `
			<form method="post">
				<button name="action" value="reset">Reset Game</button>
			</form>
		`)
	}
}

func resetGameState(w http.ResponseWriter) {
	endGameState = false
	initializeGame()
	displayGameState(w)
}
