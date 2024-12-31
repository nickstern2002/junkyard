package blackjack

import (
	"fmt"
	"io"
	"net/http"
)

func DisplayWelcomeScreen(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, `<form method="post"><button name="action" value="start">Start Game</button></form>`)
}

func DisplayPlaceBetScreen(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, `
    	<form method="post">
        	<label for="chips">Enter number of chips:</label>
        	<input type="number" id="chips" name="bet">
        	<button type="submit" name="action" value="placeBet">Submit</button>
    	</form>
	`)

}

func DisplayInvalidInputScreen(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, "<p>Invalid input</p>")
	io.WriteString(w, `
    	<form method="post">
        	<label for="chips">Enter number of chips:</label>
        	<input type="number" id="chips" name="bet">
        	<button type="submit" name="action" value="placeBet">Submit</button>
    	</form>
	`)
}

func DisplayHitStandState(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, fmt.Sprintf("<p>Your bet amount: $%d</p>", game.activeBet))

	// Display the player's hand and score
	playerHandOutput := ""
	for _, card := range game.playerHand {
		playerHandOutput += card.Value + " " + card.Suit + ", "
	}
	io.WriteString(w, fmt.Sprintf("<p>Your hand: %s </p>", playerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Your score: %d</p>", game.playerScore))

	// Display the dealer's hand (only one card visible initially)
	io.WriteString(w, fmt.Sprintf("<p>Dealer shows: %s %s</p>", game.dealerHand[0].Value, game.dealerHand[0].Suit))

	io.WriteString(w, `
			<form method="post">
				<button name="action" value="hit">Hit</button>
				<button name="action" value="stand">Stand</button>
			</form>
		`)
}

func DisplayPlayerBust(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, fmt.Sprintf("<p>Your bet amount: $%d</p>", game.activeBet))

	// Display the player's hand and score
	playerHandOutput := ""
	for _, card := range game.playerHand {
		playerHandOutput += card.Value + " " + card.Suit + ", "
	}
	io.WriteString(w, fmt.Sprintf("<p>Your hand: %s </p>", playerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Your score: %d</p>", game.playerScore))

	// Display the dealer's hand (only one card visible initially)
	io.WriteString(w, fmt.Sprintf("<p>Dealer shows: %s %s</p>", game.dealerHand[0].Value, game.dealerHand[0].Suit))

	io.WriteString(w, "<h1>Busted! You lose!</h1>")
	DisplayNextRoundOrGameOver(w, game)
}

func DisplayDealerBust(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, fmt.Sprintf("<p>Your bet amount: $%d</p>", game.activeBet))

	// Display the player's hand and score
	playerHandOutput := ""
	for _, card := range game.playerHand {
		playerHandOutput += card.Value + " " + card.Suit + ", "
	}
	io.WriteString(w, fmt.Sprintf("<p>Your hand: %s </p>", playerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Your score: %d</p>", game.playerScore))

	// Display the Dealer's full hand and score
	dealerHandOutput := ""
	for _, card := range game.dealerHand {
		dealerHandOutput += card.Value + " " + card.Suit + ", "
	}

	io.WriteString(w, fmt.Sprintf("<p>Dealer's hand: %s </p>", dealerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Dealer's score: %d</p>", game.dealerScore))
	io.WriteString(w, "<h1>Dealer Busted! You win!</h1>")
	DisplayNextRoundOrGameOver(w, game)
}

func DisplayPlayerLose(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, fmt.Sprintf("<p>Your bet amount: $%d</p>", game.activeBet))

	// Display the player's hand and score
	playerHandOutput := ""
	for _, card := range game.playerHand {
		playerHandOutput += card.Value + " " + card.Suit + ", "
	}
	io.WriteString(w, fmt.Sprintf("<p>Your hand: %s </p>", playerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Your score: %d</p>", game.playerScore))

	// Display the Dealer's full hand and score
	dealerHandOutput := ""
	for _, card := range game.dealerHand {
		dealerHandOutput += card.Value + " " + card.Suit + ", "
	}

	io.WriteString(w, fmt.Sprintf("<p>Dealer's hand: %s </p>", dealerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Dealer's score: %d</p>", game.dealerScore))
	io.WriteString(w, "<h1>You Lost!</h1>")
	DisplayNextRoundOrGameOver(w, game)
}

func DisplayPlayerWin(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, fmt.Sprintf("<p>Your bet amount: $%d</p>", game.activeBet))

	// Display the player's hand and score
	playerHandOutput := ""
	for _, card := range game.playerHand {
		playerHandOutput += card.Value + " " + card.Suit + ", "
	}
	io.WriteString(w, fmt.Sprintf("<p>Your hand: %s </p>", playerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Your score: %d</p>", game.playerScore))

	// Display the Dealer's full hand and score
	dealerHandOutput := ""
	for _, card := range game.dealerHand {
		dealerHandOutput += card.Value + " " + card.Suit + ", "
	}

	io.WriteString(w, fmt.Sprintf("<p>Dealer's hand: %s </p>", dealerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Dealer's score: %d</p>", game.dealerScore))
	io.WriteString(w, "<h1>You Won!</h1>")
	DisplayNextRoundOrGameOver(w, game)
}

func DisplayPlayerTie(w http.ResponseWriter, game *Game) {
	io.WriteString(w, "<h1>Welcome to Blackjack</h1>")
	io.WriteString(w, fmt.Sprintf("<p>Your balance: $%d</p>", game.chipAmount))
	io.WriteString(w, fmt.Sprintf("<p>Your bet amount: $%d</p>", game.activeBet))

	// Display the player's hand and score
	playerHandOutput := ""
	for _, card := range game.playerHand {
		playerHandOutput += card.Value + " " + card.Suit + ", "
	}
	io.WriteString(w, fmt.Sprintf("<p>Your hand: %s </p>", playerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Your score: %d</p>", game.playerScore))

	// Display the Dealer's full hand and score
	dealerHandOutput := ""
	for _, card := range game.dealerHand {
		dealerHandOutput += card.Value + " " + card.Suit + ", "
	}

	io.WriteString(w, fmt.Sprintf("<p>Dealer's hand: %s </p>", dealerHandOutput))
	io.WriteString(w, fmt.Sprintf("<p>Dealer's score: %d</p>", game.dealerScore))
	io.WriteString(w, "<h1>You Tied!</h1>")
	DisplayNextRoundOrGameOver(w, game)
}

func DisplayNextRoundOrGameOver(w http.ResponseWriter, game *Game) {
	if game.chipAmount <= 0 {
		io.WriteString(w, "<h1>You're out of Chips! Time to Restart</h1>")
		io.WriteString(w, `
			<form method="post">
				<button name="action" value="restart">Restart</button>
			</form>
		`)
	} else {
		io.WriteString(w, `
			<form method="post">
				<button name="action" value="nextRound">Next Round</button>
			</form>
		`)
	}
}
