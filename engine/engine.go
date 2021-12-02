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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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

func inCheck(b Board) bool {
	panic("Not implimented")
	return false
}

func copyBoard(b Board) Board {
	copy := b
	return copy
}

func makeMove(b Board, m Move) Board {
	if validMove(b, m) {

		e := Square{}
		square := b.b[m.start.f][m.start.r]
		b.b[m.end.f][m.end.r] = square
		b.b[m.start.f][m.start.r] = e
		return b
	}
	panic("Cannot make invalid move")
}

func validMove(b Board, m Move) bool {
	s := b.b[m.start.f][m.start.r]
	end := b.b[m.end.f][m.end.r]
	turn := b.turn
	// chk if there is a piece at square or if the color of the piece isn't the turn
	// piece on the squre we end up on is owned by the active player
	if s.kind == Empty || s.color != turn || end.color == turn {
		return false
	}
	switch s.kind {
	case Empty:
		// no piece on square selected
		return false
	case Pawn:
		if s.color == White {
			if m.start.r == 1 && b.b[m.start.f][m.start.r+1].kind == Empty && m.end.r == 3 && m.end.f == m.start.f { // pawn is on starting rank can move two squares
				return true
			} else if m.end.r == m.start.r+1 && m.start.f == m.end.f { // moving one space forward
				return true
			} else if m.end.r == m.start.r+1 && abs(m.start.f-m.end.f) == 1 && end.kind != Empty { // take a piece
				return true
			} // take a piece ADD EN PASSAUNT HERE!!!!
		} else if s.color == Black {
			if m.start.r == 6 && b.b[m.start.f][m.start.r-1].kind == Empty && m.end.r == 4 && m.end.f == m.start.f { // pawn is on starting rank can move two squares
				return true
			} else if m.end.r == m.start.r-1 && m.start.f == m.end.f { // moving one space forward
				return true
			} else if m.end.r == m.start.r-1 && abs(m.start.f-m.end.f) == 1 && end.kind != Empty { // take a piece
				return true
			} // take a piece ADD EN PASSAUNT HERE!!!!
		}
		panic("Our piece should have a color!!")

	case Knight:
		return (abs(m.start.f-m.end.f) == 2 && abs(m.start.r-m.end.r) == 1) || (abs(m.start.f-m.end.f) == 1 && abs(m.start.r-m.end.r) == 2)
	case Bishop:
		return validBishop(b, m)
	case Rook:
		return validRook(b, m)
	case Queen:
		return validBishop(b, m) || validRook(b, m)
	case King:
		return (abs(m.start.f-m.end.f)+abs(m.start.r-m.end.r) == 1) || (abs(m.start.f-m.end.f) == 1 && abs(m.start.r-m.end.r) == 1)
	default:
		panic(fmt.Sprintf("invalid kind %v", s.kind))
	}
	// rules to move a piece
	// 1: no piece can be in the way --!!knights ignore this rule!!--
	// 2: cannot land on a square with a piece you own // universal rule!
	// 3: cannot place your king in check
	// 4: must move out of check // if in check
	// 5a: pawns must move forward
	//	5b: pawns can only take sideways
	//  5c: if pawns make it to the final rank they should get an choice to be promoted
	// 6: cannot move off the board

	return false
}

func validRook(b Board, m Move) bool {
	if m.start.f == m.end.f || m.start.r == m.end.r {
		return false
	}
	var dx, dy int
	switch {
	default:
		panic("Impossible rook Possition")
	case m.end.f == m.start.f && m.end.r > m.start.r:
		dx, dy = 0, 1
	case m.end.f == m.start.f && m.end.r < m.start.r:
		dx, dy = 0, -1
	case m.end.r == m.start.r && m.end.f > m.start.f:
		dx, dy = 1, 0
	case m.end.r == m.start.r && m.end.f < m.start.f:
		dx, dy = -1, 0
	}
	for f, r := m.start.f, m.start.r; f != m.end.f; f, r = f+dx, r+dy { // step and check for pieces
		if b.b[f][r].kind != Empty {
			return false
		}
	}
	return true
}

func validBishop(b Board, m Move) bool {
	if abs(m.start.f-m.end.f) != abs(m.start.r-m.end.r) {
		return false
	}
	var dx, dy int
	switch {
	default:
		panic("Impossible bishop possition")
	case m.end.f > m.start.f && m.end.r > m.start.r:
		dx, dy = 1, 1
	case m.end.f > m.start.f && m.end.r < m.start.r:
		dx, dy = 1, -1
	case m.end.f < m.start.f && m.end.r > m.start.r:
		dx, dy = -1, 1
	case m.end.f < m.start.f && m.end.r < m.start.r:
		dx, dy = -1, -1
	}
	for f, r := m.start.f, m.start.r; f != m.end.f; f, r = f+dx, r+dy { // step and check for pieces
		if b.b[f][r].kind != Empty {
			return false
		}
	}
	return true
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
	b.b[4][7] = Square{King, Black}
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
