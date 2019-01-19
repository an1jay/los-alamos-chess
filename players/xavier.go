package players

import (
	"fmt"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
)

// Xavier is an AI
type Xavier struct {
	MinDepth uint
	MaxDepth uint
	Ev       *Evaluator
}

// ChooseMove asks Xavier to choose a move
func (x *Xavier) ChooseMove(pos *game.Position) *game.Ply {
	t0 := time.Now()
	fmt.Println("Xavier thinks...")
	mv, nodecnt := ChooseMinimaxAlphaBetaQuiescenceConcurrently(pos, x.Ev, x.MinDepth, x.MaxDepth, -1*DefaultVal, DefaultVal)
	dt := time.Since(t0).Seconds()
	fmt.Printf("Xavier explored %d nodes in %.02f seconds at %.02f nodes/s \n", nodecnt, dt, float64(nodecnt)/dt)
	return mv[0]
}
