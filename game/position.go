package game

import (
	"fmt"
	"math/bits"
)

// Position stores the entire gamestate at a point in time.
type Position struct {
	Bd            *Board
	Turn          Color
	MoveNumber    uint
	HalfMoveClock uint
	InCheck       bool
}

// NewPosition constructs a new position.
func NewPosition(bd *Board, turn Color, moveNumber, halfMoveClock uint) *Position {
	return &Position{
		Bd:            bd,
		Turn:          turn,
		MoveNumber:    moveNumber,
		HalfMoveClock: halfMoveClock,
		InCheck:       bd.InCheck(turn),
	}
}

// NewGamePosition returns a position at the start of a game.
func NewGamePosition() *Position {
	bd := BoardFromMap(
		map[Square]Piece{
			A1: WhiteRook,
			B1: WhiteKnight,
			C1: WhiteQueen,
			D1: WhiteKing,
			E1: WhiteKnight,
			F1: WhiteRook,
			A2: WhitePawn,
			B2: WhitePawn,
			C2: WhitePawn,
			D2: WhitePawn,
			E2: WhitePawn,
			F2: WhitePawn,
			A3: NoPiece,
			B3: NoPiece,
			C3: NoPiece,
			D3: NoPiece,
			E3: NoPiece,
			F3: NoPiece,
			A4: NoPiece,
			B4: NoPiece,
			C4: NoPiece,
			D4: NoPiece,
			E4: NoPiece,
			F4: NoPiece,
			A5: BlackPawn,
			B5: BlackPawn,
			C5: BlackPawn,
			D5: BlackPawn,
			E5: BlackPawn,
			F5: BlackPawn,
			A6: BlackRook,
			B6: BlackKnight,
			C6: BlackQueen,
			D6: BlackKing,
			E6: BlackKnight,
			F6: BlackRook,
		})
	return &Position{
		Bd:            bd,
		Turn:          White,
		MoveNumber:    0,
		HalfMoveClock: 0,
		InCheck:       bd.InCheck(White),
	}
}

// Copy returns a pointer to a copy of this position, with a new copy of Board.
func (pos *Position) Copy() *Position {
	return &Position{
		Bd:            pos.Bd.Copy(),
		Turn:          pos.Turn,
		MoveNumber:    pos.MoveNumber,
		HalfMoveClock: pos.HalfMoveClock,
		InCheck:       pos.InCheck,
	}
}

// Display prints a board representation of the position to stdout
func (pos *Position) Display(symb bool) {
	pos.Bd.Display(symb)
	checkStr := "not "
	if pos.InCheck {
		checkStr = ""
	}
	fmt.Printf("%s is to move. %s is %sin check\n", pos.Turn.String(), pos.Turn.String(), checkStr)
}

// Move updates the position with the given ply, returning false if illegal or invalid move (does not make the move).
func (pos *Position) Move(p *Ply) bool {
	if !pos.LegalMove(p) {
		return false
	}
	pos.UnsafeMove(p)
	return true
}

// UnsafeMove updates the position with the given ply, not checking for legality.
func (pos *Position) UnsafeMove(p *Ply) {
	pos.Bd.Move(p.SourceSq, p.DestinationSq, NewPiece(p.Promotion, pos.Turn))
	if pos.Turn == Black {
		pos.MoveNumber++
	}
	pos.HalfMoveClock++
	// fmt.Println(p.Side, "makes a move")
	// fmt.Println(pos.Turn, "is the current Turn")
	pos.Turn = pos.Turn.Other()
	// fmt.Println(pos.Turn, "is the current Turn after turn.Other")
	if pos.Bd.InCheck(pos.Turn) {
		// fmt.Println("WTF")
		pos.InCheck = true
	} else {
		pos.InCheck = false
	}
	// pos.InCheck = pos.Bd.InCheck(pos.Turn)
	// fmt.Println("FInal pos.InCheck: ", pos.InCheck)
	// fmt.Println("FInal pos.InCheck assignment: ", pos.Bd.InCheck(pos.Turn))

	pos.InCheck = false

	if pos.Bd.InCheck(White) && pos.Turn == White {
		pos.InCheck = true
	} else if pos.Bd.InCheck(Black) && pos.Turn == Black {
		pos.InCheck = true
	}
	// fmt.Println("White in Check: ", pos.Bd.InCheck(White))
	// fmt.Println("Black in Check: ", pos.Bd.InCheck(Black))
	// fmt.Println("Final pos.InCheck: ", pos.InCheck)

}

