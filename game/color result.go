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

// Coefficient returns 1 for white, -1 for black, 0 for NoColor
func (c Color) Coefficient() int {
	switch c {
	case White:
		return 1
	case Black:
		return -1
	}
	return 0
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

// Evaluation returns 100 for a WhiteWin, -100 for a BlackWin, and 0 for a Draw
func (r Result) Evaluation() float32 {
	switch r {
	case BlackWin:
		return -100
	case WhiteWin:
		return 100
	case Draw:
		return 0
	}
	panic("Should not be evaluating in play position for win/loss/draw")
	return 0
}

// EvaluationCoefficient returns -1 for Black, 1 for White
func (r Result) EvaluationCoefficient() float32 {
	switch r {
	case BlackWin:
		return -1
	case WhiteWin:
		return 1
	case Draw:
		return 0
	}
	panic("Should not be evaluating in play position for win/loss/draw")
	return 0
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
