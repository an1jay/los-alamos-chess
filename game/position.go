package game

import (
	"fmt"
)

// Position stores the entire gamestate at a point in time.
type Position struct {
	Bd            Board
	Turn          Color
	MoveNumber    uint
	HalfMoveClock uint
	InCheck       bool
}

// Move updates the position with the given ply, returning false if illegal or invalid move (does not make the move).
func (pos *Position) Move(p *Ply, c Color) bool {
	if c != pos.Turn {
		fmt.Println("Position.Move: Wrong side tries to move")
		return false
	}

	// moveValid := pos.LegalMove(p)
	if pos.Turn == Black {
		pos.MoveNumber++
	}
	pos.HalfMoveClock++
	pos.Turn = c.Other()
	// pos.inCheck := pos.brd.InCheck(pos.Turn)

	return false
}

// LegalMove returns whether a move is legal.
func (pos *Position) LegalMove(p *Ply) bool {
	// Check side to move
	if p.Side != pos.Turn {
		return false
	}
	// Check destination square is empty if capture is true
	if p.Capture {
		return pos.Bd.Piece(p.DestinationSq) == NoPiece
	}
	// Check if the piece can actually move there
	piece := pos.Bd.Piece(p.SourceSq)

	return true
}

func (pos *Position)