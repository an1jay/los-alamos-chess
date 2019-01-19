package main

import (
	"fmt"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
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

// PlayFromPos plays a game of Los Alamos Chess from given position returing a game.Result
func (g Game) PlayFromPos(white, black Player, verbose bool, posToPlayFrom *game.Position) game.Result {
	// make move list
	g.moveHistory = []*game.Ply{}

	// make new game position
	pos := posToPlayFrom

	// fmt.Println("Legal Move Check")

	// main game loop
	for true {
		time.Sleep(100 * time.Millisecond)

		if verbose {
			// fmt.Println("")
			pos.Display(false)
		}

		// depending on whose move, get move
		if pos.Turn == game.White {
			mW := white.ChooseMove(pos.Copy())
			g.moveHistory = append(g.moveHistory, mW)
			if !pos.LegalPly(mW) {
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
			if !pos.LegalPly(mB) {
				fmt.Printf("Black plays illegal move - %s\n\n", mB.String())
				return game.WhiteWin
			}
			pos.Move(mB)
			if verbose {
				fmt.Printf("Black plays %s\n\n", mB.String())
			}
		}

		// if pos.HalfMoveClock == 8 {
		// 	panic("Test")
		// }

		// check game over
		res := pos.Result()
		if res != game.InPlay {
			if verbose {
				fmt.Println("")
				fmt.Println("Game Over! \nFinal Position")
				fmt.Println("Move History: ", g.moveHistory)
				fmt.Println(res)
				fmt.Printf("Full Moves: %d\nHalf-Moves: %d\n", pos.MoveNumber, pos.HalfMoveClock)
				pos.Display(false)
			}
			return res
		}

	}
	panic("Game somehow not over")
}
