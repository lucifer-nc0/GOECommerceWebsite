package main

import (
	"net/http"

	"github.com/lucifer-nc0/test/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
