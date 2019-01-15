package evaluators

import "github.com/an1jay/los-alamos-chess/game"

// FirstEvaluator evaluates a position
type FirstEvaluator struct{}

// Evaluate evaluates a position and returns a score depending on how favourable the position is
func (b FirstEvaluator) Evaluate(pos *game.Position) float32 {
	// If game over, return evaluation of the game over state
	res := pos.Result()
	if res != game.InPlay {
		return res.Evaluation()
	}

	WhiteMaterial := pos.Bd.MaterialCount(game.White)
	BlackMaterial := pos.Bd.MaterialCount(game.Black)

	materialScore := StdMaterialWeighter(WhiteMaterial) - StdMaterialWeighter(BlackMaterial)

	return materialScore
}

// StdMaterialWeighter is a helper function to calculate the material weights
func StdMaterialWeighter(m map[game.PieceType]int) float32 {
	weights := map[game.PieceType]float32{
		game.Pawn:   1,
		game.Knight: 3,
		game.Rook:   5,
		game.Queen:  9,
		game.King:   0,
	}
	var total float32
	for _, pt := range game.AllPieceTypes {
		total += float32(m[pt]) * weights[pt]
	}
	return total
}
