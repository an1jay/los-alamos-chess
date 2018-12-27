package game

import (
	"fmt"
	"math/bits"
)

// Board represents the pieces on the Los Alamos chess board
type Board struct {
	wPawns   BitBoard
	wKnights BitBoard
	wRooks   BitBoard
	wQueen   BitBoard
	wKing    BitBoard

	bPawns   BitBoard
	bKnights BitBoard
	bRooks   BitBoard
	bQueen   BitBoard
	bKing    BitBoard

	bPieces  BitBoard
	wPieces  BitBoard
	emptySqs BitBoard
}

// SquareMap returns mapping of squares to pieces (at those squares).
// Squares with 'NoPiece' are not included in the map
func (b *Board) SquareMap() map[Square]Piece {
	m := map[Square]Piece{}
	for sq := 0; sq < numSquaresInBoard; sq++ {
		p := b.Piece(Square(sq))
		if p != NoPiece {
			m[Square(sq)] = p
		}
	}
	return m
}

// Piece returns the piece at that square
func (b *Board) Piece(sq Square) Piece {
	for _, p := range AllPieces {
		bb := b.BitBoardForPiece(p)
		if bb.Occupied(sq) {
			return p
		}
	}
	return NoPiece
}

// Display prints a representation of the Board to stdout using Unicode symbols.
func (b *Board) Display(symb bool) {
	m := PieceCharMap
	if symb {
		m = PieceSymbMap
	}
	b.FuncDisplay(func(sq Square) string { return m[b.Piece(sq)] })
}

// SquareDisplay prints the squares in their positions
func (b *Board) SquareDisplay() {
	b.FuncDisplay(func(sq Square) string { return sq.String() })
}

// FuncDisplay outputs a string representation given by function f to stdout
func (b *Board) FuncDisplay(f func(Square) string) {
	fmt.Printf(
		"   A B C D E F\n"+
			"6. %s %s %s %s %s %s\n"+
			"5. %s %s %s %s %s %s\n"+
			"4. %s %s %s %s %s %s\n"+
			"3. %s %s %s %s %s %s\n"+
			"2. %s %s %s %s %s %s\n"+
			"1. %s %s %s %s %s %s\n",
		f(A6), f(B6), f(C6), f(D6), f(E6), f(F6),
		f(A5), f(B5), f(C5), f(D5), f(E5), f(F5),
		f(A4), f(B4), f(C4), f(D4), f(E4), f(F4),
		f(A3), f(B3), f(C3), f(D3), f(E3), f(F3),
		f(A2), f(B2), f(C2), f(D2), f(E2), f(F2),
		f(A1), f(B1), f(C1), f(D1), f(E1), f(F1))
}

// BoardFromMap generates a NewBoard from a Square->Piece map
func BoardFromMap(m map[Square]Piece) *Board {
	b := Board{}
	for _, p := range AllPieces {
		pMap := map[Square]bool{}
		for sq := 0; sq < numSquaresInBoard; sq++ {
			pMap[Square(sq)] = m[Square(sq)] == p
		}
		b.SetBBForPiece(p, BitBoardFromMap(pMap))
	}
	b.UpdateConvenienceBBs()
	return &b
}

// CheckBBsAreValid checks that the most significant 28 bits are zero.
func (b *Board) CheckBBsAreValid() bool {
	rv := true
	for _, p := range AllPieces {
		rv = rv && b.BitBoardForPiece(p)&BBValidityCheck == 0
	}
	return rv
}

// DebugBitBoards prints out, for every piece, the chaacter representation, and displays the bitboard for that piece.
func (b *Board) DebugBitBoards() {
	for _, p := range AllPieces {
		fmt.Println(PieceCharMap[p])
		b.BitBoardForPiece(p).Display()
	}
}

