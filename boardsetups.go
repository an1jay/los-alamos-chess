package main

import "github.com/an1jay/los-alamos-chess/game"

// NewGame Position
var NewGame = map[game.Square]game.Piece{
	game.A1: game.WhiteRook,
	game.B1: game.WhiteKnight,
	game.C1: game.WhiteQueen,
	game.D1: game.WhiteKing,
	game.E1: game.WhiteKnight,
	game.F1: game.WhiteRook,
	game.A2: game.WhitePawn,
	game.B2: game.WhitePawn,
	game.C2: game.WhitePawn,
	game.D2: game.WhitePawn,
	game.E2: game.WhitePawn,
	game.F2: game.WhitePawn,
	game.A3: game.NoPiece,
	game.B3: game.NoPiece,
	game.C3: game.NoPiece,
	game.D3: game.NoPiece,
	game.E3: game.NoPiece,
	game.F3: game.NoPiece,
	game.A4: game.NoPiece,
	game.B4: game.NoPiece,
	game.C4: game.NoPiece,
	game.D4: game.NoPiece,
	game.E4: game.NoPiece,
	game.F4: game.NoPiece,
	game.A5: game.BlackPawn,
	game.B5: game.BlackPawn,
	game.C5: game.BlackPawn,
	game.D5: game.BlackPawn,
	game.E5: game.BlackPawn,
	game.F5: game.BlackPawn,
	game.A6: game.BlackRook,
	game.B6: game.BlackKnight,
	game.C6: game.BlackQueen,
	game.D6: game.BlackKing,
	game.E6: game.BlackKnight,
	game.F6: game.BlackRook,
}

// PuzzlingPos Position
var PuzzlingPos = map[game.Square]game.Piece{
	game.A1: game.WhiteRook,
	game.B1: game.WhiteKnight,
	game.C1: game.NoPiece,
	game.D1: game.WhiteKing,
	game.E1: game.WhiteKnight,
	game.F1: game.NoPiece,
	game.A2: game.NoPiece,
	game.B2: game.NoPiece,
	game.C2: game.NoPiece,
	game.D2: game.NoPiece,
	game.E2: game.NoPiece,
	game.F2: game.WhiteRook,
	game.A3: game.WhitePawn,
	game.B3: game.NoPiece,
	game.C3: game.NoPiece,
	game.D3: game.NoPiece,
	game.E3: game.NoPiece,
	game.F3: game.WhitePawn,
	game.A4: game.BlackKnight,
	game.B4: game.NoPiece,
	game.C4: game.NoPiece,
	game.D4: game.BlackPawn,
	game.E4: game.BlackKnight,
	game.F4: game.BlackRook,
	game.A5: game.BlackRook,
	game.B5: game.BlackPawn,
	game.C5: game.NoPiece,
	game.D5: game.BlackPawn,
	game.E5: game.NoPiece,
	game.F5: game.NoPiece,
	game.A6: game.NoPiece,
	game.B6: game.NoPiece,
	game.C6: game.NoPiece,
	game.D6: game.NoPiece,
	game.E6: game.NoPiece,
	game.F6: game.BlackKing,
}

// PuzzlingPos2 Position
var PuzzlingPos2 = map[game.Square]game.Piece{
	game.A1: game.WhiteRook,
	game.B1: game.WhiteKnight,
	game.C1: game.NoPiece,
	game.D1: game.WhiteKing,
	game.E1: game.WhiteKnight,
	game.F1: game.NoPiece,
	game.A2: game.NoPiece,
	game.B2: game.NoPiece,
	game.C2: game.WhiteRook,
	game.D2: game.NoPiece,
	game.E2: game.NoPiece,
	game.F2: game.NoPiece,
	game.A3: game.WhitePawn,
	game.B3: game.NoPiece,
	game.C3: game.NoPiece,
	game.D3: game.NoPiece,
	game.E3: game.NoPiece,
	game.F3: game.WhitePawn,
	game.A4: game.BlackKnight,
	game.B4: game.NoPiece,
	game.C4: game.NoPiece,
	game.D4: game.BlackPawn,
	game.E4: game.BlackKnight,
	game.F4: game.BlackRook,
	game.A5: game.BlackRook,
	game.B5: game.BlackPawn,
	game.C5: game.NoPiece,
	game.D5: game.BlackPawn,
	game.E5: game.NoPiece,
	game.F5: game.NoPiece,
	game.A6: game.NoPiece,
	game.B6: game.NoPiece,
	game.C6: game.NoPiece,
	game.D6: game.NoPiece,
	game.E6: game.NoPiece,
	game.F6: game.BlackKing,
}

