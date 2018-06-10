package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gobuffalo/packr"
)

// Consider using echo for http library

func mediaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recived path: " + r.URL.Path)

	s := strings.Split(r.URL.Path, "/")

	switch s[2] {
	case "files":
		if len(s) < 4 {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		file := mediaFiles[s[3]]
		if file == nil {
			fmt.Println("Could not find media file: " + s[3])
			errorHandler(w, r, http.StatusNotFound)
		} else {
			http.ServeFile(w, r, file.path)
		}
	default:
		fmt.Println("Unknown path: " + r.URL.Path)
		errorHandler(w, r, http.StatusNotFound)
	}
}

func main() {
	loadMedia("/home/brian/Pictures")
	box := packr.NewBox("./web-interface")

	http.Handle("/", http.FileServer(box))
	http.HandleFunc("/media/", mediaHandler)
	http.ListenAndServe(":3000", nil)
}

// Borrowed from:
// https://stackoverflow.com/questions/9996767/showing-custom-404-error-page-with-standard-http-package
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 Not found")
	}
}
