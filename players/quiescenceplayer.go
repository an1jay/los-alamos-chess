package players

import (
	"fmt"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
	"github.com/an1jay/los-alamos-chess/players/evaluators"
)

// QuiescencePlayer plays according to the minimax algorithm and alpha-beta pruning
type QuiescencePlayer struct {
	Ev        evaluators.Evaluator
	MaxDepth  uint
	NodeCount uint
}

// ChooseMove returns a move chosen by Minimax
func (m *QuiescencePlayer) ChooseMove(pos *game.Position) *game.Ply {
	fmt.Println("Quiescence Player thinks ...")
	m.NodeCount = 0
	t0 := time.Now()
	bestScore := DefaultVal * game.NewResultWin(pos.Turn.Other()).EvaluationCoefficient() // -1000 for white; 1000 for black
	var bestMove *game.Ply
	legalMoves := pos.GenerateLegalMoves()

	alpha := -DefaultVal
	beta := DefaultVal

	switch pos.Turn {
	case game.White:
		for _, lgm := range legalMoves {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			scr := m.MinimaxAB(m.MaxDepth, game.Black, alpha, beta, newPos)
			if scr == game.NewResultWin(game.White).Evaluation() {
				return lgm
			}
			if scr >= bestScore {
				bestScore = scr
				bestMove = lgm
			}
		}
	case game.Black:
		for _, lgm := range legalMoves {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			scr := m.MinimaxAB(m.MaxDepth, game.White, alpha, beta, newPos)
			if scr == game.NewResultWin(game.Black).Evaluation() {
				return lgm
			}
			if scr <= bestScore {
				bestScore = scr
				bestMove = lgm
			}
		}
	}

	dt := time.Since(t0).Seconds()
	fmt.Printf("Quiescence Player explored %d nodes in %.02f seconds at %.02f nodes/s \n", m.NodeCount, dt, float64(m.NodeCount)/dt)
	fmt.Printf("Best Score was %f\n", bestScore)
	return bestMove
}

// MinimaxAB calculates the minimax value for a position
func (m *QuiescencePlayer) MinimaxAB(depth uint, side game.Color, alpha, beta float32, pos *game.Position) float32 {
	m.NodeCount++

	// if at a terminal node, evaluate:
	res := pos.Result()
	if res != game.InPlay || depth == 0 {
		ev := m.Ev.Evaluate(pos)
		// pos.Display(true)
		// fmt.Printf("Evaluation: %.10f\n\n", ev)
		return ev
	}
	value := DefaultVal

	switch side {
	case game.White:
		value *= -1
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			if lgm.Capture || lgm.Promotion != game.NoPieceType {
				value = max(value, m.MinimaxAB(umax(depth-1, 1), game.Black, alpha, beta, newPos))
			} else {
				value = max(value, m.MinimaxAB(depth-1, game.Black, alpha, beta, newPos))
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
				value = min(value, m.MinimaxAB(umax(depth-1, 1), game.White, alpha, beta, newPos))
			} else {
				value = min(value, m.MinimaxAB(depth-1, game.White, alpha, beta, newPos))
			}
			beta = min(beta, value)
			if alpha >= beta {
				break
			}
		}
	}

	return value
}

func umax(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}
