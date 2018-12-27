package players

import (
	"fmt"

	"github.com/an1jay/los-alamos-chess/game"
)

// HumanPlayer is a construct to allow a human player to play
type HumanPlayer struct{}

// ChooseMove asks for a legal move through stdout/stdin
func (h *HumanPlayer) ChooseMove(pos *game.Position) *game.Ply {
	fmt.Println("Legal Moves: ", pos.GenerateLegalMoves())
	p := &game.Ply{}
	for true {
		fmt.Println("Enter Move (e.g. 'e5xd6=Knight' denotes a capture from e5 to e6, with promotion to Knight): ")
		var move string
		_, err := fmt.Scan(&move)
		if err != nil {
			fmt.Println("Try Again")
			continue
		}
		p.SourceSq = game.StringToSquareMap[move[0:2]]
		p.DestinationSq = game.StringToSquareMap[move[3:5]]
		if move[2] == 'x' {
			p.Capture = true
		}
		p.Side = pos.Turn
		p.Promotion = game.PieceTypeFromString(move[6:])

		if pos.LegalMove(p) {
			break
		}
		fmt.Println("Try Again")
	}
	return p
}
