package game

// Ply represents one half-move made by either White or Black.
type Ply struct {
	// Origin square
	SourceSq Square

	// Destination square
	DestinationSq Square

	// What piece the moved piece promotes to - e.g.
	// A ply in which no promotion happens has promotion NoPieceType
	Promotion PieceType

	// Whether the ply involves a capture
	Capture bool

	// The side making the move
	Side Color
}

// String provides a string representation of the Ply - Source Square, Destination Square, Promotion.
// Implements the fmt.Stringer interface.
func (p *Ply) String() string {
	return p.SourceSq.String() + p.DestinationSq.String() + "=" + p.Promotion.String()
}
