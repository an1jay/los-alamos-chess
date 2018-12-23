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
	capStr := "x"
	if p.Capture == false {
		capStr = "-"
	}
	return p.SourceSq.String() + capStr + p.DestinationSq.String() + "=" + p.Promotion.String()
}

// Equals checks whether two Plies are the same.
func (p *Ply) Equals(p2 *Ply) bool {
	return (p.SourceSq == p2.SourceSq) &&
		(p.DestinationSq == p2.DestinationSq) &&
		(p.Promotion == p2.Promotion) &&
		(p.Capture == p2.Capture) &&
		(p.Side == p2.Side)
}
