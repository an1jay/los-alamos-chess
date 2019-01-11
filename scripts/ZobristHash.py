import random
for letter in ['a','b','c','d','e','f']:
    for num in ['1','2','3','4','5','6']:
        print(letter.upper() + num + ': map[Piece]uint64{' , end = ' ')
        for piece in ['NoPiece', 'WhitePawn', 'WhiteKnight', 
        'WhiteRook', 'WhiteQueen', 'WhiteKing', 'BlackPawn', 
        'BlackKnight', 'BlackRook', 'BlackQueen', 'BlackKing']:
            print(piece + ': ' + str(random.randint(0, 2**64 - 1)), end = ', ')
        print('} , ')
