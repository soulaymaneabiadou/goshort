package main

import (
	"net/http"

	"github.com/soulaymaneabiadou/goshort/api"
)

func main() {
	srv := api.NewServer()

	http.ListenAndServe(":7000", srv)
}
