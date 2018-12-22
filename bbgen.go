package main

import (
	"fmt"

	"github.com/an1jay/los-alamos-chess/game"
)

func generateRookMoves() []game.BitBoard {
	moves := []game.BitBoard{}
	for sq := 0; sq < game.NumSquaresInBoard; sq++ {
		square := game.Square(sq)
		// fmt.Println("Square ", square)
		// (square.RankBB() ^ square.FileBB()).Display()
		moves = append(moves, square.RankBB()^square.FileBB())
		fmt.Println(uint(square.RankBB()^square.FileBB()), ",")
		// fmt.Println()
	}
	// fmt.Println(moves)
	return moves
}

func generateKingMoves() []game.BitBoard {
	moves := []game.BitBoard{}

	kingMoveVectors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for _, vec := range kingMoveVectors {
			sqx, sqy := sqNoToCoord(sq)
			sqx += vec[0]
			sqy += vec[1]
			// fmt.Println("In loop, sqx, sqy", sqx, sqy)
			if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
				e[game.Square(genSqNo(sqx, sqy))] = true
				// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
			}
		}
		b := game.BitBoardFromMap(e)
		fmt.Println(uint(b), ",")
		// b.Display()

		moves = append(moves, b)
	}
	return moves
}

func generateKnightMoves() []game.BitBoard {
	moves := []game.BitBoard{}

	knightMoveVectors := [][]int{{2, 1}, {-2, 1}, {2, -1}, {-2, -1},
		{1, 2}, {-1, 2}, {1, -2}, {-1, -2},
	}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for _, vec := range knightMoveVectors {
			sqx, sqy := sqNoToCoord(sq)
			sqx += vec[0]
			sqy += vec[1]
			// fmt.Println("In loop, sqx, sqy", sqx, sqy)
			if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
				e[game.Square(genSqNo(sqx, sqy))] = true
				// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
			}
		}
		b := game.BitBoardFromMap(e)
		fmt.Println(uint(b), ",")
		// b.Display()

		moves = append(moves, b)
	}
	return moves
}

func generateQueenMoves() []game.BitBoard {
	moves := []game.BitBoard{}

	bishopMoveVectors := [][]int{{1, -1}, {1, 1}, {-1, 1}, {-1, -1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for i := 1; i < 7; i++ {
			for _, vec := range bishopMoveVectors {
				sqx, sqy := sqNoToCoord(sq)
				sqx += i * vec[0]
				sqy += i * vec[1]
				// fmt.Println("In loop, sqx, sqy", sqx, sqy)
				if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
					e[game.Square(genSqNo(sqx, sqy))] = true
					// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
				}
			}
		}
		b := game.BitBoardFromMap(e)

		square := game.Square(sq)
		b = b ^ square.RankBB() ^ square.FileBB()

		fmt.Println(uint(b), ",")
		// b.Display()
		moves = append(moves, b)
	}
	return moves
}

func generateWhitePawnPushes() []game.BitBoard {
	moves := []game.BitBoard{}

	WhitePawnPushes := [][]int{{0, 1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for _, vec := range WhitePawnPushes {
			sqx, sqy := sqNoToCoord(sq)
			sqx += vec[0]
			sqy += vec[1]
			// fmt.Println("In loop, sqx, sqy", sqx, sqy)
			if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
				e[game.Square(genSqNo(sqx, sqy))] = true
				// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
			}
		}
		b := game.BitBoardFromMap(e)
		fmt.Println(uint(b), ",")
		// b.Display()

		moves = append(moves, b)
	}
	return moves
}

func generateWhitePawnCaptures() []game.BitBoard {
	moves := []game.BitBoard{}

	WhitePawnCaptures := [][]int{{1, 1}, {-1, 1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for _, vec := range WhitePawnCaptures {
			sqx, sqy := sqNoToCoord(sq)
			sqx += vec[0]
			sqy += vec[1]
			// fmt.Println("In loop, sqx, sqy", sqx, sqy)
			if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
				e[game.Square(genSqNo(sqx, sqy))] = true
				// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
			}
		}
		b := game.BitBoardFromMap(e)
		fmt.Println(uint(b), ",")
		// b.Display()

		moves = append(moves, b)
	}
	return moves
}

