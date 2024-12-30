package cards

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateCardScore(t *testing.T) {

	type testCase struct {
		actualValue   string
		expectedScore int
	}

	testCases := []testCase{
		{
			actualValue:   "A",
			expectedScore: 11,
		},
		{
			actualValue:   "J",
			expectedScore: 10,
		},
		{
			actualValue:   "5",
			expectedScore: 5,
		},
		{
			actualValue:   "10",
			expectedScore: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.actualValue, func(t *testing.T) {
			score := calculateCardScore(tc.actualValue)
			require.Equal(t, tc.expectedScore, score)
		})

	}
}

func TestCalculateHandScore(t *testing.T) {
	type testCase struct {
		actualHand    []Card
		expectedScore int
	}

	testCases := map[string]testCase{

		"normalInts": {
			actualHand: []Card{
				{Suit: "Hearts", Value: "2", Score: 2},
				{Suit: "Hearts", Value: "8", Score: 8},
			},
			expectedScore: 10,
		},
		"faceCards": {
			actualHand: []Card{
				{Suit: "Hearts", Value: "J", Score: 10},
				{Suit: "Hearts", Value: "K", Score: 10},
			},
			expectedScore: 20,
		},
		"threeCards": {
			actualHand: []Card{
				{Suit: "Hearts", Value: "2", Score: 2},
				{Suit: "Hearts", Value: "J", Score: 10},
				{Suit: "Hearts", Value: "9", Score: 9},
			},
			expectedScore: 21,
		},
		"Aces": {
			actualHand: []Card{
				{Suit: "Hearts", Value: "A", Score: 2},
				{Suit: "Hearts", Value: "A", Score: 10},
			},
			expectedScore: 12,
		},
		"BustScore": {
			actualHand: []Card{
				{Suit: "Hearts", Value: "J", Score: 10},
				{Suit: "Hearts", Value: "K", Score: 10},
				{Suit: "Hearts", Value: "4", Score: 4},
			},
			expectedScore: 24,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			score := calculateHandScore(tc.actualHand)
			require.Equal(t, tc.expectedScore, score)
		})
	}
}
