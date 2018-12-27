package evaluators

import "github.com/an1jay/los-alamos-chess/game"

type Evaluator interface {
	Evaluate(*game.Position) float32
}
