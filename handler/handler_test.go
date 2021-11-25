package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_Ping(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/ping", nil)
	fmt.Println("the request is")
	r.Write(os.Stdout)
	Ping(w, r)
	resp := w.Result()
	fmt.Println("the response is")
	resp.Write(os.Stdout)
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected the status %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}
