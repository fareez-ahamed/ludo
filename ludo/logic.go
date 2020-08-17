package ludo

import "math/rand"

// RollDice and returns the outcome.
func RollDice() DiceResult {
	return DiceResult(rand.Intn(100) + 1)
}
