package players

import "github.com/an1jay/los-alamos-chess/game"

// HumanPlayer is a construct to allow a human player to play
type HumanPlayer struct{}

// Move asks for a move through stdout/stdin
func (h HumanPlayer) Move(pos *game.Position) *game.Ply {

}
