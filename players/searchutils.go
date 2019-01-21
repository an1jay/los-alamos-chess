package players

import (
	"fmt"

	"github.com/an1jay/los-alamos-chess/game"
)

// DefaultVal is larger than maximum possible evaluation
const DefaultVal float32 = 1000

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

func sideMinMax(side game.Color, x, y float32) float32 {
	if side == game.White {
		return max(x, y)
	}
	return min(x, y)
}

func sideGeqLeq(side game.Color, x, y float32) bool {
	if side == game.White {
		return x >= y
	}
	return x <= y
}

func umax(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func sumUintSlice(arr []uint) uint {
	var total uint
	for _, num := range arr {
		total += num
	}
	return total
}

type evaluation struct {
	move    *game.Ply
	nodecnt uint
	eval    float32
}

type moveAndPosition struct {
	pos  *game.Position
	move *game.Ply
}

func (ev evaluation) String() string {
	return fmt.Sprintf("Move: %s, NodeCount: %d, eval: %.2f", (*ev.move).String(), ev.nodecnt, ev.eval)
}
