package ludo

import "math/rand"

// RollDice and returns the outcome.
func RollDice() DiceResult {
	return DiceResult(rand.Intn(6) + 1)
}

// NewPawnState returns the new state of the Pawn when we give a DiceResult
func NewPawnState(pawn Pawn, dice DiceResult) Pawn {
	switch p := pawn.(type) {
	case NotStarted:
		if dice == 1 {
			return Safe(1)
		}
		return p
	case Safe:
		n := int(p) + int(dice)
		if n == 59 {
			return Completed{}
		} else if n > 59 {
			return p
		}
		return getPawnType(n)
	case Unsafe:
		return getPawnType(int(p) + int(dice))
	case Completed:
		return p
	}
	return pawn
}

func isSafe(n int) bool {
	switch n {
	case 1, 9, 14, 22, 27, 35, 40, 49,
		53, 54, 55, 56, 57, 58:
		return true
	}
	return false
}

func getPawnType(n int) Pawn {
	if isSafe(n) {
		return Safe(n)
	}
	return Unsafe(n)
}
