package players

import "github.com/an1jay/los-alamos-chess/game"

// materialEvaluation evaluates a position and returns a score depending material amounts
func materialEvaluation(pos *game.Position, MaterialWeights map[game.PieceType]float32) float32 {
	// If game over, return evaluation of the game over state

	WhiteMaterial := pos.Bd.MaterialCount(game.White)
	BlackMaterial := pos.Bd.MaterialCount(game.Black)

	materialScore := materialWeighter(WhiteMaterial, MaterialWeights) - materialWeighter(BlackMaterial, MaterialWeights)

	return materialScore
}

// materialWeighter calculates material score based on weights
func materialWeighter(m map[game.PieceType]int, weights map[game.PieceType]float32) float32 {

	var total float32
	for _, pt := range game.AllPieceTypes {
		total += float32(m[pt]) * weights[pt]
	}
	return total
}

// legalMovesEvaluation evaluates based on number of available legal moves
func legalMovesEvaluation(pos *game.Position) float32 {

	backUpTurn := pos.Turn

	pos.Turn = game.White
	WhiteLegalMoves := float32(pos.GenerateCountOfPseudoLegalMoves())

	pos.Turn = game.Black
	BlackLegalMoves := float32(pos.GenerateCountOfPseudoLegalMoves())

	pos.Turn = backUpTurn

	legalMoveScore := WhiteLegalMoves - BlackLegalMoves

	return legalMoveScore
}

// squareControlEvaluation takes in a map of square weights for each of black and White
// It returns a float32 based on the weighted control that each player has on each square
func squareControlEvaluation(pos *game.Position, whtsqrWeights map[game.Square]float32, blksqrWeights map[game.Square]float32) float32 {
	return squareWeighter(pos.Bd.NumAttacksPerSquare(game.White), whtsqrWeights) -
		squareWeighter(pos.Bd.NumAttacksPerSquare(game.Black), blksqrWeights)
}

// squareWeighter calculates a score based on a weight for each square and the number of attacks on that square
func squareWeighter(squareCount map[game.Square]int8, squareWeights map[game.Square]float32) float32 {
	var score float32
	for sq, wt := range squareWeights {
		score += float32(squareCount[sq]) * wt
	}
	return score
}
