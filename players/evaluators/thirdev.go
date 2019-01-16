package evaluators

import "github.com/an1jay/los-alamos-chess/game"

// ThirdEvaluator evaluates a position
type ThirdEvaluator struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (b ThirdEvaluator) Evaluate(pos *game.Position) float32 {
	// If game over, return evaluation of the game over state
	res := pos.Result()
	if res != game.InPlay {
		return res.Evaluation()
	}

	WhiteMaterial := pos.Bd.MaterialCount(game.White)
	BlackMaterial := pos.Bd.MaterialCount(game.Black)

	whiteMaterialScore := NewMaterialWeighter(WhiteMaterial)
	blackMaterialScore := NewMaterialWeighter(BlackMaterial)
	materialScore := whiteMaterialScore - blackMaterialScore

	backUpTurn := pos.Turn

	pos.Turn = game.White
	WhiteLegalMoves := float32(pos.GenerateCountOfPseudoLegalMoves())

	pos.Turn = game.Black
	BlackLegalMoves := float32(pos.GenerateCountOfPseudoLegalMoves())

	pos.Turn = backUpTurn

	legalMoveScore := WhiteLegalMoves - BlackLegalMoves

	whiteAttacks := pos.Bd.AttackedSquares(game.White)
	blackAttacks := pos.Bd.AttackedSquares(game.Black)

	WhiteCentreControl := centreControl(whiteAttacks)
	BlackCentreControl := centreControl(blackAttacks)

	centreControlScore := float32(WhiteCentreControl - BlackCentreControl)
	if whiteMaterialScore+blackMaterialScore <= 34 {
		return materialScore + 0.2*legalMoveScore
	}
	return materialScore + 0.15*legalMoveScore + 0.1*centreControlScore
}

// NewMaterialWeighter is a helper function to calculate the material weights
func NewMaterialWeighter(m map[game.PieceType]int) float32 {
	weights := map[game.PieceType]float32{
		game.Pawn:   1,
		game.Knight: 2.5,
		game.Rook:   4,
		game.Queen:  6,
		game.King:   0,
	}
	var total float32
	for _, pt := range game.AllPieceTypes {
		total += float32(m[pt]) * weights[pt]
	}
	return total
}

// centreControl
func centreControl(b game.BitBoard) int8 {
	var cnt int8
	for _, sq := range []game.Square{game.C3, game.C4, game.D3, game.D4} {
		if b.Occupied(sq) {
			cnt++
		}
	}
	return cnt
}