// InterestingPos is  an Interesting Position
var InterestingPos = map[game.Square]game.Piece{
	game.A1: game.WhiteRook,
	game.B1: game.WhiteKnight,
	game.C1: game.WhiteQueen,
	game.D1: game.WhiteKing,
	game.E1: game.WhiteKnight,
	game.F1: game.WhiteRook,
	game.A2: game.WhitePawn,
	game.B2: game.WhitePawn,
	game.C2: game.WhitePawn,
	game.D2: game.NoPiece,
	game.E2: game.WhitePawn,
	game.F2: game.WhitePawn,
	game.A3: game.NoPiece,
	game.B3: game.NoPiece,
	game.C3: game.NoPiece,
	game.D3: game.WhitePawn,
	game.E3: game.NoPiece,
	game.F3: game.NoPiece,
	game.A4: game.NoPiece,
	game.B4: game.NoPiece,
	game.C4: game.BlackKnight,
	game.D4: game.NoPiece,
	game.E4: game.NoPiece,
	game.F4: game.NoPiece,
	game.A5: game.BlackPawn,
	game.B5: game.BlackPawn,
	game.C5: game.BlackPawn,
	game.D5: game.BlackPawn,
	game.E5: game.BlackPawn,
	game.F5: game.BlackPawn,
	game.A6: game.BlackRook,
	game.B6: game.NoPiece,
	game.C6: game.BlackQueen,
	game.D6: game.BlackKing,
	game.E6: game.BlackKnight,
	game.F6: game.BlackRook,
}

// CheckPos is  an Interesting Position
var CheckPos = map[game.Square]game.Piece{
	game.A1: game.WhiteRook,
	game.B1: game.WhiteKnight,
	game.C1: game.NoPiece,
	game.D1: game.WhiteKing,
	game.E1: game.WhiteKnight,
	game.F1: game.WhiteRook,
	game.A2: game.WhitePawn,
	game.B2: game.WhitePawn,
	game.C2: game.WhitePawn,
	game.D2: game.NoPiece,
	game.E2: game.WhitePawn,
	game.F2: game.WhitePawn,
	game.A3: game.NoPiece,
	game.B3: game.NoPiece,
	game.C3: game.NoPiece,
	game.D3: game.WhitePawn,
	game.E3: game.NoPiece,
	game.F3: game.NoPiece,
	game.A4: game.NoPiece,
	game.B4: game.NoPiece,
	game.C4: game.BlackKnight,
	game.D4: game.NoPiece,
	game.E4: game.BlackPawn,
	game.F4: game.WhiteQueen,
	game.A5: game.BlackPawn,
	game.B5: game.BlackPawn,
	game.C5: game.BlackPawn,
	game.D5: game.BlackPawn,
	game.E5: game.NoPiece,
	game.F5: game.BlackPawn,
	game.A6: game.BlackRook,
	game.B6: game.NoPiece,
	game.C6: game.BlackQueen,
	game.D6: game.BlackKing,
	game.E6: game.BlackKnight,
	game.F6: game.BlackRook,
}

// KNPawnGame Position; black to play
var KNPawnGame = map[game.Square]game.Piece{
	game.A1: game.NoPiece,
	game.B1: game.NoPiece,
	game.C1: game.NoPiece,
	game.D1: game.WhiteRook,
	game.E1: game.NoPiece,
	game.F1: game.NoPiece,
	game.A2: game.WhiteRook,
	game.B2: game.NoPiece,
	game.C2: game.WhitePawn,
	game.D2: game.NoPiece,
	game.E2: game.WhiteKing,
	game.F2: game.WhitePawn,
	game.A3: game.BlackPawn,
	game.B3: game.WhitePawn,
	game.C3: game.BlackPawn,
	game.D3: game.NoPiece,
	game.E3: game.NoPiece,
	game.F3: game.NoPiece,
	game.A4: game.NoPiece,
	game.B4: game.NoPiece,
	game.C4: game.NoPiece,
	game.D4: game.NoPiece,
	game.E4: game.BlackPawn,
	game.F4: game.NoPiece,
	game.A5: game.NoPiece,
	game.B5: game.NoPiece,
	game.C5: game.BlackPawn,
	game.D5: game.BlackKnight,
	game.E5: game.NoPiece,
	game.F5: game.NoPiece,
	game.A6: game.BlackRook,
	game.B6: game.NoPiece,
	game.C6: game.NoPiece,
	game.D6: game.BlackKing,
	game.E6: game.NoPiece,
	game.F6: game.BlackRook,
}

// KNPawnOpening Position; black to play
var KNPawnOpening = map[game.Square]game.Piece{
	game.A1: game.WhiteRook,
	game.B1: game.WhiteKnight,
	game.C1: game.WhiteQueen,
	game.D1: game.NoPiece,
	game.E1: game.WhiteKnight,
	game.F1: game.WhiteRook,
	game.A2: game.WhitePawn,
	game.B2: game.WhitePawn,
	game.C2: game.WhitePawn,
	game.D2: game.WhitePawn,
	game.E2: game.WhiteKing,
	game.F2: game.WhitePawn,
	game.A3: game.NoPiece,
	game.B3: game.NoPiece,
	game.C3: game.NoPiece,
	game.D3: game.NoPiece,
	game.E3: game.WhitePawn,
	game.F3: game.NoPiece,
	game.A4: game.NoPiece,
	game.B4: game.NoPiece,
	game.C4: game.NoPiece,
	game.D4: game.BlackPawn,
	game.E4: game.NoPiece,
	game.F4: game.NoPiece,
	game.A5: game.BlackPawn,
	game.B5: game.BlackPawn,
	game.C5: game.BlackPawn,
	game.D5: game.NoPiece,
	game.E5: game.BlackPawn,
	game.F5: game.BlackPawn,
	game.A6: game.BlackRook,
	game.B6: game.BlackKnight,
	game.C6: game.BlackQueen,
	game.D6: game.BlackKing,
	game.E6: game.BlackKnight,
	game.F6: game.BlackRook,
}
