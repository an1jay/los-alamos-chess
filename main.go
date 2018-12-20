package main

import (
	"fmt"

	"github.com/an1jay/los-alamos-chess/game"
)

func main() {
	testBoardMove()
}

func testBoardMove() {
	m := map[game.Square]game.Piece{
		game.A1: game.WhiteRook,
		game.B1: game.WhiteKnight,
		game.C1: game.WhiteQueen,
		game.D1: game.WhiteKing,
		game.E1: game.WhiteKnight,
		game.F1: game.WhiteRook,
		game.A2: game.WhitePawn,
		game.B2: game.WhitePawn,
		game.C2: game.WhitePawn,
		game.D2: game.WhitePawn,
		game.E2: game.WhitePawn,
		game.F2: game.WhitePawn,
		game.A3: game.NoPiece,
		game.B3: game.NoPiece,
		game.C3: game.NoPiece,
		game.D3: game.NoPiece,
		game.E3: game.NoPiece,
		game.F3: game.NoPiece,
		game.A4: game.NoPiece,
		game.B4: game.NoPiece,
		game.C4: game.NoPiece,
		game.D4: game.NoPiece,
		game.E4: game.NoPiece,
		game.F4: game.NoPiece,
		game.A5: game.BlackPawn,
		game.B5: game.BlackPawn,
		game.C5: game.BlackPawn,
		game.D5: game.BlackPawn,
		game.E5: game.BlackPawn,
		game.F5: game.BlackPawn,
		game.A6: game.BlackRook,
		game.B6: game.BlackKnight,
		game.C6: game.BlackQueen,
		game.D6: game.BlackKing,
		game.E6: game.BlackKnight,
		game.F6: game.BlackRook,
	}

	b := game.BoardFromMap(m)
	fmt.Println("Before: ")
	b.Display(true)
	b.Move(game.E2, game.E3, game.NoPiece)
	fmt.Println("After: ")
	b.Display(true)
	//engine.Ply{S1: E2, S2: E3, Promotion: NoPieceType, Capture: false, Cehck: false, InCheck: false})
}
