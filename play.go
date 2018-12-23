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
}

// Play plays a game of Los Alamos Chess returing a game.Result
func (g Game) Play(white, black Player, verbose bool) game.Result {
	// make move list
	g.moveHistory = []*game.Ply{}

	// make new game position
	pos := game.NewGamePosition()

	// main game loop
	for true {
		time.Sleep(100 * time.Millisecond)

		// depending on whose move, get move
		if pos.Turn == game.White {
			if verbose {
				// fmt.Println("")
				pos.Display(false)
			}
			mW := white.ChooseMove(pos)
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
			if verbose {
				fmt.Println("")
				pos.Display(false)
			}
			mB := black.ChooseMove(pos)
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

		// check game over
		res := pos.Result()
		if res != game.InPlay {
			if verbose {
				fmt.Println("")
				fmt.Println("Game Over! \nFinal Position")
				fmt.Println("Move History: ", g.moveHistory)
				fmt.Println(res)
				pos.Display(false)
			}
			return res
		}

	}
	panic("Game somehow not over")
}