// LegalMove returns whether a move is legal.
func (pos *Position) LegalMove(p *Ply) bool {
	// fmt.Println("LEgalPly: ", p)
	// Check side to move
	if !(p.Side == pos.Turn) {
		return false
	}
	// Check if piece being moved is yours
	if !pos.Bd.PieceOccupancy(p.Side).Occupied(p.SourceSq) {
		return false
	}
	// Check destination square is empty if capture is true
	if !(pos.Bd.Piece(p.DestinationSq) == NoPiece) && (p.Capture == false) {
		return false
	}
	// Check if the piece can actually move there
	if !(pos.Bd.MovesVector(p.SourceSq)&BitBoard(0).SetSquareOnBB(p.DestinationSq, true) != 0) {
		return false
	}
	// Check if move causes the side to move to be in check
	newBoard := pos.Bd.Copy()
	newBoard.Move(p.SourceSq, p.DestinationSq, NewPiece(p.Promotion, p.Side))
	// if the side moving is in check after the move, the move is illegal
	if newBoard.InCheck(p.Side) {
		return false
	}
	return true
}

// GenerateLegalMoves generates a slice of pointers to legal moves.
func (pos *Position) GenerateLegalMoves() []*Ply {
	PseudoLegalPlies := pos.GeneratePseudoLegalMoves()
	var LegalPlies []*Ply
	for _, plp := range PseudoLegalPlies {
		newBd := pos.Bd.Copy()
		newBd.Move(plp.SourceSq, plp.DestinationSq, NewPiece(plp.Promotion, pos.Turn))
		if !newBd.InCheck(pos.Turn) {
			LegalPlies = append(LegalPlies, plp)
		}
	}
	return LegalPlies
}

// GenerateCountOfLegalMoves returns number of legal moves.
func (pos *Position) GenerateCountOfLegalMoves() int {
	PseudoLegalPlies := pos.GeneratePseudoLegalMoves()
	var LegalMoveCount int
	for _, plp := range PseudoLegalPlies {
		newBd := pos.Bd.Copy()
		newBd.Move(plp.SourceSq, plp.DestinationSq, NewPiece(plp.Promotion, pos.Turn))
		if !newBd.InCheck(pos.Turn) {
			LegalMoveCount++
		}
	}
	return LegalMoveCount
}

