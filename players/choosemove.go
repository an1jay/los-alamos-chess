package players

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
	// e "github.com/derekparker/delve/service/test"
)

// ChooseRandomMove asks a random player to move randomly from choice of legal moves.
func ChooseRandomMove(pos *game.Position) *game.Ply {
	rand.Seed(time.Now().UnixNano())
	lm := pos.GenerateLegalMoves()
	return lm[rand.Intn(len(lm))]
}

// ChooseMinimaxMove returns a slice of the best moves (according to Minimax with the specified Evaluator)
func ChooseMinimaxMove(pos *game.Position, ev *Evaluator, maxDepth uint) ([]*game.Ply, uint) {
	var nodeCount uint
	var bestMoves []*game.Ply
	var bestScore = DefaultVal * float32(pos.Turn.Other().Coefficient()) // -1000 for pos.Turn = white; 1000 for pos.Turn = black
	var legalMoves = pos.GenerateLegalMoves()
	var scoreList = make([]float32, len(legalMoves))

	for num, lgm := range legalMoves {
		newPos := pos.Copy()
		newPos.UnsafeMove(lgm)
		scr := Minimax(maxDepth, pos.Turn.Other(), newPos, &nodeCount, ev)
		scoreList[num] = scr
		if sideGeqLeq(pos.Turn, scr, bestScore) {
			bestScore = scr
		}
	}

	for moveNum, score := range scoreList {
		if score == bestScore {
			bestMoves = append(bestMoves, legalMoves[moveNum])
		}
	}
	return bestMoves, nodeCount
}

// ChooseMinimaxAlphaBetaMove returns a slice of the best moves (according to Minimax with the specified Evaluator)
func ChooseMinimaxAlphaBetaMove(pos *game.Position, ev *Evaluator, maxDepth uint, alpha, beta float32) ([]*game.Ply, uint) {
	var nodeCount uint
	var bestMoves []*game.Ply
	var bestScore = DefaultVal * float32(pos.Turn.Other().Coefficient()) // -1000 for pos.Turn = white; 1000 for pos.Turn = black
	var legalMoves = pos.GenerateLegalMoves()
	var scoreList = make([]float32, len(legalMoves))

	for num, lgm := range legalMoves {
		newPos := pos.Copy()
		newPos.UnsafeMove(lgm)
		scr := MinimaxAlphaBeta(maxDepth, pos.Turn.Other(), newPos, &nodeCount, ev, alpha, beta)
		scoreList[num] = scr
		if sideGeqLeq(pos.Turn, scr, bestScore) {
			bestScore = scr
		}
	}

	for moveNum, score := range scoreList {
		if score == bestScore {
			bestMoves = append(bestMoves, legalMoves[moveNum])
		}
	}
	return bestMoves, nodeCount
}

// ChooseMinimaxAlphaBetaQuiescence returns a slice of the best moves (according to Minimax with the specified Evaluator)
func ChooseMinimaxAlphaBetaQuiescence(pos *game.Position, ev *Evaluator, minDepth, maxDepth uint, alpha, beta float32) ([]*game.Ply, uint) {
	var nodeCount uint
	var bestMoves []*game.Ply
	var bestScore = DefaultVal * float32(pos.Turn.Other().Coefficient()) // -1000 for pos.Turn = white; 1000 for pos.Turn = black
	var legalMoves = pos.GenerateLegalMoves()
	var scoreList = make([]float32, len(legalMoves))
	fmt.Println("Legal Moves", legalMoves)

	for num, lgm := range legalMoves {
		newPos := pos.Copy()
		newPos.UnsafeMove(lgm)
		scr := MinimaxAlphaBetaQuiescence(minDepth, maxDepth, 0, pos.Turn.Other(), newPos, &nodeCount, ev, alpha, beta)
		// fmt.Println("Score in Lopp: ", scr)
		scoreList[num] = scr
		if sideGeqLeq(pos.Turn, scr, bestScore) {
			bestScore = scr
		}
	}

	for moveNum, score := range scoreList {
		if score == bestScore {
			bestMoves = append(bestMoves, legalMoves[moveNum])
		}
	}
	fmt.Println("Score List: ", scoreList)
	return bestMoves, nodeCount
}

