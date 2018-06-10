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

	if r.URL.Path == "/media/" {
		html := buildMediaHTML()
		fmt.Printf("Built html: \n%s\n", html)
		fmt.Fprintf(w, "%s", html)
		return
	}

	s := strings.Split(r.URL.Path, "/")

	switch s[2] {
	case "icons":
		fmt.Println("Test")
		if len(s) < 4 {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		file := mediaFiles[s[3]]
		if file == nil {
			fmt.Println("Could not find media file: " + s[3])
			errorHandler(w, r, http.StatusNotFound)
		} else {
			fmt.Println(file.iconPath)
			http.ServeFile(w, r, file.iconPath)
		}

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

// Borrowed from:
// https://stackoverflow.com/questions/9996767/showing-custom-404-error-page-with-standard-http-package
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 Not found")
	}
}

func buildMediaHTML() string {
	var html = ""
	for k, v := range mediaFiles {
		var itemCard card
		itemCard.Src = "/media/icons/" + k
		itemCard.Name = v.name
		html += itemCard.loadHTML() + "\n"
	}

	return html
}

func main() {
	go loadMedia("/home/brian/Pictures")
	go loadTemplates()

	mainPage := packr.NewBox("./web-interface")

	http.Handle("/", http.FileServer(mainPage))
	http.HandleFunc("/media/", mediaHandler)
	http.ListenAndServe(":3000", nil)
}