// GeneratePseudoLegalMoves generates a slice of pointers to pseudolegal moves.
func (pos *Position) GeneratePseudoLegalMoves() []*Ply {
	// update convenience bitboards
	pos.Bd.UpdateConvenienceBBs()

	var PseudoLegalPlies []*Ply
	piece := NoPiece

	// get all possible pieces able to move
	moverPieces := pos.Bd.PieceOccupancy(pos.Turn)
	// get other side's pieces - needed for calculating captures
	otherPieces := pos.Bd.PieceOccupancy(pos.Turn.Other())

	// loop over source squares
	for startSq := 0; startSq < NumSquaresInBoard; startSq++ {
		// if the mover has a piece there
		if moverPieces.Occupied(Square(startSq)) {
			piece = pos.Bd.Piece(Square(startSq))
			// calculate all the pieces moves
			mvVector := pos.Bd.MovesVector(Square(startSq))
			for endSq := 0; endSq < NumSquaresInBoard; endSq++ {
				// for each piece move
				if mvVector.Occupied(Square(endSq)) {
					// check if mover is a pawn and destination square is on either the 1st or 6th ranks and add promotions
					if piece.PieceType() == Pawn && (Square(endSq).Rank() == Rank1 || Square(endSq).Rank() == Rank6) {
						for _, pt := range PromotionPieceTypes {
							PseudoLegalPlies = append(PseudoLegalPlies, &Ply{
								SourceSq:      Square(startSq),
								DestinationSq: Square(endSq),
								Capture:       otherPieces.Occupied(Square(endSq)),
								Promotion:     pt,
								Side:          pos.Turn,
							})
						}
					} else {
						// if not a pawn, add move to pseudolegal plies
						PseudoLegalPlies = append(PseudoLegalPlies, &Ply{
							SourceSq:      Square(startSq),
							DestinationSq: Square(endSq),
							Capture:       otherPieces.Occupied(Square(endSq)),
							Promotion:     NoPieceType,
							Side:          pos.Turn,
						})
					}
				}
			}
		}
	}
	return PseudoLegalPlies
}

// GenerateCountOfPseudoLegalMoves returns the number pseudolegal moves.
func (pos *Position) GenerateCountOfPseudoLegalMoves() int {
	// update convenience bitboards
	pos.Bd.UpdateConvenienceBBs()

	var plplies int
	piece := NoPiece

	// get all possible pieces able to move
	moverPieces := pos.Bd.PieceOccupancy(pos.Turn)

	// loop over source squares
	for startSq := 0; startSq < NumSquaresInBoard; startSq++ {
		// if the mover has a piece there
		if moverPieces.Occupied(Square(startSq)) {
			piece = pos.Bd.Piece(Square(startSq))
			// calculate all the pieces moves
			mvVector := pos.Bd.MovesVector(Square(startSq))
			for endSq := 0; endSq < NumSquaresInBoard; endSq++ {
				// for each piece move
				if mvVector.Occupied(Square(endSq)) {
					// check if mover is a pawn and destination square is on either the 1st or 6th ranks and add promotions
					if piece.PieceType() == Pawn && (Square(endSq).Rank() == Rank1 || Square(endSq).Rank() == Rank6) {
						plplies += 3 // can promote to knight, queen, or rook, alternatively, len(PromotionPieceTypes)
					} else {
						// if not a pawn, increment pseudolegal plies
						plplies++
					}
				}
			}
		}
	}
	return plplies
}

// InsufficientMaterial returns true if both sides only have kings
func (pos *Position) InsufficientMaterial() bool {
	// fmt.Println("Piece Count: ", bits.OnesCount64(uint64(pos.Bd.PieceOccupancy(pos.Turn))))
	return bits.OnesCount64(uint64(pos.Bd.PieceOccupancy(pos.Turn))) < 2 && bits.OnesCount64(uint64(pos.Bd.PieceOccupancy(pos.Turn.Other()))) < 2
}

// Result returns the current result of the position
func (pos *Position) Result() Result {
	// check insufficient material
	if pos.InsufficientMaterial() {
		return Draw
	}
	// check checkmate
	numLegalMoves := pos.GenerateCountOfLegalMoves()
	if numLegalMoves == 0 {
		// if pos.Bd.InCheck(White) && White == pos.Turn {
		// 	return NewResultWin(Black)
		// } else if pos.Bd.InCheck(Black) && Black == pos.Turn {
		// 	return NewResultWin(White)
		// } else {
		// 	fmt.Println("SOmethign Worng")
		// }

		if pos.Bd.InCheck(pos.Turn) {
			return NewResultWin(pos.Turn.Other())
		}
		return Draw

	}
	return InPlay
}

// ZobristHash is a map of random numbers used to calculate the Zobrist hash of a position
var ZobristHash map[Square]map[Piece]uint64

// ZobristHash returns a Zobrist hash of the position using the table
func (pos *Position) ZobristHash() uint64 {
	return 0
}
