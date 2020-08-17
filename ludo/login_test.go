package ludo

import "testing"

func TestRollDice(t *testing.T) {
	got := RollDice()
	if got >= 1 && got <= 6 {
		return
	}
	t.Errorf("Expected DiceResult between 1 & 6; Got %d", got)
}
