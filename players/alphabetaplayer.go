package players

import (
	"fmt"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
	"github.com/an1jay/los-alamos-chess/players/evaluators"
)

// AlphaBetaPlayer plays according to the minimax algorithm and alpha-beta pruning
type AlphaBetaPlayer struct {
	Ev        evaluators.Evaluator
	MaxDepth  uint
	NodeCount uint
}

// ChooseMove returns a move chosen by Minimax
func (m *AlphaBetaPlayer) ChooseMove(pos *game.Position) *game.Ply {
	fmt.Println("Alpha Beta Player thinks ...")
	m.NodeCount = 0
	t0 := time.Now()
	bestScore := DefaultVal * game.NewResultWin(pos.Turn.Other()).EvaluationCoefficient() // * 10   // -100 for white; 100 for black
	// fmt.Println("Best Score: ", bestScore)
	var bestMove *game.Ply
	legalMoves := pos.GenerateLegalMoves()
	// fmt.Println("Legal Moves: \n", legalMoves)

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
				// fmt.Println("bestMove", bestMove)
				bestMove = lgm
			}
		}
	}

	dt := time.Since(t0).Seconds()
	fmt.Printf("Minimax Player explored %d nodes in %.02f seconds at %.02f nodes/s \n", m.NodeCount, dt, float64(m.NodeCount)/dt)
	fmt.Println(pos.GenerateLegalMoves())
	fmt.Printf("Best Score was %f\n", bestScore)
	return bestMove
}

// MinimaxAB calculates the minimax value for a position
func (m *AlphaBetaPlayer) MinimaxAB(depth uint, side game.Color, alpha, beta float32, pos *game.Position) float32 {
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
			value = max(value, m.MinimaxAB(depth-1, game.Black, alpha, beta, newPos))
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
			value = min(value, m.MinimaxAB(depth-1, game.White, alpha, beta, newPos))
			beta = min(beta, value)
			if alpha >= beta {
				break
			}
		}
	}

	return value
}
