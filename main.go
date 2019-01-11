package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
	"github.com/an1jay/los-alamos-chess/players"
	"github.com/an1jay/los-alamos-chess/players/evaluators"
)

func main() {
	// defer profile.Start().Stop()
	// testBoardMove()
	// timeBBReverse()
	g := Game{}
	// g.Play(players.HumanPlayer{}, players.HumanPlayer{}, true)

	// m1 := &players.AlphaBetaPlayer{
	// 	Ev:        evaluators.SecondEvaluator{},
	// 	MaxDepth:  5,
	// 	NodeCount: 0,
	// }

	m2 := &players.AlphaBetaPlayer{
		Ev:        evaluators.SecondEvaluator{},
		MaxDepth:  5,
		NodeCount: 0,
	}
	// r := &players.RandomPlayer{}
	m1 := &players.HumanPlayer{}
	b := game.BoardFromMap(NewGame)
	pos := game.NewPosition(b, game.White, 0, 0, []uint64{})
	g.PlayFromPos(m1, m2, true, pos)
}

func testMove() {
	b := game.BoardFromMap(InterestingPos)
	pos := game.NewPosition(b, game.White, 0, 0, []uint64{})

	fmt.Println("New Game pos")
	pos.Display(false)
	fmt.Println("Pseudo Legal ", pos.GeneratePseudoLegalMoves())
	fmt.Println("Legal ------ ", pos.GenerateLegalMoves())

	fmt.Println(pos.Move(&game.Ply{
		SourceSq:      game.D2,
		DestinationSq: game.D3,
		Promotion:     game.NoPieceType,
		Capture:       false,
		Side:          game.White,
	}))
	pos.Display(false)
	fmt.Println("Pseudo Legal ", pos.GeneratePseudoLegalMoves())
	fmt.Println("Legal ------ ", pos.GenerateLegalMoves())

	fmt.Println(pos.Move(&game.Ply{
		SourceSq:      game.A5,
		DestinationSq: game.A4,
		Promotion:     game.NoPieceType,
		Capture:       false,
		Side:          game.Black,
	}))
	pos.Display(false)
	fmt.Println("Pseudo Legal ", pos.GeneratePseudoLegalMoves())
	fmt.Println("Legal ------ ", pos.GenerateLegalMoves())

	fmt.Println(pos.Move(&game.Ply{
		SourceSq:      game.C1,
		DestinationSq: game.F4,
		Promotion:     game.NoPieceType,
		Capture:       false,
		Side:          game.White,
	}))
	pos.Display(false)
	fmt.Println("Pseudo Legal ", pos.GeneratePseudoLegalMoves())
	fmt.Println("Legal ------ ", pos.GenerateLegalMoves())

}

func testPseudoLegalMoves() {
	b := game.BoardFromMap(NewGame)
	b.Display(false)
	fmt.Println("Pseudo Legal Moves ")
	pos := game.NewPosition(b, game.White, 0, 0, []uint64{})
	pos.Display(false)
	plm := pos.GeneratePseudoLegalMoves()
	fmt.Println(plm)

}

func testCheck() {
	b := game.BoardFromMap(CheckPos)
	b.Display(false)
	fmt.Println("White in Check ", b.InCheck(game.White))
	fmt.Println("Black in Check ", b.InCheck(game.Black))
}
func testMovesVector() {
	b := game.BoardFromMap(CheckPos)
	for sq := 0; sq < game.NumSquaresInBoard; sq++ {
		fmt.Println("Square ", game.Square(sq))
		fmt.Println("")
		b.Display(false)
		fmt.Println("")
		b.MovesVector(game.Square(sq)).Display()
		fmt.Println("++++++++++++++++++++++++++++++++++++")
	}
}

func testGeneratedBBs() {
	for sq := 0; sq < game.NumSquaresInBoard; sq++ {
		fmt.Println(game.Square(sq))
		fmt.Println("King Moves")
		game.BBKingMoves[sq].Display()
		fmt.Println("Queen Moves")
		game.BBQueenMoves[sq].Display()
		fmt.Println("Rook Moves")
		game.BBRookMoves[sq].Display()
		fmt.Println("Knight Moves")
		game.BBKnightMoves[sq].Display()
		fmt.Println("White Pawn Pushes")
		game.BBWhitePawnPushes[sq].Display()
		fmt.Println("White Pawn Captures")
		game.BBWhitePawnPushes[sq].Display()
		fmt.Println("Black Pawn Pushes")
		game.BBBlackPawnPushes[sq].Display()
		fmt.Println("Black Pawn Captures")
		game.BBBlackPawnCaptures[sq].Display()
		fmt.Println("Diagonals")
		game.BBDiagonals[sq].Display()
		fmt.Println("AntiDiagonals")
		game.BBAntiDiagonals[sq].Display()
		fmt.Println("+++++++++++++++++++++++++++++++++")
	}
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
	//game.Ply{S1: E2, S2: E3, Promotion: NoPieceType, Capture: false, Cehck: false, InCheck: false})
}

func testBBReverse() {
	b := bitboardFromString("0000000000000000000000000000110000111111111111111111111111111111")
	fmt.Println("Before", b)
	b.Display()
	rb := b.Reverse()
	fmt.Println("After", rb)
	rb.Display()
}

func timeBBReverse() {
	b := bitboardFromString("0000000000000000000000000000110000111101111111111111111111111111")
	bb := bitboardFromString("0000000000000000000000000000110000111110111111111111111111111111")
	bbb := bitboardFromString("0000000000000000000000000000110000111111111111111111111111111111")
	bbbb := bitboardFromString("0000000000000000000000000000110000111100111111111111111111111111")
	bbbbb := bitboardFromString("0000000000000000000000000000110000111110011110011111111111111111")

	start := time.Now()
	b.Reverse()
	bb.Reverse()
	bbb.Reverse()
	bbbb.Reverse()
	bbbbb.Reverse()
	fmt.Printf("5 bbreverse took %d nanosecs\n", time.Since(start).Nanoseconds())

}

func bitboardFromString(str string) game.BitBoard {
	i, err := strconv.ParseUint(str, 2, 64)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	board := game.BitBoard(i)
	return board
}

func genSqNo(x, y int) int {
	return (y-1)*6 + (x - 1)
}

func sqNoToCoord(sq int) (int, int) {
	x := sq%6 + 1
	y := sq/6 + 1
	return x, y
}