// ChooseMinimaxAlphaBetaQuiescenceConcurrently returns a slice of the best moves (according to Minimax with the specified Evaluator)
func ChooseMinimaxAlphaBetaQuiescenceConcurrently(pos *game.Position, ev *Evaluator, minDepth, maxDepth uint, alpha, beta float32) ([]*game.Ply, uint) {
	// type empty {}
	// ...
	// data := make([]float, N);
	// res := make([]float, N);
	// sem := make(chan empty, N);  // semaphore pattern
	// ...
	// for i,xi := range data {
	// 	go func (i int, xi float) {
	// 		res[i] = doSomething(i,xi);
	// 		sem <- empty{};
	// 	} (i, xi);
	// }
	// // wait for goroutines to finish
	// for i := 0; i < N; ++i { <-sem }

	var bestMoves []*game.Ply
	var bestScore = DefaultVal * float32(pos.Turn.Other().Coefficient()) // -1000 for pos.Turn = white; 1000 for pos.Turn = black

	var legalMoves = pos.GenerateLegalMoves()
	var numMoves = len(legalMoves)

	// var scoreList = make([]float32, numMoves)
	// var nodeCount = make([]uint, numMoves)
	var nctotal uint
	var sem = make(chan evaluation, numMoves)

	fmt.Println("Legal Moves", legalMoves)
	var wg sync.WaitGroup

	wg.Add(len(legalMoves))

	for num, lgm := range legalMoves {
		newPos := pos.Copy()
		newPos.UnsafeMove(lgm)

		// wg *sync.WaitGroup,
		// fmt.Println("in for loop")
		go asyncSearch(lgm, &wg, &sem, num, minDepth, maxDepth, 0, pos.Turn.Other(), newPos, ev, alpha, beta)
	}
	// &wg,

	wg.Wait()
	var evList []*evaluation
	// var evList = make([]*evaluation, numMoves)

	close(sem)
	for eval := range sem {
		evList = append(evList, &eval)
	}
	fmt.Println("evList ", evList)
	// scr := MinimaxAlphaBetaQuiescenceConcurrently(minDepth, maxDepth, 0, pos.Turn.Other(), newPos, &nodeCount, ev, alpha, beta)
	// fmt.Println("Score in Lopp: ", scr)
	// scoreList[num] = scr
	for _, ev1 := range evList {
		if sideGeqLeq(pos.Turn, ev1.evalua, bestScore) {
			bestScore = ev1.evalua
		}
		nctotal += ev1.nodecnt
	}

	for _, ev2 := range evList {
		if ev2.evalua == bestScore {
			bestMoves = append(bestMoves, ev2.move)
		}
	}
	// fmt.Println("Score List: ", scoreList)
	fmt.Println("Node Count: ", nctotal)
	return bestMoves, nctotal
}

type evaluation struct {
	move    *game.Ply
	nodecnt uint
	evalua  float32
}

func (ev evaluation) String() string {
	return fmt.Sprintf("Move: %s, NodeCount: %d, eval: %.2f", (*ev.move).String(), ev.nodecnt, ev.evalua)
}

func asyncSearch(prevMove *game.Ply, wg *sync.WaitGroup, sem *chan evaluation, no int, depth, maxDepth, depthCount uint, side game.Color, pos *game.Position, evaluator *Evaluator, alpha, beta float32) {
	// fmt.Println("in func")
	mv, eval, nc := MinimaxAlphaBetaQuiescenceConcurrently(prevMove, depth, maxDepth, depthCount, side, pos, evaluator, alpha, beta)
	*sem <- evaluation{
		move:    mv,
		nodecnt: nc,
		evalua:  eval,
	}
	// fmt.Println("scoreList", *sclist)
	wg.Done()
}
