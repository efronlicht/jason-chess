package engine

import (
	"bytes"
	"fmt"
	"log"
	"text/tabwriter"
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

func (c Color) other() Color {
	if c == Black {
		return White
	} else if c == White {
		return Black
	}
	panic("cannot call other on a non color")
}

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
func Eval(b Board, m Move, turn Color) (next Board, state gameState, err error) {
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
	if inCheck(chkNBoard, turn) {
		return b, Active, fmt.Errorf("cannot move into check")
	} else if inCheckMate(nBoard, turn) {
		return nBoard, Checkmate, nil
	} else if isDraw(nBoard, turn) {
		return nBoard, Stalemate, nil
	}
	return nBoard, Active, nil
	//	if both false m we will then check to see if the game has ended
}

func listOfAllPlayerPieces(b Board, color Color) []loc {
	pieces := []loc{}
	for r := 0; r < 8; r++ {
		for f := 0; f < 8; f++ {
			if b.b[r][f].color == color {
				pieces = append(pieces, loc{r, f})
			}
		}
	}
	return pieces
}

func inCheckMate(b Board, threttened Color) bool {
	pieces := listOfAllPlayerPieces(b, threttened)
	// check every move to see if it will result in a check
	for _, l := range pieces {
		moves := movesFromSquare(b, l)
		// make a move and test if current player still in check
		for _, m := range moves { // iterate though all the current player's moves checking for check each time
			fakeBoard := copyBoard(b)
			fakeBoard = makeMove(fakeBoard, m)
			// if any move will result in a non check position then return false
			if !inCheck(fakeBoard, threttened) {
				log.Printf(fakeBoard.String())
				return false
			}
		}
	}
	// after checking all pieces and no vailid safe moves return true
	return true
}

func isDraw(b Board, turn Color) bool {
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
	buf := new(bytes.Buffer)
	tw := tabwriter.NewWriter(buf, 1, 1, 0, ' ', 0)
	tw.Write([]byte("\n"))
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			// empty squares are colored black or white
			if (b.b[f][r] == Piece{}) {
				if (f+r)%2 == 0 {
					tw.Write([]byte("■"))
				} else {
					tw.Write([]byte("□"))
				}
			}
			tw.Write([]byte(b.b[f][r].String()))
			tw.Write([]byte("\t"))
		}
		tw.Write([]byte("\n"))
	}
	tw.Flush()
	return buf.String()
}

func (s Piece) String() string {
	if s.color == 0 && s.kind == 0 {
		return " "
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
			return "♙"
		case Bishop:
			return "♗"
		case Knight:
			return "♘"
		case Queen:
			return "♕"
		case King:
			return "♔"
		case Rook:
			return "♖"
		}
	case Black:
		switch s.kind {
		default:
			panic("Black piece sqare with no piece, check moving code")
			// this is an error, since the Empty kind is only OK for the zero value, which we already handled.
		case Pawn:
			return "♟"
		case Bishop:
			return "♝"
		case Knight:
			return "♞"
		case Queen:
			return "♛"
		case King:
			return "♚"
		case Rook:
			return "♜"
		}
	}
}

// EFRON NOTE: no need to use this, but it might help.
/*
func (c Color) Enemy() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	default:
		panic(fmt.Errorf("no enemy color for unknown color %d",c))
	}
}
*/

func inCheck(b Board, inThreat Color) bool {
	// create fake game where last player can go again to see if king is in danger
	chkboard := copyBoard(b)

	// find the square with the king with the correct color on it loop once
	var kloc loc
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			if b.b[f][r].kind == King && b.b[f][r].color == inThreat {
				kloc.f, kloc.r = f, r
				break
			}
		}
		// EFRON NOTE: what if you don't find the king?
		if r == 0 {
			panic(fmt.Sprintf("no king %v on board", inThreat))
		}
	}

	// check if any piece of the oppsite color can legally move to attack the king
	// loop over the entire board
	// if piece coor != b.color check it against the kings square

	pieces := listOfAllPlayerPieces(chkboard, inThreat.other())
	for _, p := range pieces {
		m := Move{p, kloc}
		if validMove(chkboard, m) == nil {
			return true
		}
	}
	return false
}

// EFRON NOTE: this function isn't necessary.
func copyBoard(b Board) Board {
	copy := b
	return copy
}

func displayBoard(b Board) {
	fmt.Printf("%v", b)
}
