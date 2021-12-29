package engine

import (
	"fmt"
)

type Board struct {
	b    [8][8]Piece
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

// Piece is a white piece, a black piece, or empty.
// The zero value is the empty square.
type Piece struct {
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

func listOfAllPlayerPieces(b Board) []loc {
	pieces := []loc{}
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			if b.b[r][f].color == b.turn {
				pieces = append(pieces, loc{r, f})
			}
		}
	}
	return pieces
}

func inCheckMate(b Board) bool {
	pieces := listOfAllPlayerPieces(b)
	// check every move to see if it will result in a check
	for _, l := range pieces {
		moves := movesFromSquare(b, l)
		// make a move and test if current player still in check
		for _, m := range moves { // iterate though all the current player's moves checking for check each time
			fakeBoard := copyBoard(b)
			fakeBoard = makeMove(fakeBoard, m)
			if fakeBoard.turn == White {
				fakeBoard.turn = Black
			} else if fakeBoard.turn == Black {
				fakeBoard.turn = White
			}
			// if any move will result in a non check position then return false
			if !inCheck(fakeBoard) {
				return false
			}
		}
	}
	// after checking all pieces and no vailid safe moves return true
	return true
}

func isDraw(b Board) bool {
	// chk for stalemate
	if stalemate(b) {
		return true
	}
	// chk for rep of possition NEed make changes for the not part of mvp
	// check for 50 moves without taking a piece
	// not enough pieces to mate
	return false
}

func stalemate(b Board) bool {
	panic("not yet implimented")
}

func (b Board) String() string {
	out := " "
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			out += b.b[f][r].String()
		}
		out += "\n"
	}
	return out
}

func (s Piece) String() string {
	if s.color == 0 && s.kind == 0 {
		return fmt.Sprintf("%3v", "")
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
			return fmt.Sprintf("%3v", "♙")
		case Bishop:
			return fmt.Sprintf("%3v", "♗")
		case Knight:
			return fmt.Sprintf("%3v", "♘")
		case Queen:
			return fmt.Sprintf("%3v", "♕")
		case King:
			return fmt.Sprintf("%3v", "♔")
		case Rook:
			return fmt.Sprintf("%3v", "♖")
		}
	case Black:
		switch s.kind {
		default:
			panic("Black piece sqare with no piece, check moving code")
			// this is an error, since the Empty kind is only OK for the zero value, which we already handled.
		case Pawn:
			return fmt.Sprintf("%3v", "♟")
		case Bishop:
			return fmt.Sprintf("%3v", "♝")
		case Knight:
			return fmt.Sprintf("%3v", "♞")
		case Queen:
			return fmt.Sprintf("%3v", "♛")
		case King:
			return fmt.Sprintf("%3v", "♚")
		case Rook:
			return fmt.Sprintf("%3v", "♜")
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
