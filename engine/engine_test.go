package engine

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// func TestMain(t *testing.M) {

// }

func Test_displayBoard(t *testing.T) {
	b := setupBoard()
	displayBoard(b)
	nb := parseBoardFromString(White, "WKe2 WQa1 bnh7 bRc4")
	displayBoard(nb)
}

func parseBoardFromString(t Color, s string) Board {
	// example turn White
	// example string "Wkg8 BQg7 bkg6"
	var mp = map[string]Piece{
		"WP": Piece{Pawn, White},
		"WB": Piece{Bishop, White},
		"WN": Piece{Knight, White},
		"WR": Piece{Rook, White},
		"WQ": Piece{Queen, White},
		"WK": Piece{King, White},
		"BP": Piece{Pawn, Black},
		"BB": Piece{Bishop, Black},
		"BN": Piece{Knight, Black},
		"BR": Piece{Rook, Black},
		"BQ": Piece{Queen, Black},
		"BK": Piece{King, Black},
	}
	nb := Board{turn: t}
	pieces := strings.Fields(s)
	for _, p := range pieces {
		piece := strings.ToUpper(p[:2])
		loc := locFromNotation(p[2:])
		if (mp[piece] == Piece{}) {
			panic(fmt.Sprintf("unkonwn piece %v cannot parse", piece))
		}
		nb.b[loc.r][loc.f] = mp[piece]
	}
	return nb
}

func locFromNotation(s string) (p loc) {
	s = strings.ToLower(s)
	switch f := s[0]; f {
	case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h':
		p.f = int(f - 'a')
	default:
		panic(fmt.Sprintf("unknown file %c", f))
	}
	switch r := s[1]; r {
	case '1', '2', '3', '4', '5', '6', '7', '8':
		p.r = int(r - '1')
	default:
		panic(fmt.Sprintf("unknown rank %c", r))

	}
	return p
}
func moveFromNotation(s string) (m Move) {
	split := strings.Fields(s)
	return Move{
		start: locFromNotation(split[0]),
		end:   locFromNotation(split[1]),
	}

}
func Test_validMove(t *testing.T) {
	type test = struct {
		name  string
		board Board
		move  Move
		want  bool
	}
	t.Run("valid white starting moves", func(t *testing.T) {
		for name, move := range map[string]string{
			"pawn e4":   "E2 E4",
			"knight c3": "b1 c3",
			"pawn h3":   "h2 h3",
			"knight f3": "g1 f3",
		} {
			if validMove(setupBoard(), moveFromNotation(move)) != nil {
				t.Errorf("%s: %s", name, move)
			}
		}
	})
	t.Run("valid black starting moves", func(t *testing.T) {
		for name, move := range map[string]string{
			"pawn 2x":   "E7 E5",
			"knight c6": "b8 c6",
			"pawn a6":   "a7 a6",
		} {
			board := setupBoard() // TODO: some valid start for black, like after e4
			board.turn = Black
			if validMove(board, moveFromNotation(move)) != nil {
				t.Errorf("%s: %s", name, move)
			}
		}
	})
	t.Run("invalid white starting moves", func(t *testing.T) {
		for name, move := range map[string]string{
			"king e4":  "E1 e4",
			"Queen h6": "d1 h5",
			"rook a7":  "a1 a7",
		} {
			if validMove(setupBoard(), moveFromNotation(move)) == nil {
				t.Errorf("%s: %s", name, move)
			}

		}
	})
}

func Test_Create_Game(t *testing.T) {
	r, err := http.NewRequest(http.MethodPost, "http://localhost:8080/new", nil)
	if err != nil {
		t.Fatalf("expected no error making request, but got %v", err)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	const wantStatus = http.StatusCreated
	if resp.StatusCode != wantStatus {
		t.Fatalf("expected status code %d, but got %d", wantStatus, resp.StatusCode)
	}
	t.Fatalf(`
	TODO: the response body should have the default chessboard. 
	We should also figure out a way to get the GAME ID.`,
	)
}