// BitBoardForPiece returns a copy of the bitboard for piece p
func (b *Board) BitBoardForPiece(p Piece) BitBoard {
	switch p {
	case WhitePawn:
		return b.wPawns
	case WhiteKnight:
		return b.wKnights
	case WhiteRook:
		return b.wRooks
	case WhiteQueen:
		return b.wQueen
	case WhiteKing:
		return b.wKing
	case BlackPawn:
		return b.bPawns
	case BlackKnight:
		return b.bKnights
	case BlackRook:
		return b.bRooks
	case BlackQueen:
		return b.bQueen
	case BlackKing:
		return b.bKing
	}
	return BitBoard(1 << 63)
}

// MaterialCount returns a copy of the bitboard for piece p
func (b *Board) MaterialCount(c Color) map[PieceType]int {
	m := map[PieceType]int{}
	for _, pt := range AllPieceTypes {
		m[pt] = bits.OnesCount64(uint64(b.BitBoardForPiece(NewPiece(pt, c))))
	}
	return m
}

// SetBBForPiece sets the bitboard for piece p to b
func (b *Board) SetBBForPiece(p Piece, bb BitBoard) {
	switch p {
	case WhitePawn:
		b.wPawns = bb
	case WhiteKnight:
		b.wKnights = bb
	case WhiteRook:
		b.wRooks = bb
	case WhiteQueen:
		b.wQueen = bb
	case WhiteKing:
		b.wKing = bb
	case BlackPawn:
		b.bPawns = bb
	case BlackKnight:
		b.bKnights = bb
	case BlackRook:
		b.bRooks = bb
	case BlackQueen:
		b.bQueen = bb
	case BlackKing:
		b.bKing = bb
	default:
		panic("Board.SetBBForPiece() -> Invalid Piece")
	}
}

// PieceOccupancy returns a bitboards of the pieces of a certain color.
// *Call UpdateConvenienceBoards() before using this*
func (b *Board) PieceOccupancy(c Color) BitBoard {
	if c == White {
		return b.wPieces
	}
	return b.bPieces
}

// UpdateConvenienceBBs updates the wPieces, bPieces and emptySqs bitboards
func (b *Board) UpdateConvenienceBBs() {
	b.bPieces = b.bKing | b.bQueen | b.bRooks | b.bKnights | b.bPawns
	b.wPieces = b.wKing | b.wQueen | b.wRooks | b.wKnights | b.wPawns
	b.emptySqs = ^(b.bPieces | b.wPieces)
}

// Move simply puts a move on the board - does not perform legality check (will simply make the move if an illegal move is entered).
func (b *Board) Move(source, destination Square, promotionPiece Piece) {
	movePiece := b.Piece(source)
	capturePiece := b.Piece(destination)
	// set bitboard for movePiece at source square zero
	b.SetBBForPiece(movePiece, b.BitBoardForPiece(movePiece).SetSquareOnBB(source, false))
	// set bitboard for capPiece at destination square zero
	if capturePiece != NoPiece {
		b.SetBBForPiece(capturePiece, b.BitBoardForPiece(capturePiece).SetSquareOnBB(destination, false))
		// b.setBBForPiece(capturePiece, b.bbForPiece(capturePiece)^BBForSquare(destination))
	}
	// if promotionPiece, make destination square of movepiece bitboard 1, else make the destination square of promotionPiece bitboard 1.
	if promotionPiece == NoPiece {
		// set bitboard for movePiece at destination square one
		b.SetBBForPiece(movePiece, b.BitBoardForPiece(movePiece).SetSquareOnBB(destination, true))
		// b.setBBForPiece(movePiece, b.bbForPiece(movePiece)^BBForSquare(destination))
	} else {
		b.SetBBForPiece(promotionPiece, b.BitBoardForPiece(promotionPiece).SetSquareOnBB(destination, true))
	}
}

