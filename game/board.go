package game

import "fmt"

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
		" \tA\tB\tC\tD\tE\tF\n"+
			"6\t%s\t%s\t%s\t%s\t%s\t%s\n"+
			"5\t%s\t%s\t%s\t%s\t%s\t%s\n"+
			"4\t%s\t%s\t%s\t%s\t%s\t%s\n"+
			"3\t%s\t%s\t%s\t%s\t%s\t%s\n"+
			"2\t%s\t%s\t%s\t%s\t%s\t%s\n"+
			"1\t%s\t%s\t%s\t%s\t%s\t%s\n",
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
		panic("Board.setBBForPiece() -> Invalid Piece")
	}
}

// UpdateConvenienceBBs updates the wPieces, bPieces and emptySqs bitboards
func (b *Board) UpdateConvenienceBBs() {
	b.bPieces = b.bKing | b.bQueen | b.bRooks | b.bKnights | b.bPawns
	b.wPieces = b.wKing | b.wQueen | b.wRooks | b.wKnights | b.wPawns
	b.emptySqs = ^(b.bPieces | b.wPieces)
}

// Move simply puts a move on the board - does not perform legality check (will simply make the move if an illegal move is entered).
func (b *Board) Move(source, destination Square, promotionPiece Piece) bool {
	movePiece := b.Piece(source)
	capturePiece := b.Piece(destination)
	if movePiece == NoPiece {
		fmt.Println("Board.Move: trying to move NoPiece")
		return false
	}

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

	return true
}

// AttackVector attack vector returns a bitboard with the attacked squares highlighted.
func (b *Board) AttackVector(p Piece, sq Square) BitBoard {
	b.UpdateConvenienceBBs()

	panic("Todo")
	return BitBoard(0)
}

// CalcLinearAttack calculates attacks for a sliding piece located at PiecePos, which can legally move to pieceMoves, and is blocked by occupied
func CalcLinearAttack(occupied, pieceMoves, piecePos BitBoard) BitBoard {
	OccupiedInMask := occupied & pieceMoves
	return ((OccupiedInMask - 2*piecePos) ^ (OccupiedInMask.Reverse() - 2*piecePos.Reverse()).Reverse()) & pieceMoves //& bbMask
}

// RookAttacks returns the moves of a rook located at sq
func RookAttacks(sq Square, occupied BitBoard) BitBoard {
	return CalcLinearAttack(occupied, sq.FileBB(), sq.BitBoard()) ^ CalcLinearAttack(occupied, sq.RankBB(), sq.BitBoard())
}

// BishopAttacks returns the moves of a rook located at sq
func BishopAttacks(sq Square, occupied BitBoard) BitBoard {
	return CalcLinearAttack(occupied, BBDiagonals[int(sq)], sq.BitBoard()) ^ CalcLinearAttack(occupied, BBAntiDiagonals[int(sq)], sq.BitBoard())
}
