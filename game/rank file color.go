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

//--------------------------------------------------------------------------------

// File is one of the 6 columns of the board (e.g. the 'a' File)
type File int8

const fileChars string = "abcdef"

// String gives a string representation of the file - e.g. 'a' file
// Implements the fmt.Stringer interface.
func (f File) String() string {
	return fileChars[f : f+1]
}

// Enumerating all Files on the board
const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
)

//--------------------------------------------------------------------------------

// Rank is one of the 6 rows of the board (e.g. '1st' Rank)
type Rank int8

const rankChars string = "123456"

// String gives a string representation of the rank - e.g. rank '1'
// Implements the fmt.Stringer interface.
func (r Rank) String() string {
	return rankChars[r : r+1]
}

// Enumerating all 6 Ranks on the board
const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
)
