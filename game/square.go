package game

const (
	numSquaresInBoard = 36
	numSquaresInRow   = 6
)

// Square is a combination of rank and file, one of the 36 squares on the board.
type Square int8

// Rank gets the Rank of the square
func (sq Square) Rank() Rank {
	return Rank(int(sq) / numSquaresInRow)
}

// File gets the File of the square
func (sq Square) File() File {
	return File(int(sq) % numSquaresInRow)
}

// BitBoard returns a blank bitboard with the square = 1
func (sq Square) BitBoard() BitBoard {
	return BitBoard(1 << uint8(sq))
}

// Color gets the color of the square
func (sq Square) Color() Color {
	if ((sq / 6) % 2) == (sq % 2) {
		return Black
	}
	return White
}

// String gives a string representation of a square -- e.g. "a1"
// Implements the fmt.Stringer interface.
func (sq Square) String() string {
	return sq.File().String() + sq.Rank().String()
}

// GetSquare constructs a Square from Rank and File
func GetSquare(f File, r Rank) Square {
	return Square(int(r)*numSquaresInRow + int(f))
}

// RankBB returns a bitboard for the rank of the square
func (sq Square) RankBB() BitBoard {
	return BBRanks[sq.Rank()]
}

// FileBB returns a bitboard for the file of the square
func (sq Square) FileBB() BitBoard {
	return BBFiles[sq.File()]
}

//--------------------------------------------------------------------------------

var (
	// StringToSquareMap converts from string to Square
	StringToSquareMap = map[string]Square{
		"a1": A1, "a2": A2, "a3": A3, "a4": A4, "a5": A5, "a6": A6,
		"b1": B1, "b2": B2, "b3": B3, "b4": B4, "b5": B5, "b6": B6,
		"c1": C1, "c2": C2, "c3": C3, "c4": C4, "c5": C5, "c6": C6,
		"d1": D1, "d2": D2, "d3": D3, "d4": D4, "d5": D5, "d6": D6,
		"e1": E1, "e2": E2, "e3": E3, "e4": E4, "e5": E5, "e6": E6,
		"f1": F1, "f2": F2, "f3": F3, "f4": F4, "f5": F5, "f6": F6,
	}

	// SquareToStringMap converts from square to string
	SquareToStringMap = map[Square]string{
		A1: "a1", A2: "a2", A3: "a3", A4: "a4", A5: "a5", A6: "a6",
		B1: "b1", B2: "b2", B3: "b3", B4: "b4", B5: "b5", B6: "b6",
		C1: "c1", C2: "c2", C3: "c3", C4: "c4", C5: "c5", C6: "c6",
		D1: "d1", D2: "d2", D3: "d3", D4: "d4", D5: "d5", D6: "d6",
		E1: "e1", E2: "e2", E3: "e3", E4: "e4", E5: "e5", E6: "e6",
		F1: "f1", F2: "f2", F3: "f3", F4: "f4", F5: "f5", F6: "f6",
	}
)

// Enumerating all Squares on the board
const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	A2
	B2
	C2
	D2
	E2
	F2
	A3
	B3
	C3
	D3
	E3
	F3
	A4
	B4
	C4
	D4
	E4
	F4
	A5
	B5
	C5
	D5
	E5
	F5
	A6
	B6
	C6
	D6
	E6
	F6
)