func generateBlackPawnPushes() []game.BitBoard {
	moves := []game.BitBoard{}

	BlackPawnPushes := [][]int{{0, -1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for _, vec := range BlackPawnPushes {
			sqx, sqy := sqNoToCoord(sq)
			sqx += vec[0]
			sqy += vec[1]
			// fmt.Println("In loop, sqx, sqy", sqx, sqy)
			if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
				e[game.Square(genSqNo(sqx, sqy))] = true
				// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
			}
		}
		b := game.BitBoardFromMap(e)
		fmt.Println(uint(b), ",")
		// b.Display()

		moves = append(moves, b)
	}
	return moves
}

func generateBlackPawnCaptures() []game.BitBoard {
	moves := []game.BitBoard{}

	BlackPawnCaptures := [][]int{{1, -1}, {-1, -1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for _, vec := range BlackPawnCaptures {
			sqx, sqy := sqNoToCoord(sq)
			sqx += vec[0]
			sqy += vec[1]
			// fmt.Println("In loop, sqx, sqy", sqx, sqy)
			if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
				e[game.Square(genSqNo(sqx, sqy))] = true
				// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
			}
		}
		b := game.BitBoardFromMap(e)
		fmt.Println(uint(b), ",")
		// b.Display()

		moves = append(moves, b)
	}
	return moves
}

func generateDiagonals() []game.BitBoard {
	moves := []game.BitBoard{}

	bishopMoveVectors := [][]int{{-1, -1}, {1, 1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for i := 1; i < 7; i++ {
			for _, vec := range bishopMoveVectors {
				sqx, sqy := sqNoToCoord(sq)
				sqx += i * vec[0]
				sqy += i * vec[1]
				// fmt.Println("In loop, sqx, sqy", sqx, sqy)
				if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
					e[game.Square(genSqNo(sqx, sqy))] = true
					// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
				}
			}
		}
		b := game.BitBoardFromMap(e)

		fmt.Println(uint(b), ",")
		// b.Display()
		moves = append(moves, b)
	}
	return moves
}

func generateAntiDiagonals() []game.BitBoard {
	moves := []game.BitBoard{}

	bishopMoveVectors := [][]int{{-1, 1}, {1, -1}}

	// for sq := 0; sq < game.NumSquaresInBoard; sq++ {
	for sq := 0; sq < 36; sq++ {
		// fmt.Println("Square ", game.Square(sq))
		e := map[game.Square]bool{
			game.A1: false, game.B1: false, game.C1: false, game.D1: false, game.E1: false, game.F1: false,
			game.A2: false, game.B2: false, game.C2: false, game.D2: false, game.E2: false, game.F2: false,
			game.A3: false, game.B3: false, game.C3: false, game.D3: false, game.E3: false, game.F3: false,
			game.A4: false, game.B4: false, game.C4: false, game.D4: false, game.E4: false, game.F4: false,
			game.A5: false, game.B5: false, game.C5: false, game.D5: false, game.E5: false, game.F5: false,
			game.A6: false, game.B6: false, game.C6: false, game.D6: false, game.E6: false, game.F6: false,
		}
		// fmt.Println(e)
		for i := 1; i < 7; i++ {
			for _, vec := range bishopMoveVectors {
				sqx, sqy := sqNoToCoord(sq)
				sqx += i * vec[0]
				sqy += i * vec[1]
				// fmt.Println("In loop, sqx, sqy", sqx, sqy)
				if sqx >= 1 && sqx <= 6 && sqy >= 1 && sqy <= 6 {
					e[game.Square(genSqNo(sqx, sqy))] = true
					// fmt.Println("Valid SQ ", game.Square(genSqNo(sqx, sqy)))
				}
			}
		}
		b := game.BitBoardFromMap(e)

		fmt.Println(uint(b), ",")
		// b.Display()
		moves = append(moves, b)
	}
	return moves
}
