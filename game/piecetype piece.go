package game

// PieceType represents a type of piece (e.g. Pawn)
type PieceType int8

// String returns a string representation of PieceType - e.g. "Pawn" for a Pawn PieceType.
// Implements the fmt.Stringer interface.
func (p PieceType) String() string {
	switch p {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	}
	return ""
}

const (
	// NoPieceType represents the type of 'no piece'
	NoPieceType PieceType = iota

	// Pawn represents a pawn
	Pawn
	// Knight represents a knight
	Knight
	// Rook represents a rook
	Rook
	// Queen represents a queen
	Queen
	// King represents a king
	King
)

//--------------------------------------------------------------------------------

// Piece is one of the 10 PieceType and Color combinations.
type Piece int8

// PieceType gets the PieceType
func (p Piece) PieceType() PieceType {
	switch p {
	case BlackPawn, WhitePawn:
		return Pawn
	case BlackKnight, WhiteKnight:
		return Knight
	case BlackRook, WhiteRook:
		return Rook
	case BlackQueen, WhiteQueen:
		return Queen
	case BlackKing, WhiteKing:
		return King
	}
	return NoPieceType
}

// String gets a string representation
func (p Piece) String() string {
	switch p {
	case WhitePawn:
		return "White Pawn"
	case WhiteKnight:
		return "White Knight"
	case WhiteRook:
		return "White Rook"
	case WhiteQueen:
		return "White Queen"
	case WhiteKing:
		return "White King"

	case BlackPawn:
		return "Black Pawn"
	case BlackKnight:
		return "Black Knight"
	case BlackRook:
		return "Black Rook"
	case BlackQueen:
		return "Black Queen"
	case BlackKing:
		return "Black King"
	}
	return "NoPieceType"
}

// Color gets the Color
func (p Piece) Color() Color {
	switch p {
	case WhiteKing, WhiteQueen, WhiteRook, WhiteKnight, WhitePawn:
		return White
	case BlackKing, BlackQueen, BlackRook, BlackKnight, BlackPawn:
		return Black
	}
	return NoColor
}

const (
	// NoPiece represents no piece (i.e. empty square).
	NoPiece Piece = iota

	// WhitePawn is a white pawn.
	WhitePawn
	// WhiteKnight is a white knight.
	WhiteKnight
	// WhiteRook is a white rook.
	WhiteRook
	// WhiteQueen is a white queen.
	WhiteQueen
	// WhiteKing is a white king.
	WhiteKing

	// BlackPawn is a black pawn.
	BlackPawn
	// BlackKnight is a black knight.
	BlackKnight
	// BlackRook is a black rook.
	BlackRook
	// BlackQueen is a black queen.
	BlackQueen
	// BlackKing is a black king.
	BlackKing
)

var (
	// AllPieces is a slice of all Pieces.
	AllPieces = []Piece{
		WhitePawn, WhiteKnight, WhiteRook, WhiteQueen, WhiteKing,
		BlackPawn, BlackKnight, BlackRook, BlackQueen, BlackKing,
	}

	// PieceSymbMap is a mapping from Pieces to their Unicode symbols.
	PieceSymbMap = map[Piece]string{
		NoPiece: ".",

		WhitePawn:   "♙",
		WhiteKnight: "♘",
		WhiteRook:   "♖",
		WhiteQueen:  "♕",
		WhiteKing:   "♔",

		BlackPawn:   "♟",
		BlackKnight: "♞",
		BlackRook:   "♜",
		BlackQueen:  "♛",
		BlackKing:   "♚",
	}

	// PieceCharMap is a mapping from Pieces to their FEN characters (e.g. White Rook = 'R').
	PieceCharMap = map[Piece]string{
		NoPiece: "_",

		WhitePawn:   "P",
		WhiteKnight: "N",
		WhiteRook:   "R",
		WhiteQueen:  "Q",
		WhiteKing:   "K",

		BlackPawn:   "p",
		BlackKnight: "n",
		BlackRook:   "r",
		BlackQueen:  "q",
		BlackKing:   "k",
	}
)
