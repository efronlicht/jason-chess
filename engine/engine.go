package engine

import (
	"fmt"
)

type Board struct {
	b    [8][8]Square
	turn Color
}

type Move struct {
	start, end loc
}

type loc struct {
	f, r int
}

// var File int(
// 	'a' = 0,
// 	'b' = 1,
// 	'c' = 2,
// 	'd' = 3,
// 	'e' = 4,
// 	'f' = 5,
// 	'g' = 6,
// 	'h' = 7,
// )

type Kind uint8

const (
	Empty  Kind = 0
	Pawn   Kind = 1
	Bishop Kind = 2
	Queen  Kind = 3
	King   Kind = 4
	Rook   Kind = 5
	Knight Kind = 6
)

type Color uint8

const (
	Black Color = 1
	White Color = 2
)

// Square is a white piece, a black piece, or empty.
// The zero value is the empty square.
type Square struct {
	kind  Kind
	color Color
}

// Evaluate a board position.
func Eval(b Board, m Move) (next Board, winning bool, err error) {
	panic("not implemented")
}

func (b Board) String() string {
	out := ""
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			out += b.b[f][r].String()
		}
		out += "\n"
	}
	return out
}

func (s Square) String() string {
	if s.color == 0 && s.kind == 0 {
		return "   "
	}
	switch s.color {
	default:
		// this is an error, since a non White or Black Color is only OK for the zero value, which we already handled.
		panic("No color asigned to piece but piece assigned")
	case White:
		switch s.kind {
		default:
			panic("White square with no piece, check piece moving code")
			// this is an error, since the Empty kind is only OK for the zero value, which we already handled.
		case Pawn:
			return " ♙ "
		case Bishop:
			return " ♗ "
		case Knight:
			return " ♘ "
		case Queen:
			return " ♕ "
		case King:
			return " ♔ "
		case Rook:
			return " ♖ "
		}
	case Black:
		switch s.kind {
		default:
			panic("Black piece sqare with no piece, check moving code")
			// this is an error, since the Empty kind is only OK for the zero value, which we already handled.
		case Pawn:
			return " ♟ "
		case Bishop:
			return " ♝ "
		case Knight:
			return " ♞ "
		case Queen:
			return " ♛ "
		case King:
			return " ♚ "
		case Rook:
			return " ♜ "
		}
	}
}

func inCheck(b Board)bool {
	panic("Not implimented")
}

func validMove(b Board, m Move) bool {
	p := b.b[m.start.f][m.start.r]
	turn := b.turn
	// rules to move a piece
		// 1: no piece can be in the way --!!knights ignore this rule!!--
		// 2: cannot land on a square with a piece you own
		// 3: cannot place your king in check
		// 4: must move out of check 
		// 5a: pawns must move forward
		//	5b: pawns can only take sideways
		//  5c: if pawns make it to the final rank they should get an choice to be promoted
		// 6: cannot move off the board
	switch p {
	default:
		// no piece on square selected
		return false
	case :

	}
	return false
}

func setupBoard() Board {
	b := Board{}
	b.turn = White
	// set White/black back rank
	b.b[0][0] = Square{Rook, White}
	b.b[0][7] = Square{Rook, Black}
	b.b[1][0] = Square{Knight, White}
	b.b[1][7] = Square{Knight, Black}
	b.b[2][0] = Square{Bishop, White}
	b.b[2][7] = Square{Bishop, Black}
	b.b[3][0] = Square{Queen, White}
	b.b[3][7] = Square{Queen, Black}
	b.b[4][0] = Square{King, White}
	b.b[4][7] = Square{King, White}
	b.b[5][0] = Square{Bishop, White}
	b.b[5][7] = Square{Bishop, Black}
	b.b[6][0] = Square{Knight, White}
	b.b[6][7] = Square{Knight, Black}
	b.b[7][0] = Square{Rook, White}
	b.b[7][7] = Square{Rook, Black}
	// set White/black pawns pawns
	for i := 0; i < len(b.b); i++ {
		b.b[i][1] = Square{Pawn, White}
		b.b[i][6] = Square{Pawn, Black}
	}
	return b
}

func displayBoard(b Board) {
	fmt.Printf("%v", b)

}
