package main

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("spinning up test server")
	go http.ListenAndServe(":8080", buildRouter())
	log.Println("running tests")
	code := m.Run()
	os.Exit(code)
}

func Test_Ping(t *testing.T) {
	resp, err := http.DefaultClient.Get("http://localhost:8080/ping")
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	const wantStatus = http.StatusOK
	if resp.StatusCode != wantStatus {
		t.Fatalf("expected status code %d, but got %d", wantStatus, resp.StatusCode)
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

func Test_View_Game(t *testing.T) {
	// trying to view a game that can't exist should give a 404.
	resp, err := http.Get("http://localhost:8080/game/-1")
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}
	const wantStatus = http.StatusNotFound
	if resp.StatusCode != wantStatus {
		t.Fatalf("expected status code %d, but got %d", wantStatus, resp.StatusCode)
	}
	//
}
