package game

// Color is one of the two colors that squares and peices can be - Black and White
type Color int8

// Enumerating all possible colors.
const (
	Black Color = iota
	White
	NoColor
)

// Other returns the opposite color
func (c Color) Other() Color {
	if c == Black {
		return White
	}
	return Black
}

func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	}
	return "No Color"
}

//--------------------------------------------------------------------------------

// Result is one of the 4 Results a game can be in: InPlay, BlackWin, WhiteWin, Draw
type Result int8

// Enumerating all possible GameStates
const (
	InPlay = iota
	BlackWin
	WhiteWin
	Draw
)

// GameOver returns true if the game state is not equal to GameOver
func (r Result) GameOver() bool {
	if r == InPlay {
		return false
	}
	return true
}

func (r Result) String() string {
	switch r {
	case BlackWin:
		return "Black Win"
	case WhiteWin:
		return "White Win"
	case Draw:
		return "Draw"
	}
	return "In Play"
}

// NewResultWin returns a win for the color c. If c is NoColor, then Draw is returned.
func NewResultWin(c Color) Result {
	if c == White {
		return WhiteWin
	} else if c == Black {
		return BlackWin
	}
	return Draw
}
