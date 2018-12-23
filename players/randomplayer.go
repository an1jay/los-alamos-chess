package players

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/an1jay/los-alamos-chess/game"
)

// RandomPlayer plays Los Alamos Chess randomly
type RandomPlayer struct{}

// ChooseMove asks a random player to move randomly from choice of legal moves.
func (r RandomPlayer) ChooseMove(pos *game.Position) *game.Ply {
	rand.Seed(time.Now().UnixNano())
	lm := pos.GenerateLegalMoves()
	fmt.Println("Legal Moves: ", lm)
	return lm[rand.Intn(len(lm))]
}
