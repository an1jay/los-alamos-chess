package game

import (
	"fmt"
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

// Display prints a board representation of the position to stdout
func (pos *Position) Display(symb bool) {
	pos.Bd.Display(symb)
	checkStr := "not "
	if pos.InCheck {
		checkStr = ""
	}
	fmt.Printf("%s is to move. %s is %sin check\n\n", pos.Turn.String(), pos.Turn.String(), checkStr)
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
	pos.Turn = pos.Turn.Other()
	pos.InCheck = pos.Bd.InCheck(pos.Turn)
}

// LegalMove returns whether a move is legal.
func (pos *Position) LegalMove(p *Ply) bool {
	// Check side to move
	if !(p.Side == pos.Turn) {
		return false
	}
	// Check destination square is empty if capture is true
	if !(pos.Bd.Piece(p.DestinationSq) == NoPiece) {
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
					}
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
	return PseudoLegalPlies
}

// Result returns the current result of the position
func (pos *Position) Result() Result {
	// check checkmate
	numLegalMoves := len(pos.GenerateLegalMoves())
	if numLegalMoves == 0 && pos.InCheck {
		return NewResultWin(pos.Turn.Other())
	} else if numLegalMoves == 0 {
		return Draw
	}
	return InPlay
}
