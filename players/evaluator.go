package players

import "github.com/an1jay/los-alamos-chess/game"

// Evaluator does all evaluations
type Evaluator struct {
	MaterialCoeff      float32
	LegalMovesCoeff    float32
	SquareControlCoeff float32
	MaterialWeights    map[game.PieceType]float32
	WhiteSquareWeights map[game.Square]float32
	BlackSquareWeights map[game.Square]float32
}

// Evaluate returns an evaluation based on coefficients and weights
func (ev Evaluator) Evaluate(pos *game.Position) float32 {
	// If game over, return evaluation of the game over state
	res := pos.Result()
	if res != game.InPlay {
		return res.Evaluation()
	}

	var score float32

	if ev.MaterialCoeff != 0 {
		score += ev.MaterialCoeff * materialEvaluation(pos, ev.MaterialWeights)
	}

	if ev.LegalMovesCoeff != 0 {
		score += ev.LegalMovesCoeff * legalMovesEvaluation(pos)
	}

	if ev.SquareControlCoeff != 0 {
		score += ev.SquareControlCoeff * squareControlEvaluation(pos, ev.WhiteSquareWeights, ev.BlackSquareWeights)
	}
	return score
}
