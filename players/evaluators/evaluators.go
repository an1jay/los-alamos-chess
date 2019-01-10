package evaluators

import "github.com/an1jay/los-alamos-chess/game"

// Evaluator is a type of evaluator
type Evaluator interface {
	Evaluate(*game.Position) float32
}
