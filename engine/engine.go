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

type gameState uint8

const (
	Active    gameState = 0
	Checkmate gameState = 1
	Stalemate gameState = 2
)

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
func Eval(b Board, m Move) (next Board, state gameState, err error) {
	// 1: check to see if the move is vaild
	err = validMove(b, m)
	if err != nil {
		return b, Active, err
	}
	// 2: check to see if it will result in active player being placed in check
	// if either is true return error illegal move as well as the current board and a fasle game ended state
	nBoard := makeMove(b, m)
	chkNBoard := nBoard
	chkNBoard.turn = b.turn
	if inCheck(chkNBoard) {
		return b, Active, fmt.Errorf("cannot move into check")
	} else if inCheckMate(nBoard) {
		return nBoard, Checkmate, nil
	} else if isDraw(nBoard) {
		return nBoard, Stalemate, nil
	}
	return nBoard, Active, nil
	//	if both false m we will then check to see if the game has ended
}

func inCheckMate(b Board) bool {
	// check every move to see if it will result in a check

	// iterate though all the current player's moves checking for check each time

	// if any move will result in a non check position then return false

	// after checking all pieces and no vailid safe moves return true
	panic("Not implimented")
}

func isDraw(b Board) bool {
	panic("Not implimented")
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
	// create fake game where last player can go again to see if king is in danger
	chkboard := copyBoard(b)
	if chkboard.turn == White {
		chkboard.turn = Black
	} else if chkboard.turn == Black {
		chkboard.turn = White
	} else {
		panic("Its is no ones turn and it should be someone turn --!! error showed inCheck funciton")
	}
	// find the square with the king with the correct color on it loop once
	var kloc = loc{}
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			if b.b[f][r].kind == King && b.b[f][r].color == b.turn {
				kloc.f, kloc.r = f, r
				break
			}
		}
	}
	// check if any piece of the oppsite color can legally move to attack the king
	// loop over the entire board
	// if piece coor != b.color check it against the kings square
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			if chkboard.b[f][r].kind != Empty && chkboard.b[f][r].color == chkboard.turn {
				chkchk := loc{f, r}
				m := Move{chkchk, kloc}
				if validMove(chkboard, m) == nil {
					return true
				}
			}
		}
	}
	return false
}

func copyBoard(b Board) Board {
	copy := b
	return copy
}

func displayBoard(b Board) {
	fmt.Printf("%v", b)
}