// Copy returns a pointer to the copy of a board
func (b *Board) Copy() *Board {
	newb := Board{
		wPawns:   b.wPawns,
		wKnights: b.wKnights,
		wRooks:   b.wRooks,
		wQueen:   b.wQueen,
		wKing:    b.wKing,
		bPawns:   b.bPawns,
		bKnights: b.bKnights,
		bRooks:   b.bRooks,
		bQueen:   b.bQueen,
		bKing:    b.bKing,
		bPieces:  b.bPieces,
		wPieces:  b.wPieces,
		emptySqs: b.emptySqs,
	}
	return &newb
}

//--------------------------------------------------------------------------------

// MovesVector returns a bitboard with the possible move destination squares highlighted for the piece at the square.
func (b *Board) MovesVector(sq Square) BitBoard {
	// need to update covenience bitboards
	b.UpdateConvenienceBBs()

	// calculate things which will be useful
	p := b.Piece(sq)
	pt := p.PieceType()
	color := p.Color()
	sqbb := sq.BitBoard()
	sqint := int(sq)
	occupied := ^b.emptySqs

	// switch on piecetype
	switch pt {
	case King:
		return BBKingMoves[sqint] &^ b.PieceOccupancy(color)
	case Queen:
		rookAtk := (CalcLinearAttack(occupied, sq.FileBB(), sqbb) ^ CalcLinearAttack(occupied, sq.RankBB(), sqbb))
		bishAtk := (CalcLinearAttack(occupied, BBDiagonals[sqint], sqbb) ^ CalcLinearAttack(occupied, BBAntiDiagonals[sqint], sqbb))
		return (rookAtk ^ bishAtk) &^ b.PieceOccupancy(color)
	case Rook:
		return (CalcLinearAttack(occupied, sq.FileBB(), sqbb) ^ CalcLinearAttack(occupied, sq.RankBB(), sqbb)) &^ b.PieceOccupancy(color)
	case Knight:
		return BBKnightMoves[sqint] &^ b.PieceOccupancy(color)
	case Pawn:
		lm := BitBoard(0)
		if color == White {
			lm = BBWhitePawnPushes[sqint] & b.emptySqs
			return lm ^ (BBWhitePawnCaptures[sqint] & b.bPieces)
		}
		lm = BBBlackPawnPushes[sqint] & b.emptySqs
		return lm ^ (BBBlackPawnCaptures[sqint] & b.wPieces)
	}

	// If no piece at that square, no moves are possible
	return BitBoard(0)
}

// CalcLinearAttack calculates attacks for a sliding piece located at PiecePos, which can legally move to pieceMoves, and is blocked by occupied
func CalcLinearAttack(occupied, pieceMoves, piecePos BitBoard) BitBoard {
	OccupiedInMask := occupied & pieceMoves
	return ((OccupiedInMask - 2*piecePos) ^ (OccupiedInMask.Reverse() - 2*piecePos.Reverse()).Reverse()) & pieceMoves
}

// KingSquare returns the square of the king of the color c
func (b *Board) KingSquare(c Color) Square {
	if c == White {
		for sq := 0; sq < NumSquaresInBoard; sq++ {
			if Square(sq).BitBoard() == b.wKing {
				return Square(sq)
			}
		}
	}
	for sq := 0; sq < NumSquaresInBoard; sq++ {
		if Square(sq).BitBoard() == b.bKing {
			return Square(sq)
		}
	}
	panic("No King on board")
}

// SquareAttacked checks whether sq is attacked by the side attacker
func (b *Board) SquareAttacked(attackSq Square, attacker Color) bool {
	attackVector := BitBoard(0)
	pieces := b.PieceOccupancy(attacker)
	for sq := 0; sq < numSquaresInBoard; sq++ {
		if pieces.Occupied(Square(sq)) {
			attackVector = attackVector | b.MovesVector(Square(sq))
		}
	}
	if attackVector.Occupied(attackSq) {
		return true
	}
	return false
}

// InCheck checks whether King sq is attacked
func (b *Board) InCheck(c Color) bool {
	KingSq := b.KingSquare(c)
	return b.SquareAttacked(KingSq, c.Other())
}
