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
	HashList      []uint64
}

// NewPosition constructs a new position.
func NewPosition(bd *Board, turn Color, moveNumber, halfMoveClock uint, hashlist []uint64) *Position {
	return &Position{
		Bd:            bd,
		Turn:          turn,
		MoveNumber:    moveNumber,
		HalfMoveClock: halfMoveClock,
		InCheck:       bd.InCheck(turn),
		HashList:      hashlist,
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
	if !pos.LegalPly(p) {
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
	pos.InCheck = false
	if pos.Turn == White && pos.Bd.InCheck(White) {
		pos.InCheck = true
	} else if pos.Turn == Black && pos.Bd.InCheck(Black) {
		pos.InCheck = true
	}
	pos.HashList = append(pos.HashList, pos.ZobristHash())

}

// LegalPly returns whether a move is legal.
func (pos *Position) LegalPly(p *Ply) bool {
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
	WhiteMaterial := pos.Bd.MaterialCount(White)
	BlackMaterial := pos.Bd.MaterialCount(Black)
	var whiteInsuff = false
	var blackInsuff = false
	// insufficient material cases:
	// King v King
	// King + Knight v King
	// King + Knight v King + Knight
	if noQueenRookPawnChecker(WhiteMaterial, White) && noQueenRookPawnChecker(BlackMaterial, Black) {

		if WhiteMaterial[Knight] == 1 || WhiteMaterial[Knight] == 0 {
			whiteInsuff = true
		}
		if BlackMaterial[Knight] == 1 || BlackMaterial[Knight] == 0 {
			blackInsuff = true
		}
	}
	// fmt.Println("Piece Count: ", bits.OnesCount64(uint64(pos.Bd.PieceOccupancy(pos.Turn))))
	//return bits.OnesCount64(uint64(pos.Bd.PieceOccupancy(pos.Turn))) < 2 && bits.OnesCount64(uint64(pos.Bd.PieceOccupancy(pos.Turn.Other()))) < 2
	return whiteInsuff && blackInsuff
}

// noQueenRookPawnChecker checks if the map from piece to int has no
func noQueenRookPawnChecker(m map[PieceType]int, c Color) bool {
	if m[Queen] == 0 && m[Rook] == 0 && m[Pawn] == 0 {
		return true
	}
	return false
}

// Result returns the current result of the position
func (pos *Position) Result() Result {
	// check insufficient material and threefold repetition
	if pos.InsufficientMaterial() || pos.threefoldRepetition() {
		return Draw
	}
	// check checkmate and stalmate
	numLegalMoves := pos.GenerateCountOfLegalMoves()
	if numLegalMoves == 0 {
		if pos.Bd.InCheck(pos.Turn) {
			return NewResultWin(pos.Turn.Other())
		}
		// if no legal moves and not in check, stalemate
		return Draw
	}
	return InPlay
}

// threefoldRepetition returns bool of whether oor not three fold repeition has occured
func (pos *Position) threefoldRepetition() bool {
	currenthash := pos.ZobristHash()
	var counter uint8
	for i := len(pos.HashList) - 1; i >= 0; i-- {
		if pos.HashList[i] == currenthash {
			counter++
			if counter > 2 {
				return true
			}
		}
	}
	return false
}

// ZobristHashMap is a map of random numbers used to calculate the Zobrist hash of a position
var ZobristHashMap = map[Square]map[Piece]uint64{
	A1: map[Piece]uint64{NoPiece: 11773386922175966982, WhitePawn: 16287117459286898574, WhiteKnight: 1596030731072225293, WhiteRook: 6908381805289842349, WhiteQueen: 1471855846966644088,
		WhiteKing: 13946128460607735265, BlackPawn: 933481857740680097, BlackKnight: 4822007064700398354, BlackRook: 6198724923208203700, BlackQueen: 17946889222579085311, BlackKing: 6270309927917970168},
	A2: map[Piece]uint64{NoPiece: 6072176031918769629, WhitePawn: 17925223329797815625, WhiteKnight: 14615379492287389764, WhiteRook: 8922529181293848742, WhiteQueen: 2186575892497350080,
		WhiteKing: 11255819383249683110, BlackPawn: 9455613954418737896, BlackKnight: 8787529630854465939, BlackRook: 15080278555260971654, BlackQueen: 10822588433193164194, BlackKing: 6098849504152491340},
	A3: map[Piece]uint64{NoPiece: 16527254978915964266, WhitePawn: 16031150834041543796, WhiteKnight: 3536296654515895987, WhiteRook: 1482412321804371265, WhiteQueen: 15815380194503815770,
		WhiteKing: 11780042619164385330, BlackPawn: 12226390934754713528, BlackKnight: 2624779950949833695, BlackRook: 5204219537379946078, BlackQueen: 5630195549636017477, BlackKing: 8992578851473221649},
	A4: map[Piece]uint64{NoPiece: 17492292331866303211, WhitePawn: 5865978329968989632, WhiteKnight: 3575330857263971148, WhiteRook: 5135785284106165524, WhiteQueen: 17346646487532452867,
		WhiteKing: 17006776898343501925, BlackPawn: 14761012425106088848, BlackKnight: 1133808653343760861, BlackRook: 11705866394343143641, BlackQueen: 810765885738391414, BlackKing: 16402293007870477108},
	A5: map[Piece]uint64{NoPiece: 5078965237013387312, WhitePawn: 10996549525619867154, WhiteKnight: 8450082418903396944, WhiteRook: 10473422123106913351, WhiteQueen: 346624852092281243,
		WhiteKing: 8073009176360924998, BlackPawn: 17934417002597962725, BlackKnight: 2015319649236122759, BlackRook: 14747494385568980386, BlackQueen: 13074925490434954809, BlackKing: 17969637990791162656},
	A6: map[Piece]uint64{NoPiece: 5730624484641922827, WhitePawn: 4377432308003386796, WhiteKnight: 96352109645192172, WhiteRook: 6125506922645864851, WhiteQueen: 2073964105513269798,
		WhiteKing: 16890814920110874952, BlackPawn: 8546445605202623096, BlackKnight: 4036538071927409025, BlackRook: 12677955725882282243, BlackQueen: 12020999644872683379, BlackKing: 6910592320002715237},
	B1: map[Piece]uint64{NoPiece: 1239935007780297662, WhitePawn: 10386346689933916061, WhiteKnight: 12483840290261758634, WhiteRook: 16746065366981394713, WhiteQueen: 12407282130762189405,
		WhiteKing: 17966870885288538231, BlackPawn: 13698259473969289459, BlackKnight: 15755580579568985283, BlackRook: 10146528574378694998, BlackQueen: 17720873775737591970, BlackKing: 2472344647272992813},
	B2: map[Piece]uint64{NoPiece: 7629937874342804614, WhitePawn: 10638063951634777872, WhiteKnight: 9527009095712169158, WhiteRook: 6014976400866454716, WhiteQueen: 11093461923709284204,
		WhiteKing: 4104818376804604415, BlackPawn: 7259830900705720713, BlackKnight: 10285184083555264822, BlackRook: 2517620139808330519, BlackQueen: 6868627547023817416, BlackKing: 11454128984415816989},
	B3: map[Piece]uint64{NoPiece: 2168322695413405769, WhitePawn: 1006578471928546547, WhiteKnight: 15239214459205079742, WhiteRook: 16582141539562978657, WhiteQueen: 10278894395748860717,
		WhiteKing: 8450509298816644077, BlackPawn: 6804526565015884229, BlackKnight: 4881063388067503983, BlackRook: 15532760663777006515, BlackQueen: 4340950135463547861, BlackKing: 12036808534414370123},
	B4: map[Piece]uint64{NoPiece: 4659925661469946721, WhitePawn: 1748211115516405655, WhiteKnight: 8272311948344737425, WhiteRook: 11159173626914545888, WhiteQueen: 4855086305095776812,
		WhiteKing: 661670590802585445, BlackPawn: 6138305466401791049, BlackKnight: 4086856245824389211, BlackRook: 16964749118989876245, BlackQueen: 6516335962816781140, BlackKing: 8714844605611571402},
	B5: map[Piece]uint64{NoPiece: 6901648874693365610, WhitePawn: 14907669940587679858, WhiteKnight: 13275990456205445542, WhiteRook: 6427450627099029935, WhiteQueen: 13572279908122645822,
		WhiteKing: 2386384972929881978, BlackPawn: 2648459995346717954, BlackKnight: 6436395110811492248, BlackRook: 14622790242863698231, BlackQueen: 9072396571794037216, BlackKing: 12092276400206542330},
	B6: map[Piece]uint64{NoPiece: 15683397801645903382, WhitePawn: 7158564836161238555, WhiteKnight: 6936053491250265091, WhiteRook: 3175677080297903372, WhiteQueen: 17112093807046837214,
		WhiteKing: 13160269491655870207, BlackPawn: 4895718224712188793, BlackKnight: 342238063497134066, BlackRook: 10740151585763382863, BlackQueen: 16972109461368274062, BlackKing: 2077955963127131129},
	C1: map[Piece]uint64{NoPiece: 6412973179820290764, WhitePawn: 2857391983283342310, WhiteKnight: 4811709422441345720, WhiteRook: 15954483095551476410, WhiteQueen: 6176328532375587924,
		WhiteKing: 15229555803117540684, BlackPawn: 6019564927511236213, BlackKnight: 5927302832190330030, BlackRook: 4343152757778933172, BlackQueen: 1056343617103278623, BlackKing: 573592684560851023},
	C2: map[Piece]uint64{NoPiece: 2566822059360505664, WhitePawn: 2646335313407054082, WhiteKnight: 9957102452461020210, WhiteRook: 4602790725508405197, WhiteQueen: 3344920104025556521,
		WhiteKing: 15982512922683758958, BlackPawn: 11499171466371752289, BlackKnight: 13440786927624358665, BlackRook: 1898715308285314618, BlackQueen: 535771651424357824, BlackKing: 11092943512157989895},
	C3: map[Piece]uint64{NoPiece: 11428892452587414914, WhitePawn: 10231758163641119656, WhiteKnight: 5348527799442395555, WhiteRook: 11260025147690344619, WhiteQueen: 3333233221876710855,
		WhiteKing: 17140011901125957097, BlackPawn: 11870358074147416952, BlackKnight: 11844390212946384503, BlackRook: 12316106461439024120, BlackQueen: 15972260403073356387, BlackKing: 17589983155673301174},
	C4: map[Piece]uint64{NoPiece: 9104620987201590065, WhitePawn: 6174065723103611825, WhiteKnight: 15893416705534328851, WhiteRook: 6678160223715031183, WhiteQueen: 8719772738382758561,
		WhiteKing: 17727663361366289203, BlackPawn: 10764198003842199362, BlackKnight: 10767615613795335979, BlackRook: 2907586754536733572, BlackQueen: 354039519418428305, BlackKing: 10703033643104006943},
	C5: map[Piece]uint64{NoPiece: 2773758589587582151, WhitePawn: 6703082939502293683, WhiteKnight: 5430892227613520690, WhiteRook: 16797971007539366787, WhiteQueen: 16561228637489628198,
		WhiteKing: 15886291964460668600, BlackPawn: 13321945168050029494, BlackKnight: 9833906374481384516, BlackRook: 1722726788765740262, BlackQueen: 4144203252800625622, BlackKing: 14882892329984662852},
	C6: map[Piece]uint64{NoPiece: 3428848391633257808, WhitePawn: 2066521119266587960, WhiteKnight: 5311678803005135959, WhiteRook: 12961880221293659866, WhiteQueen: 11365033512335886647,
		WhiteKing: 16611455622960898934, BlackPawn: 9401084867147280953, BlackKnight: 8228877743144769722, BlackRook: 9551054865368678677, BlackQueen: 5367371904371263628, BlackKing: 17725181219846962250},
	D1: map[Piece]uint64{NoPiece: 17638576701428372885, WhitePawn: 1202785698995402117, WhiteKnight: 2498941524475391162, WhiteRook: 12382573326056141144, WhiteQueen: 4933516684309646975,
		WhiteKing: 2565226072262439014, BlackPawn: 9048573062567709169, BlackKnight: 12165350819594822749, BlackRook: 3456255634607401815, BlackQueen: 9857173880552626642, BlackKing: 1588256792758203023},
	D2: map[Piece]uint64{NoPiece: 9569858698258528322, WhitePawn: 17642898734978408639, WhiteKnight: 14359492202908392017, WhiteRook: 17655903896892849175, WhiteQueen: 1330063395150230557,
		WhiteKing: 14620674357213100841, BlackPawn: 13320464234745932877, BlackKnight: 17458510489672546543, BlackRook: 6928323940484480081, BlackQueen: 9275827688279428170, BlackKing: 12238427242695405671},
	D3: map[Piece]uint64{NoPiece: 9334487589767743786, WhitePawn: 14283292896746652930, WhiteKnight: 1504954434549055802, WhiteRook: 10417478159726925200, WhiteQueen: 11350739486867002710,
		WhiteKing: 10606359796971861698, BlackPawn: 7385329502892071359, BlackKnight: 7139007937393578263, BlackRook: 16668706679756223954, BlackQueen: 17909296079823128394, BlackKing: 11506602781942324516},
	D4: map[Piece]uint64{NoPiece: 5124801153691074703, WhitePawn: 5844345924858341013, WhiteKnight: 10538737058250563884, WhiteRook: 5503838126036851703, WhiteQueen: 13249862433994079541,
		WhiteKing: 335841766725459219, BlackPawn: 11719773066685559593, BlackKnight: 1750744690127340514, BlackRook: 810700744967443909, BlackQueen: 513256583388648931, BlackKing: 10261952756938359446},
	D5: map[Piece]uint64{NoPiece: 1571054462922775895, WhitePawn: 11693330577297800729, WhiteKnight: 12864995713890574981, WhiteRook: 17932127288754035425, WhiteQueen: 6892051907783460189,
		WhiteKing: 10994195510102705059, BlackPawn: 5817031884319575247, BlackKnight: 17346603534636869025, BlackRook: 106777638686652344, BlackQueen: 5065835324994511329, BlackKing: 3906341670389511792},
	D6: map[Piece]uint64{NoPiece: 5683895659794663571, WhitePawn: 9664435531502648904, WhiteKnight: 5778951933276307539, WhiteRook: 12744509813347914554, WhiteQueen: 17108694082608486709,
		WhiteKing: 3569944486643882010, BlackPawn: 11482280678979612811, BlackKnight: 3910371669405967471, BlackRook: 14715419190142488886, BlackQueen: 7983635611872497642, BlackKing: 10674895694521208845},
	E1: map[Piece]uint64{NoPiece: 5658617193308509212, WhitePawn: 11525794121241867711, WhiteKnight: 10536599933475811259, WhiteRook: 7939481273209203936, WhiteQueen: 9457790795157968586,
		WhiteKing: 16190246637726474128, BlackPawn: 17382097824132401245, BlackKnight: 15380453529833825092, BlackRook: 1375799755191178987, BlackQueen: 12151254022552636739, BlackKing: 1186296719702877822},
	E2: map[Piece]uint64{NoPiece: 10612086517840482733, WhitePawn: 11942108215327123125, WhiteKnight: 16985085901107738691, WhiteRook: 7428016039012787895, WhiteQueen: 346395341476481881,
		WhiteKing: 2161321100412602207, BlackPawn: 11535305452995508703, BlackKnight: 10293468691176545363, BlackRook: 17868059861134303242, BlackQueen: 12687427384015048656, BlackKing: 10178860562321799581},
	E3: map[Piece]uint64{NoPiece: 11663905951255169233, WhitePawn: 427182107144632515, WhiteKnight: 7316404717623904185, WhiteRook: 14139819536556545541, WhiteQueen: 12316549601801392658,
		WhiteKing: 6904697297395536213, BlackPawn: 6789849735265882199, BlackKnight: 8055426494981633879, BlackRook: 18202428938912730070, BlackQueen: 3061086519933419853, BlackKing: 13212631844504336054},
	E4: map[Piece]uint64{NoPiece: 4624631714125480411, WhitePawn: 3907983101626336110, WhiteKnight: 11217994730132662973, WhiteRook: 10617231302332827099, WhiteQueen: 12939021186978275046,
		WhiteKing: 1135572291329184406, BlackPawn: 9755692541292980568, BlackKnight: 4755831847496153374, BlackRook: 8755350343101780224, BlackQueen: 4038446427481660896, BlackKing: 18375596241001554818},
	E5: map[Piece]uint64{NoPiece: 11727193977121028681, WhitePawn: 18180287017938245255, WhiteKnight: 13042925507364908893, WhiteRook: 16705919510350577863, WhiteQueen: 18047868130386505331,
		WhiteKing: 17346683156043913384, BlackPawn: 17205067572282183206, BlackKnight: 7574146711693955872, BlackRook: 13078297199372302824, BlackQueen: 4442906032018553964, BlackKing: 3710350131097370058},
	E6: map[Piece]uint64{NoPiece: 10458643697206784426, WhitePawn: 17966753912542058981, WhiteKnight: 16482410131381734431, WhiteRook: 14988974852773770115, WhiteQueen: 9495932568727549658,
		WhiteKing: 472014930458144894, BlackPawn: 11618783473691210058, BlackKnight: 3251418374518936931, BlackRook: 15419889733993074151, BlackQueen: 253374321517741122, BlackKing: 288575612421696812},
	F1: map[Piece]uint64{NoPiece: 1360100442136714265, WhitePawn: 13740299683458060087, WhiteKnight: 8516540595627298802, WhiteRook: 13644453717855479901, WhiteQueen: 17342672294120558105,
		WhiteKing: 15526058230052370786, BlackPawn: 13344332063336180619, BlackKnight: 11718887548006995903, BlackRook: 13524726797816805690, BlackQueen: 8184790379158690337, BlackKing: 16533666512202728587},
	F2: map[Piece]uint64{NoPiece: 5272287977708953033, WhitePawn: 1505912239496101634, WhiteKnight: 2233742591613390424, WhiteRook: 322844231205426855, WhiteQueen: 12094525015859507268,
		WhiteKing: 4162764563160837132, BlackPawn: 11419309217220353963, BlackKnight: 8845411249013691181, BlackRook: 17321360843734533330, BlackQueen: 1596953527460594128, BlackKing: 12805506799257378101},
	F3: map[Piece]uint64{NoPiece: 173625333634244122, WhitePawn: 13873301547367597972, WhiteKnight: 7953950595959262141, WhiteRook: 18220400653196628710, WhiteQueen: 9851732300812193879,
		WhiteKing: 10817932178294175381, BlackPawn: 14513047952726448946, BlackKnight: 15821471142276492618, BlackRook: 12773357221497416190, BlackQueen: 585931912418038684, BlackKing: 14511326117114479800},
	F4: map[Piece]uint64{NoPiece: 10701091667561070349, WhitePawn: 14625645431557203954, WhiteKnight: 13915509363803329956, WhiteRook: 18252245524172127747, WhiteQueen: 10180003846323696710,
		WhiteKing: 3574696389312514267, BlackPawn: 16294972208671315245, BlackKnight: 980613923437787904, BlackRook: 18400219482907826235, BlackQueen: 18426585929219480346, BlackKing: 17962694460894615475},
	F5: map[Piece]uint64{NoPiece: 14265349171563724632, WhitePawn: 14150084214488760745, WhiteKnight: 14725895992650464589, WhiteRook: 9210223609540779732, WhiteQueen: 11766142961358431158,
		WhiteKing: 13632035142610384552, BlackPawn: 14119663684533762624, BlackKnight: 6853243077979165393, BlackRook: 11082512107468916680, BlackQueen: 14703771657165611949, BlackKing: 14794854579590078302},
	F6: map[Piece]uint64{NoPiece: 15162016773395006737, WhitePawn: 15719567261319496588, WhiteKnight: 10381819373397584044, WhiteRook: 245505232520688868, WhiteQueen: 5096343611318020,
		WhiteKing: 11130238338614010427, BlackPawn: 6606077363668324549, BlackKnight: 8196222139501837355, BlackRook: 12299368809965038468, BlackQueen: 15805647502848571893, BlackKing: 11719373776528817052}}

var whiteHash uint64 = 13616133844141066041
var blackHash uint64 = 12360302557774378304

// ZobristHash returns a Zobrist hash of the position using the table
func (pos *Position) ZobristHash() uint64 {
	var hash uint64
	for sq := 0; sq < numSquaresInBoard; sq++ {
		hash ^= ZobristHashMap[Square(sq)][pos.Bd.Piece(Square(sq))]
	}

	switch pos.Turn {
	case White:
		hash ^= whiteHash
	case Black:
		hash ^= blackHash
	}

	return hash
}
