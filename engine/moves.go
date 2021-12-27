package engine

import (
	"errors"
	"fmt"
)

func makeMove(b Board, m Move) Board {
	if err := validMove(b, m); err != nil {
		panic(fmt.Sprintf("Cannot make invalid move: %v", err))
	}
	// check if this move puts current player in check
	e := Square{}
	square := b.b[m.start.f][m.start.r]
	b.b[m.end.f][m.end.r] = square
	b.b[m.start.f][m.start.r] = e
	return b
}

func validPawnMv(b Board, m Move) bool {
	s := b.b[m.start.f][m.start.r]
	end := b.b[m.end.f][m.end.r]
	if s.color == White {
		if m.start.r == 1 && b.b[m.start.f][m.start.r+1].kind == Empty && m.end.r == 3 && m.end.f == m.start.f { // pawn is on starting rank can move two squares
			return true
		} else if m.end.r == m.start.r+1 && m.start.f == m.end.f { // moving one space forward
			return true
		} else if m.end.r == m.start.r+1 && abs(m.start.f-m.end.f) == 1 && end.kind != Empty && end.color != b.turn { // take a piece
			return true
		} // add white en passant

		// take a piece ADD EN PASSAUNT HERE!!!!
	} else if s.color == Black {
		if m.start.r == 6 && b.b[m.start.f][m.start.r-1].kind == Empty && m.end.r == 4 && m.end.f == m.start.f { // pawn is on starting rank can move two squares
			return true
		} else if m.end.r == m.start.r-1 && m.start.f == m.end.f { // moving one space forward
			return true
		} else if m.end.r == m.start.r-1 && abs(m.start.f-m.end.f) == 1 && end.kind != Empty && end.color != b.turn { // take a piece
			return true
		} // take a piece ADD EN PASSAUNT HERE!!!!
	} else {
		panic(fmt.Sprintf("our piece should have a color, color is %v", s.color))
	}
	return false
}

func validKnightMv(b Board, m Move) bool {
	return (abs(m.start.f-m.end.f) == 2 && abs(m.start.r-m.end.r) == 1) || (abs(m.start.f-m.end.f) == 1 && abs(m.start.r-m.end.r) == 2)
}

func validKingMv(b Board, m Move) bool {
	return (abs(m.start.f-m.end.f)+abs(m.start.r-m.end.r) == 1) || (abs(m.start.f-m.end.f) == 1 && abs(m.start.r-m.end.r) == 1)
}

func validRook(b Board, m Move) error {
	if m.start.f != m.end.f && m.start.r != m.end.r {
		return errors.New("rooks may only move in strait lines")
	}
	var dx, dy int
	switch {
	default:
		return errors.New("impossible rook Possition")
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
			return errors.New("rooks may not jump over other pieces")
		}
	}
	return nil
}

func validBishop(b Board, m Move) error {
	if abs(m.start.f-m.end.f) != abs(m.start.r-m.end.r) {
		return errors.New("this piece may only move diagnol")
	}
	var dx, dy int
	switch {
	default:
		return errors.New("impossible bishop possition")
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
			return errors.New("this piece may not jump over pieces")
		}
	}
	return nil
}

func validQueenMv(b Board, m Move) bool {
	return validBishop(b, m) == nil || validRook(b, m) == nil
}

func wrapErrToBool(f func(Board, Move) error) func(Board, Move) bool {
	return func(b Board, m Move) bool {
		return f(b, m) == nil
	}
}

func grabValidPieceFunction(p Kind) func(Board, Move) bool {
	switch p {
	default:
		panic(fmt.Sprintf("expected a Pawn Rook Bishon King Queen or Knight got %v", p))
	case Pawn:
		return validPawnMv
	case Bishop:
		return wrapErrToBool(validBishop)
	case Knight:
		return validKnightMv
	case Rook:
		return wrapErrToBool(validRook)
	case King:
		return validKingMv
	case Queen:
		return validQueenMv
	}
}

func movesFromSquare(b Board, l loc) (moves []Move) {
	s := b.b[l.r][l.f]
	vaildMvChk := grabValidPieceFunction(s.kind)
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if vaildMvChk(b, Move{start: loc{l.r, l.f}, end: loc{x, y}}) {
				moves = append(moves, Move{l, loc{x, y}})
			}
		}
	}
	return moves
}

func validMove(b Board, m Move) error {
	s := b.b[m.start.f][m.start.r]
	end := b.b[m.end.f][m.end.r]
	turn := b.turn
	// chk if there is a piece at square or if the color of the piece isn't the turn
	// piece on the squre we end up on is owned by the active player
	if s.kind == Empty {
		return errors.New("no piece selected to move")
	}
	if s.color != turn {
		return errors.New("not your piece to move")
	}
	if end.color == turn {
		return errors.New("not valid to capture own piece")
	}
	switch s.kind {
	case Empty:
		// no piece on square selected
		return errors.New("no piece selected to move")
	case Pawn:
		if validPawnMv(b, m) {
			return nil
		}
		return errors.New("pawns cannot move that way")
	case Knight:
		if validKnightMv(b, m) {
			return nil
		}
		return errors.New("the Knight cannot move here")
	case Bishop:
		err := validBishop(b, m)
		if err == nil {
			return nil
		}
		return err
	case Rook:
		err := validRook(b, m)
		if err == nil {
			return nil
		}
		return err
	case Queen:
		if validQueenMv(b, m) {
			return nil
		}
		return errors.New("the Queen has landed on an unreachable square")
	case King:
		if validKingMv(b, m) {
			return nil
		}
		return errors.New("the King has landed on an unreachable square")
	default:
		return fmt.Errorf("invalid Piece %v", s.kind)
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
