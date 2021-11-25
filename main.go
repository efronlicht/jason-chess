package main

import (
	"net/http"
)

// build the router that assigns paths to http handlers,
// i.e, GET /ping -> Ping
// POST /game/new ->
func buildRouter() http.Handler {
	panic("not implemented")
}

func main() {
	http.ListenAndServe(":8080", buildRouter())
}
