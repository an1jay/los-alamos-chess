package players

import (
	"github.com/an1jay/los-alamos-chess/game"
)

// Minimax calculates the minimax value for a position
func Minimax(depth uint, side game.Color, pos *game.Position, NodeCount *uint, evaluator *Evaluator) float32 {
	(*NodeCount)++
	// if at a terminal node, evaluate:
	res := pos.Result()
	if res != game.InPlay || depth == 0 {
		ev := evaluator.Evaluate(pos)
		return ev
	}

	// else minimax:
	value := DefaultVal * float32(side.Other().Coefficient())
	for _, lgm := range pos.GenerateLegalMoves() {
		newPos := pos.Copy()
		newPos.UnsafeMove(lgm)
		value = sideMinMax(side, value, Minimax(depth-1, side.Other(), newPos, NodeCount, evaluator))
	}
	return value
}

// MinimaxAlphaBeta calculates the minimax value for a position
func MinimaxAlphaBeta(depth uint, side game.Color, pos *game.Position, NodeCount *uint, evaluator *Evaluator, alpha, beta float32) float32 {
	(*NodeCount)++
	// if at a terminal node, evaluate:
	res := pos.Result()
	if res != game.InPlay || depth == 0 {
		ev := evaluator.Evaluate(pos)
		return ev
	}
	value := DefaultVal * float32(pos.Turn.Other().Coefficient()) // -1000 for pos.Turn = white; 1000 for pos.Turn = black

	switch side {
	case game.White:
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			value = max(value, MinimaxAlphaBeta(depth-1, game.Black, newPos, NodeCount, evaluator, alpha, beta))
			alpha = max(alpha, value)
			if alpha >= beta {
				break
			}
		}
	case game.Black:
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			value = min(value, MinimaxAlphaBeta(depth-1, game.White, newPos, NodeCount, evaluator, alpha, beta))
			beta = min(beta, value)
			if alpha >= beta {
				break
			}
		}
	}
	return value
}

// MinimaxAlphaBetaQuiescence calculates the minimax value for a position
func MinimaxAlphaBetaQuiescence(depth, maxDepth, depthCount uint, side game.Color, pos *game.Position, NodeCount *uint, evaluator *Evaluator, alpha, beta float32) float32 {
	(*NodeCount)++
	// if at a terminal node, evaluate:
	res := pos.Result()
	if res != game.InPlay || depth == 0 {
		ev := evaluator.Evaluate(pos)
		return ev
	}
	value := DefaultVal

	switch side {
	case game.White:
		value *= -1
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			if (lgm.Capture || lgm.Promotion != game.NoPieceType) && depthCount < maxDepth {
				value = max(value, MinimaxAlphaBetaQuiescence(umax(depth-1, 1), maxDepth, depthCount+1, game.Black, newPos, NodeCount, evaluator, alpha, beta))
			} else {
				value = max(value, MinimaxAlphaBetaQuiescence(depth-1, maxDepth, depthCount+1, game.Black, newPos, NodeCount, evaluator, alpha, beta))
			}
			alpha = max(alpha, value)
			if alpha >= beta {
				break
			}
		}
	case game.Black:
		value *= 1
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			if lgm.Capture || lgm.Promotion != game.NoPieceType {
				value = min(value, MinimaxAlphaBetaQuiescence(umax(depth-1, 1), maxDepth, depthCount+1, game.White, newPos, NodeCount, evaluator, alpha, beta))
			} else {
				value = min(value, MinimaxAlphaBetaQuiescence(depth-1, maxDepth, depthCount+1, game.White, newPos, NodeCount, evaluator, alpha, beta))
			}
			beta = min(beta, value)
			if alpha >= beta {
				break
			}
		}
	}

	return value
}

// MinimaxAlphaBetaQuiescenceConcurrently calculates the minimax value for a position
func MinimaxAlphaBetaQuiescenceConcurrently(prevMove *game.Ply, depth, maxDepth, depthCount uint, side game.Color, pos *game.Position, evaluator *Evaluator, alpha, beta float32) (*game.Ply, float32, uint) {
	var NodeCount uint
	// if at a terminal node, evaluate:
	res := pos.Result()
	if res != game.InPlay || depth == 0 {
		ev := evaluator.Evaluate(pos)
		NodeCount++
		return prevMove, ev, NodeCount
	}
	value := DefaultVal

	switch side {
	case game.White:
		value *= -1
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			if (lgm.Capture || lgm.Promotion != game.NoPieceType) && depthCount < maxDepth {
				value = max(value, MinimaxAlphaBetaQuiescence(umax(depth-1, 1), maxDepth, depthCount+1, game.Black, newPos, &NodeCount, evaluator, alpha, beta))
			} else {
				value = max(value, MinimaxAlphaBetaQuiescence(depth-1, maxDepth, depthCount+1, game.Black, newPos, &NodeCount, evaluator, alpha, beta))
			}
			alpha = max(alpha, value)
			if alpha >= beta {
				break
			}
		}
	case game.Black:
		value *= 1
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			if lgm.Capture || lgm.Promotion != game.NoPieceType {
				value = min(value, MinimaxAlphaBetaQuiescence(umax(depth-1, 1), maxDepth, depthCount+1, game.White, newPos, &NodeCount, evaluator, alpha, beta))
			} else {
				value = min(value, MinimaxAlphaBetaQuiescence(depth-1, maxDepth, depthCount+1, game.White, newPos, &NodeCount, evaluator, alpha, beta))
			}
			beta = min(beta, value)
			if alpha >= beta {
				break
			}
		}
	}

	return prevMove, value, NodeCount
}
