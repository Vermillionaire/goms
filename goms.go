package main

import (
	"net/http"

	"github.com/gobuffalo/packr"
)

// Consider using echo for http library

func main() {
	loadMedia("/home/brian/Pictures")
	box := packr.NewBox("./web-interface")

	http.Handle("/", http.FileServer(box))
	http.ListenAndServe(":3000", nil)
}
