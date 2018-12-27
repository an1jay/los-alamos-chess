package main

import (
	"fmt"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
	"github.com/an1jay/los-alamos-chess/players/evaluators"
)

// Player defines a struct with a Move method, which can play Los-Alamos-Chess.
type Player interface {
	ChooseMove(*game.Position) *game.Ply
}

// Game is a game of Los Alamos Chess
type Game struct {
	moveHistory []*game.Ply
	evList      []float32
}

// Play plays a game of Los Alamos Chess returing a game.Result
func (g Game) Play(white, black Player, verbose bool) game.Result {
	// make move list
	g.moveHistory = []*game.Ply{}

	// make new game position
	pos := game.NewGamePosition()

	// fmt.Println("Legal Move Check")

	// main game loop
	for true {
		time.Sleep(100 * time.Millisecond)

		if verbose {
			// fmt.Println("")
			pos.Display(true)
		}

		// depending on whose move, get move
		if pos.Turn == game.White {
			mW := white.ChooseMove(pos.Copy())
			g.moveHistory = append(g.moveHistory, mW)
			if !pos.LegalMove(mW) {
				fmt.Printf("White plays illegal move - %s\n\n", mW.String())
				return game.BlackWin
			}
			pos.Move(mW)
			if verbose {
				fmt.Printf("White plays %s\n\n", mW.String())
			}
		} else if pos.Turn == game.Black {
			mB := black.ChooseMove(pos.Copy())
			g.moveHistory = append(g.moveHistory, mB)
			if !pos.LegalMove(mB) {
				fmt.Printf("Black plays illegal move - %s\n\n", mB.String())
				return game.WhiteWin
			}
			pos.Move(mB)
			if verbose {
				fmt.Printf("Black plays %s\n\n", mB.String())
			}
		}

		// if pos.HalfMoveClock == 4 {
		// 	panic("Test")
		// }

		if verbose {
			posEv := evaluators.SecondEvaluator{}.Evaluate(pos)
			g.evList = append(g.evList, posEv)
			fmt.Printf("Evaluation: %v\n", g.evList)
		}

		// check game over
		res := pos.Result()
		if res != game.InPlay {
			if verbose {
				fmt.Println("")
				fmt.Println("Game Over! \nFinal Position")
				fmt.Println("Move History: ", g.moveHistory)
				fmt.Println(res)
				fmt.Printf("Full Moves: %d\nHalf-Moves: %d\n", pos.MoveNumber, pos.HalfMoveClock)
				pos.Display(true)
			}
			return res
		}

	}
	panic("Game somehow not over")
}
