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

// // Reverse returns a bitboard where the bit order is reversed.
// // Implementation from: http://stackoverflow.com/questions/746171/best-algorithm-for-bit-reversal-from-msb-lsb-to-lsb-msb-in-c
// func (b BitBoard) Reverse() BitBoard {
// 	return (BitBoard((bitReverseLookupTable[b&0xff]<<56)|
// 		(bitReverseLookupTable[(b>>8)&0xff]<<48)|
// 		(bitReverseLookupTable[(b>>16)&0xff]<<40)|
// 		(bitReverseLookupTable[(b>>24)&0xff]<<32)|
// 		(bitReverseLookupTable[(b>>32)&0xff]<<24)|
// 		(bitReverseLookupTable[(b>>40)&0xff]<<16)|
// 		(bitReverseLookupTable[(b>>48)&0xff]<<8)|
// 		(bitReverseLookupTable[(b>>56)&0xff])) >> 28)
// }

// var bitReverseLookupTable = []uint64{
// 	0x00, 0x80, 0x40, 0xC0, 0x20, 0xA0, 0x60, 0xE0, 0x10, 0x90, 0x50, 0xD0, 0x30, 0xB0, 0x70, 0xF0,
// 	0x08, 0x88, 0x48, 0xC8, 0x28, 0xA8, 0x68, 0xE8, 0x18, 0x98, 0x58, 0xD8, 0x38, 0xB8, 0x78, 0xF8,
// 	0x04, 0x84, 0x44, 0xC4, 0x24, 0xA4, 0x64, 0xE4, 0x14, 0x94, 0x54, 0xD4, 0x34, 0xB4, 0x74, 0xF4,
// 	0x0C, 0x8C, 0x4C, 0xCC, 0x2C, 0xAC, 0x6C, 0xEC, 0x1C, 0x9C, 0x5C, 0xDC, 0x3C, 0xBC, 0x7C, 0xFC,
// 	0x02, 0x82, 0x42, 0xC2, 0x22, 0xA2, 0x62, 0xE2, 0x12, 0x92, 0x52, 0xD2, 0x32, 0xB2, 0x72, 0xF2,
// 	0x0A, 0x8A, 0x4A, 0xCA, 0x2A, 0xAA, 0x6A, 0xEA, 0x1A, 0x9A, 0x5A, 0xDA, 0x3A, 0xBA, 0x7A, 0xFA,
// 	0x06, 0x86, 0x46, 0xC6, 0x26, 0xA6, 0x66, 0xE6, 0x16, 0x96, 0x56, 0xD6, 0x36, 0xB6, 0x76, 0xF6,
// 	0x0E, 0x8E, 0x4E, 0xCE, 0x2E, 0xAE, 0x6E, 0xEE, 0x1E, 0x9E, 0x5E, 0xDE, 0x3E, 0xBE, 0x7E, 0xFE,
// 	0x01, 0x81, 0x41, 0xC1, 0x21, 0xA1, 0x61, 0xE1, 0x11, 0x91, 0x51, 0xD1, 0x31, 0xB1, 0x71, 0xF1,
// 	0x09, 0x89, 0x49, 0xC9, 0x29, 0xA9, 0x69, 0xE9, 0x19, 0x99, 0x59, 0xD9, 0x39, 0xB9, 0x79, 0xF9,
// 	0x05, 0x85, 0x45, 0xC5, 0x25, 0xA5, 0x65, 0xE5, 0x15, 0x95, 0x55, 0xD5, 0x35, 0xB5, 0x75, 0xF5,
// 	0x0D, 0x8D, 0x4D, 0xCD, 0x2D, 0xAD, 0x6D, 0xED, 0x1D, 0x9D, 0x5D, 0xDD, 0x3D, 0xBD, 0x7D, 0xFD,
// 	0x03, 0x83, 0x43, 0xC3, 0x23, 0xA3, 0x63, 0xE3, 0x13, 0x93, 0x53, 0xD3, 0x33, 0xB3, 0x73, 0xF3,
// 	0x0B, 0x8B, 0x4B, 0xCB, 0x2B, 0xAB, 0x6B, 0xEB, 0x1B, 0x9B, 0x5B, 0xDB, 0x3B, 0xBB, 0x7B, 0xFB,
// 	0x07, 0x87, 0x47, 0xC7, 0x27, 0xA7, 0x67, 0xE7, 0x17, 0x97, 0x57, 0xD7, 0x37, 0xB7, 0x77, 0xF7,
// 	0x0F, 0x8F, 0x4F, 0xCF, 0x2F, 0xAF, 0x6F, 0xEF, 0x1F, 0x9F, 0x5F, 0xDF, 0x3F, 0xBF, 0x7F, 0xFF,
// }

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
