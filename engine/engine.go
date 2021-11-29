package engine

import "fmt"

type Board [8][8]Piece

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

type Piece struct {
	n     rune
	white bool
	// may add complexitiy later
}

// Evaluate a board position.
func Eval(b Board, m Move) (next Board, winning bool, err error) {
	panic("not implemented")
}

func validMove(b Board, m Move) bool {

}

func setupBoard() Board {
	b := Board{}
	// set White/black back rank
	b[0][0] = Piece{'R', true}
	b[0][7] = Piece{'R', false}
	b[1][0] = Piece{'N', true}
	b[1][7] = Piece{'N', false}
	b[2][0] = Piece{'B', true}
	b[2][7] = Piece{'B', false}
	b[3][0] = Piece{'Q', true}
	b[3][7] = Piece{'Q', false}
	b[4][0] = Piece{'K', true}
	b[4][7] = Piece{'K', true}
	b[5][0] = Piece{'B', true}
	b[5][7] = Piece{'B', false}
	b[6][0] = Piece{'N', true}
	b[6][7] = Piece{'N', false}
	b[7][0] = Piece{'R', true}
	b[7][7] = Piece{'R', false}
	// set White/black pawns pawns
	for i := 0; i < len(b); i++ {
		b[i][1] = Piece{'P', true}
		b[i][6] = Piece{'P', false}
	}
	return b
}

func displayBoard(b Board) {
	for r := 7; r >= 0; r-- {
		fmt.Printf("|")
		for f := 0; f <= 8; f++ {
			fmt.Printf("%2s|", b[f][r].n)
		}
		fmt.Printf("\n")
	}
}
