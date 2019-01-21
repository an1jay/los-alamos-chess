package players

import (
	"fmt"
	"sync"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
)

// Neo is a faster AI
type Neo struct {
	MinDepth        uint
	MaxDepth        uint
	Ev              *Evaluator
	positionQueue   chan moveAndPosition
	evaluationQueue chan evaluation
	wg              *sync.WaitGroup
}

// CreateNewNeo returns a new Neo
func CreateNewNeo(minDepth uint, maxDepth uint, ev *Evaluator, threadCount int) *Neo {
	var waitg sync.WaitGroup

	NN := &Neo{
		MinDepth:        minDepth,
		MaxDepth:        maxDepth,
		Ev:              ev,
		positionQueue:   make(chan moveAndPosition, 100),
		evaluationQueue: make(chan evaluation, 100),
		wg:              &waitg,
	}

	for i := 0; i < threadCount; i++ {
		go positionSearcher(NN.positionQueue, NN.evaluationQueue, NN.wg, NN.MinDepth, NN.MaxDepth, NN.Ev)
	}

	return NN
}

// ChooseMove asks Neo to choose a move
func (n *Neo) ChooseMove(pos *game.Position) *game.Ply {
	t0 := time.Now()
	fmt.Println("Neo thinks...")

	legalMoves := pos.GenerateLegalMoves()
	numLegalMoves := len(legalMoves)

	for _, lgm := range legalMoves {
		n.wg.Add(1)
		newPos := pos.Copy()
		newPos.UnsafeMove(lgm)
		n.positionQueue <- moveAndPosition{
			move: *lgm,
			pos:  newPos,
		}
	}

	n.wg.Wait()

	var reorderLegalMoves = make([]game.Ply, numLegalMoves)
	var reorderMoveScores = make([]float32, numLegalMoves)
	var nodecnt uint
	var counter int

	// for e := range n.evaluationQueue {
	// 	reorderLegalMoves[counter] = e.move
	// 	reorderMoveScores[counter] = e.eval
	// 	nodecnt += e.nodecnt
	// }

	var ok = true
	var item evaluation
	for ok {
		select {
		case item = <-n.evaluationQueue:
			reorderLegalMoves[counter] = item.move
			reorderMoveScores[counter] = item.eval
			nodecnt += item.nodecnt
			counter++
			ok = true
		default:
			ok = false
		}
	}

	var bestScore = float32(pos.Turn.Other().Coefficient()) * DefaultVal
	var bestMove *game.Ply

	for i := 0; i < numLegalMoves; i++ {
		if sideGeqLeq(pos.Turn, reorderMoveScores[i], bestScore) {
			bestScore = reorderMoveScores[i]
			bestMove = &reorderLegalMoves[i]
		}
	}

	dt := time.Since(t0).Seconds()
	fmt.Printf("Neo explored %d nodes in %.02f seconds at %.02f nodes/s \n", nodecnt, dt, float64(nodecnt)/dt)
	return bestMove
}

func positionSearcher(in chan moveAndPosition, out chan evaluation, wg *sync.WaitGroup, minDepth, maxDepth uint, ev *Evaluator) {
	for candidateNode := range in {
		val, ndct := MinimaxAlphaBetaQuiescenceConcurrently(minDepth, maxDepth, 0, candidateNode.pos.Turn, candidateNode.pos, ev, -1*DefaultVal, DefaultVal)
		message := evaluation{
			move:    candidateNode.move,
			nodecnt: ndct,
			eval:    val,
		}

		fmt.Println("message: ", message.String())
		out <- message

		wg.Done()
	}
}
