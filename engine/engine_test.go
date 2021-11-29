package engine

import (
	"net/http"
	"testing"
)

func TestMain(t *testing.M) {

}

func Test_displayBoard(t *testing.T) {

}

func Test_validMove(t *testing.T) {
	type test = struct {
		b    Board
		m    Move
		want bool
	}
	for _, tt := range []test{
		{},
		{},
		{},
	} {
		if got := validMove(tt.b, tt.m); got != tt.want {
			t.Fatalf("validMove(%b, %b) should be %v, but is: %v", tt.a, tt.b, tt.want, got)
		}
	}
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
