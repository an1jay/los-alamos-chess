package players

import (
	"fmt"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
	"github.com/an1jay/los-alamos-chess/players/evaluators"
)

// MinimaxPlayer plays according to the minimax algorithm
type MinimaxPlayer struct {
	Ev        evaluators.Evaluator
	MaxDepth  uint
	NodeCount uint
}

// ChooseMove returns a move chosen by Minimax
func (m *MinimaxPlayer) ChooseMove(pos *game.Position) *game.Ply {
	fmt.Println("Minimax Player thinks ...")
	m.NodeCount = 0
	t0 := time.Now()
	bestScore := DefaultVal * game.NewResultWin(pos.Turn.Other()).EvaluationCoefficient() // * 10   // -100 for white; 100 for black
	// fmt.Println("Best Score: ", bestScore)
	var bestMove *game.Ply
	switch pos.Turn {
	case game.White:
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			scr := m.Minimax(m.MaxDepth, game.White, newPos)
			if scr == game.NewResultWin(game.White).Evaluation() {
				return lgm
			}
			if scr >= bestScore {
				bestScore = scr
				bestMove = lgm
			}
		}
	case game.Black:
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			scr := m.Minimax(m.MaxDepth, game.Black, newPos)
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

// DefaultVal is larger than maximum possible evaluation
const DefaultVal float32 = 1000

// Minimax calculates the minimax value for a position
func (m *MinimaxPlayer) Minimax(depth uint, side game.Color, pos *game.Position) float32 {
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
			value = max(value, m.Minimax(depth-1, game.Black, newPos))
		}
	case game.Black:
		value *= 1
		for _, lgm := range pos.GenerateLegalMoves() {
			newPos := pos.Copy()
			newPos.UnsafeMove(lgm)
			value = min(value, m.Minimax(depth-1, game.White, newPos))
		}
	}

	return value
}

func max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

func min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}
