package blackjack

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	gameSessions = make(map[string]*Game)
	mu           sync.Mutex
	suits        = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values       = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
)

// Card is the base unit in blackjack
type Card struct {
	Suit  string
	Value string
	Score int
}

type Game struct {
	Deck         []Card
	playerScore  int
	dealerScore  int
	playerHand   []Card
	dealerHand   []Card
	chipAmount   int
	deckBookMark int
	activeBet    int
}

// Seed the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
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

// shuffleDeck function to shuffle the cards in a random order
func shuffleDeck(deck []Card) []Card {
	shuffledDeck := append([]Card(nil), deck...) // create a copy of the deck
	rand.Shuffle(len(shuffledDeck), func(i, j int) {
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	})
	return shuffledDeck
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

func initializeGame() *Game {
	return &Game{
		Deck:         shuffleDeck(createDeck()),
		playerScore:  0,
		dealerScore:  0,
		playerHand:   []Card{},
		dealerHand:   []Card{},
		chipAmount:   100,
		deckBookMark: 0,
		activeBet:    0,
	}
}

// Helper to get the game from the session ID
func getGameSession(sessionID string) *Game {
	mu.Lock()
	defer mu.Unlock()
	game, exists := gameSessions[sessionID]
	if !exists {
		game = initializeGame()
		gameSessions[sessionID] = game
	}
	return game
}

func (g *Game) hit(w http.ResponseWriter) {
	g.playerHand = append(g.playerHand, g.Deck[g.deckBookMark])
	g.deckBookMark++
	g.playerScore = calculateHandScore(g.playerHand)
	if g.playerScore > 21 {
		DisplayPlayerBust(w, g)
	} else {
		DisplayHitStandState(w, g)
	}
}

func (g *Game) stand(w http.ResponseWriter) {
	for g.dealerScore < 17 {
		g.dealerHand = append(g.dealerHand, g.Deck[g.deckBookMark])
		g.deckBookMark++
		g.dealerScore = calculateHandScore(g.dealerHand)
	}

	g.determineRoundEnd(w)
}

func (g *Game) determineRoundEnd(w http.ResponseWriter) {
	if g.dealerScore > 21 {
		// Dealer Bust
		g.chipAmount += g.activeBet * 2
		g.activeBet = 0
		DisplayDealerBust(w, g)
	} else if g.dealerScore > g.playerScore {
		// Player Loss
		g.activeBet = 0
		DisplayPlayerLose(w, g)
	} else if g.dealerScore < g.playerScore {
		// Player Win
		g.chipAmount += g.activeBet * 2
		g.activeBet = 0
		DisplayPlayerWin(w, g)
	} else if g.dealerScore == g.playerScore {
		// Game Tie
		g.chipAmount += g.activeBet
		g.activeBet = 0
		DisplayPlayerTie(w, g)
	}
}

func (g *Game) startRound() {

	// Deal Cards
	g.playerHand = []Card{g.Deck[g.deckBookMark], g.Deck[g.deckBookMark+2]}
	g.dealerHand = []Card{g.Deck[g.deckBookMark+1], g.Deck[g.deckBookMark+3]}
	g.deckBookMark += 4

	// Calculate Scores
	g.playerScore = calculateHandScore(g.playerHand)
	g.dealerScore = calculateHandScore(g.dealerHand)
	return
}

func (g *Game) betPlaced(w http.ResponseWriter) {
	g.startRound()
	DisplayHitStandState(w, g)
}

func (g *Game) checkBet(val int) bool {
	if val <= 0 {
		return false
	}

	if val > g.chipAmount {
		return false
	}

	return true
}

func (g *Game) restart(w http.ResponseWriter) {
	*g = *initializeGame()
	DisplayWelcomeScreen(w, g)
}

func GetBlackjack(w http.ResponseWriter, r *http.Request) {

	sessionID := "default"

	game := getGameSession(sessionID)

	if r.Method == http.MethodPost {
		action := r.FormValue("action")

		if action == "hit" {
			game.hit(w)
		}
		if action == "stand" {
			game.stand(w)

		}
		if action == "placeBet" {

			properInput := false
			betValAsString := r.FormValue("bet")
			betVal, err := strconv.Atoi(betValAsString)
			if err != nil {
				log.Println(err)
			} else {
				properInput = game.checkBet(betVal)
			}
			if !properInput {
				DisplayInvalidInputScreen(w, game)
			} else {
				game.activeBet = betVal
				game.chipAmount -= betVal
				game.betPlaced(w)
			}
		}
		if action == "start" {
			DisplayPlaceBetScreen(w, game)
		}
		if action == "restart" {
			game.restart(w)
		}
		if action == "nextRound" {
			DisplayPlaceBetScreen(w, game)
		}
	} else {
		DisplayWelcomeScreen(w, game)
	}
}
