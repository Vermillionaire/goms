package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Consider using echo for http library

func mediaHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().
		Str("type", r.Method).
		Str("path", r.URL.Path).
		Msg("Recieved http reqest to the server.")

	if r.URL.Path == "/media/" {
		html := buildMediaHTML()
		fmt.Fprintf(w, "%s", html)
		return
	}

	s := strings.Split(r.URL.Path, "/")

	switch s[2] {
	case "icons":
		if len(s) < 4 {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		file := mediaFiles[s[3]]
		if file == nil {
			log.Debug().
				Str("id", s[3]).
				Msg("Could not find media file.")
			errorHandler(w, r, http.StatusNotFound)
		} else {
			http.ServeFile(w, r, file.iconPath)
		}

	case "files":
		if len(s) < 4 {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		file := mediaFiles[s[3]]
		if file == nil {
			log.Debug().
				Str("id", s[3]).
				Msg("Could not find media file.")
			errorHandler(w, r, http.StatusNotFound)
		} else {
			http.ServeFile(w, r, file.path)
		}
	default:
		errorHandler(w, r, http.StatusNotFound)
	}
}

// Borrowed from:
// https://stackoverflow.com/questions/9996767/showing-custom-404-error-page-with-standard-http-package
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	log.Error().Int("code", status).Str("path", r.URL.Path).Msg("Returned a error to client.")
}

var maxWidth float32 = 400

func buildMediaHTML() string {
	var html = ""
	for k, v := range mediaFiles {
		var itemCard card
		itemCard.Src = "/media/icons/" + k
		itemCard.Name = v.name
		itemCard.Width = int(maxWidth)
		itemCard.Height = int(maxWidth / float32(v.width) * float32(v.height))
		html += itemCard.loadHTML() + "\n"
	}

	return html
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("GOMS server is starting!")
	go loadMedia("/home/brian/Pictures")
	go loadTemplates()

	mainPage := packr.NewBox("./web-interface")

	http.Handle("/", http.FileServer(mainPage))
	http.HandleFunc("/media/", mediaHandler)
	http.ListenAndServe(":3000", nil)
}
