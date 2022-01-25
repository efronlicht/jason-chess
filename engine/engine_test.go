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

var mp = map[string]Piece{
	"WP": {Pawn, White},
	"WB": {Bishop, White},
	"WN": {Knight, White},
	"WR": {Rook, White},
	"WQ": {Queen, White},
	"WK": {King, White},
	"BP": {Pawn, Black},
	"BB": {Bishop, Black},
	"BN": {Knight, Black},
	"BR": {Rook, Black},
	"BQ": {Queen, Black},
	"BK": {King, Black},
}

func parseBoardFromString(t Color, s string) Board {
	// example turn White
	// example string "Wkg8 BQg7 bkg6"
	nb := Board{turn: t}
	pieces := strings.Fields(s)
	for _, p := range pieces {
		piece := strings.ToUpper(p[:2])
		loc := locFromNotation(p[2:])
		if (mp[piece] == Piece{}) {
			panic(fmt.Sprintf("unkonwn piece %v cannot parse", piece))
		}
		nb.b[loc.f][loc.r] = mp[piece]
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
	type test struct {
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
			b := setupBoard()
			if validMove(b, moveFromNotation(move)) == nil {
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

type test struct {
	name  string
	input string
	want  string
}

func Test_EndingPositions(t *testing.T) {
	t.Run("black to move: white checkmate", func(t *testing.T) {
		for _, notation := range []string{
			"Bkh8 Wkg6 Wqh7", "Bke8 wra8 wrc7 wkc2", "wke6 bke8 wra8",
		} {
			board := parseBoardFromString(Black, notation)
			if !inCheckMate(board, Black) {
				t.Errorf("expected white checkmate\n%s", board)
			}
		}
	})
	t.Run("white to move: black checkmate", func(t *testing.T) {
		for _, notation := range []string{
			"Bka6 brg1 bqh8 wkh4", "wkd1 bkc3 bnd3 bbf3", "wkd1 bkd3 bbd2 bne3",
		} {
			board := parseBoardFromString(White, notation)
			if !inCheckMate(board, White) {
				t.Errorf("expected white checkmate\n%s", board)
			}
		}
	})
	t.Run("white to move: check but not checkmate", func(t *testing.T) {
		for _, notation := range []string{
			// TODO
		} {
			board := parseBoardFromString(Black, notation)
			if check, mate := inCheck(board, White), inCheckMate(board, White); !check || mate {
				t.Errorf("expected black to be in check (%v) but NOT in checkmate (%v) \n%s", check, mate, board)
			}
		}
	})
	t.Run("black to move: check but not checkmate", func(t *testing.T) {
		for _, notation := range []string{
			// TODO
		} {
			board := parseBoardFromString(Black, notation)
			if check, mate := inCheck(board, Black), inCheckMate(board, Black); !check || mate {
				t.Errorf("expected black to be in check (%v) but NOT in checkmate (%v) \n%s", check, mate, board)
			}
		}
	})
}

var boardStates = map[string][]string{
	"simple WCM":  {"Bkh8 Wkg6 Wqh7", "Bke8 wra8 wrc7 wkc2", "wke6 bke8 wra8"},
	"simple WSM":  {"Bkh8 Wqh6 WKb5", "Bkh8 Wqh6 WKb5 wpd5 bpd6", "Bkh8 Wqh6 WKb5 wpd5 bpd6 bpe7 wpe6 bpb6 bpa7 bnc8 wpa6"},
	"simple BCM":  {"Bka6 brg1 bqh8 wkh4", "wkd1 bkc3 bnd3 bbf3", "wkd1 bkd3 bbd2 bne3"},
	"simple BSM":  {""},
	"complex WSM": {"bkd3 wkh1", "bkd3 wkh1 bnf7", "bkd3 wkh1 bbf7"},
}
