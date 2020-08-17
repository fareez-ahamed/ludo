// Package ludo has the type and logic for Ludo.
package ludo

// Color represensts one of the four colors
type Color string

// Constants for Colors
const (
	Red    = Color("Red")
	Green  = Color("Green")
	Blue   = Color("Blue")
	Yellow = Color("Yellow")
)

// Player represents a Player in Ludo.
type Player interface{}

// NoPlayer represents an empty slot in the list
// of players.
type NoPlayer struct{}

// ActivePlayer represents an active player in the game.
type ActivePlayer struct {
	Color Color
	Pawn  [4]Pawn
}

// Winner represents a won player along with the winning position.
type Winner int

// Pawn represents one of four pawns of a player.
type Pawn interface{}

// NotStarted represents a Pawn which is not started yet.
type NotStarted struct{}

// Unsafe represents a Pawn sitting on a safe position
type Unsafe int

// Safe represents a Pawn sitting on unsafe position
type Safe int

// Completed represents a Pawn which is has completed
type Completed struct{}

// DiceResult represents the outcome of a dice.
type DiceResult int
