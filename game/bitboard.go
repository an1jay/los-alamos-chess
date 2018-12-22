package game

import (
	"fmt"
	"math/bits"
)

// Bitboard encodes as following:
// 0000000000000000000000000000 | 1  1  1  1  1  1  | 1  1  1  1  1  1  | 1  1  1  1  1  1  | 1  1  1  1  1  1  | 1  1  1  1  1  1  | 1  1  1  1  1  1
//		Squares					| F6 E6 D6 C6 B6 A6 | F5 E5 D5 C5 B5 A5 | F4 E4 D4 C4 B4 A4 | F3 E3 D3 C3 B3 A3 | F2 E2 D2 C2 B2 A2 | F1 E1 D1 C1 B1 A1
//		Number					| 35 34 33 32 31 30 | 29 28 27 26 25 24 | 23 22 21 20 19 18 | 17 16 15 14 13 12 | 11 10 9  8  7  6  | 5  4  3  2  1  0

// BitBoard is a board representation encoded in an unsigned 64-bit integer.
// Stores the 36 squares in the 36 least significant bits. All other bits are zero.
type BitBoard uint64

// BitBoardFromMap generates a bitboard, from a Square->bool map
func BitBoardFromMap(m map[Square]bool) BitBoard {
	var rv BitBoard
	var one uint64 = 1
	for sq, fill := range m {
		if fill {
			rv += BitBoard(one << uint(sq))
		}
	}
	return rv
}

// Mapping returns a bitboard from a Square -> bool map, =1 if bool for corresponding square is true.
func (b BitBoard) Mapping() map[Square]bool {
	m := make(map[Square]bool)
	for sq := 0; sq < NumSquaresInBoard; sq++ {
		var comp uint64 = 1
		comp = comp >> uint(sq)
		if bits.OnesCount64(comp&uint64(b)) == 1 {
			m[Square(sq)] = true
		}
	}
	return m
}

// SetSquareOnBB returns a bitboard with the specific square set to 1 or 0 (true or false).
func (b BitBoard) SetSquareOnBB(sq Square, value bool) BitBoard {
	if value {
		return b | BitBoard(1<<uint(sq))
	}
	return ^BitBoard(1<<uint(sq)) & b
}

// Occupied returns true if the square's bitboard position is 1.
func (b BitBoard) Occupied(sq Square) bool {
	return (uint64(b) >> uint64(sq) & 1) == 1
}

// Reverse returns a reversed version of the bitboard - useful in calculating legal moves
func (b BitBoard) Reverse() BitBoard {
	return BitBoard(bits.Reverse64(uint64(b)) >> 28)
}

// String returns a 36 character string of 1s and 0s starting with most significant bit.
// Implements the fmt.Stringer interface.
func (b BitBoard) String() string {
	return fmt.Sprintf("%064b", uint64(b))
}

// FuncDisplay outputs a string representation given by function f to stdout
func (b BitBoard) FuncDisplay(f func(Square) string) {
	fmt.Printf(
		"   A B C D E F\n"+
			"6. %s %s %s %s %s %s\n"+
			"5. %s %s %s %s %s %s\n"+
			"4. %s %s %s %s %s %s\n"+
			"3. %s %s %s %s %s %s\n"+
			"2. %s %s %s %s %s %s\n"+
			"1. %s %s %s %s %s %s\n",
		f(A6), f(B6), f(C6), f(D6), f(E6), f(F6),
		f(A5), f(B5), f(C5), f(D5), f(E5), f(F5),
		f(A4), f(B4), f(C4), f(D4), f(E4), f(F4),
		f(A3), f(B3), f(C3), f(D3), f(E3), f(F3),
		f(A2), f(B2), f(C2), f(D2), f(E2), f(F2),
		f(A1), f(B1), f(C1), f(D1), f(E1), f(F1))
}

// Display prints a representation to stdout.
func (b BitBoard) Display() {
	b.FuncDisplay(func(sq Square) string { return BoolToIntString(b.Occupied(sq)) })
}

//--------------------------------------------------------------------------------

const (
	// BBMask is equal to 0000000000000000000000000000111111111111111111111111111111111111
	BBMask BitBoard = 68719476735

	// BBValidityCheck is used to check whether a bitboard is valid
	// Any bitboard b is only valid iff b & validBBCheck == 0
	// Equal to 1111111111111111111111111111000000000000000000000000000000000000
	BBValidityCheck BitBoard = 18446744004990074880
)

// Bitboards for each file
const (
	BBFileA BitBoard = 1090785345
	BBFileB BitBoard = 2181570690
	BBFileC BitBoard = 4363141380
	BBFileD BitBoard = 8726282760
	BBFileE BitBoard = 17452565520
	BBFileF BitBoard = 34905131040
)

// Bitboards for each rank
const (
	BBRank1 BitBoard = 63
	BBRank2 BitBoard = 4032
	BBRank3 BitBoard = 258048
	BBRank4 BitBoard = 16515072
	BBRank5 BitBoard = 1056964608
	BBRank6 BitBoard = 67645734912
)
const (
	// NumSquaresInBoard is 36, the number of squares on the board
	NumSquaresInBoard int = 36
	// NumSquaresInRow is 6, the number of squares in a row
	NumSquaresInRow int = 6
)

var (
	// EmptyBitBoardMap is a map from all Squares to false - i.e. representing an empty bitboard.
	EmptyBitBoardMap = map[Square]bool{
		A1: false, B1: false, C1: false, D1: false, E1: false, F1: false,
		A2: false, B2: false, C2: false, D2: false, E2: false, F2: false,
		A3: false, B3: false, C3: false, D3: false, E3: false, F3: false,
		A4: false, B4: false, C4: false, D4: false, E4: false, F4: false,
		A5: false, B5: false, C5: false, D5: false, E5: false, F5: false,
		A6: false, B6: false, C6: false, D6: false, E6: false, F6: false,
	}

	// BBFiles is an array of bitboards for each file
	BBFiles = [6]BitBoard{BBFileA, BBFileB, BBFileC, BBFileD, BBFileE, BBFileF}

	// BBRanks is an array of bitboards for each rank
	BBRanks = [6]BitBoard{BBRank1, BBRank2, BBRank3, BBRank4, BBRank5, BBRank6}
)
