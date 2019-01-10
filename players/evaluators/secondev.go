package evaluators

import "github.com/an1jay/los-alamos-chess/game"

// SecondEvaluator evaluates a position
type SecondEvaluator struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (b SecondEvaluator) Evaluate(pos *game.Position) float32 {
	// If game over, return evaluation of the game over state
	res := pos.Result()
	if res != game.InPlay {
		return res.Evaluation()
	}

	WhiteMaterial := pos.Bd.MaterialCount(game.White)
	BlackMaterial := pos.Bd.MaterialCount(game.Black)

	materialScore := MaterialWeighter(WhiteMaterial) - MaterialWeighter(BlackMaterial)

	backUpTurn := pos.Turn

	pos.Turn = game.White
	WhiteLegalMoves := float32(pos.GenerateCountOfPseudoLegalMoves())

	pos.Turn = game.Black
	BlackLegalMoves := float32(pos.GenerateCountOfPseudoLegalMoves())

	pos.Turn = backUpTurn

	legalMoveScore := WhiteLegalMoves - BlackLegalMoves

	return materialScore + 0.05*legalMoveScore
}
