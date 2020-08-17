package ludo

import "testing"

func TestRollDice(t *testing.T) {
	got := RollDice()
	if got >= 1 && got <= 6 {
		return
	}
	t.Errorf("Expected DiceResult between 1 & 6; Got %d", got)
}

func TestNewPawnState(t *testing.T) {
	type ts struct {
		state     Pawn
		number    int
		nextState Pawn
	}
	tests := []ts{
		{NotStarted{}, 3, NotStarted{}},
		{Completed{}, 3, Completed{}},
		{NotStarted{}, 1, Safe(1)},
		{Safe(1), 1, Unsafe(2)},
		{Unsafe(6), 3, Safe(9)},
		{Safe(14), 2, Unsafe(16)},
		{Unsafe(52), 1, Safe(53)},
		{Unsafe(52), 3, Safe(55)},
		{Safe(58), 1, Completed{}},
		{Safe(53), 6, Completed{}},
		{Safe(55), 6, Safe(55)},
	}

	for _, test := range tests {
		if got := NewPawnState(test.state, DiceResult(test.number)); got != test.nextState {
			t.Errorf("Expected %T %v, Got %T %v", test.nextState, test.nextState, got, got)
		}
	}
}
