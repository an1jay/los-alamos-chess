package main

import (
	"github.com/an1jay/los-alamos-chess/game"
	"github.com/an1jay/los-alamos-chess/players"
)

func main() {
	// defer profile.Start().Stop()
	// testBoardMove()
	// timeBBReverse()
	g := Game{}
	// g.Play(players.HumanPlayer{}, players.HumanPlayer{}, true)

	ev1 := players.Evaluator{
		MaterialCoeff:      1,
		LegalMovesCoeff:    0.1,
		SquareControlCoeff: 0.08,
		MaterialWeights:    pawnrulesweights,
		WhiteSquareWeights: loosecentrecontrolweights,
		BlackSquareWeights: loosecentrecontrolweights,
	}
	_ = ev1
	ev2 := players.Evaluator{
		MaterialCoeff:      1,
		LegalMovesCoeff:    0.1,
		SquareControlCoeff: 0.05,
		MaterialWeights:    pawnrulesweights,
		WhiteSquareWeights: loosecentrecontrolweights,
		BlackSquareWeights: loosecentrecontrolweights,
	}

	m1 := players.CreateNewNeo(4, 8, &ev1, 6, true)

	m2 := players.CreateNewNeo(4, 8, &ev2, 6, true)

	// b := game.BoardFromMap(NewGame)
	// pos := game.NewPosition(b, game.White, 0, 0, []uint64{})
	pos := game.NewGamePosition()
	g.PlayFromPos(m1, m2, true, pos)
}

var standardmaterialweights = map[game.PieceType]float32{
	game.Pawn:   1,
	game.Knight: 3,
	game.Rook:   5,
	game.Queen:  9,
	game.King:   0,
}

var knightrulesweights = map[game.PieceType]float32{
	game.Pawn:   1,
	game.Knight: 3.5,
	game.Rook:   4.5,
	game.Queen:  8,
	game.King:   0,
}

var pawnrulesweights = map[game.PieceType]float32{
	game.Pawn:   1,
	game.Knight: 2.5,
	game.Rook:   3.5,
	game.Queen:  5,
	game.King:   0,
}

var tightcentrecontrolweights = map[game.Square]float32{
	game.C3: 1,
	game.C4: 1,
	game.D3: 1,
	game.D4: 1,
}

var loosecentrecontrolweights = map[game.Square]float32{
	game.C3: 1,
	game.C4: 1,
	game.D3: 1,
	game.D4: 1,

	game.B2: 0.5,
	game.B3: 0.5,
	game.B4: 0.5,
	game.B5: 0.5,

	game.E2: 0.5,
	game.E3: 0.5,
	game.E4: 0.5,
	game.E5: 0.5,

	game.C2: 0.5,
	game.C5: 0.5,
	game.D2: 0.5,
	game.D5: 0.5,
}
