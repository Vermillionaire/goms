package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
)

type media struct {
	name     string
	id       string
	path     string
	iconPath string
}

var mediaFiles map[string]*media

func loadMedia(path string) {
	log.Info().Msg("Scanning for media to serve...")
	if mediaFiles == nil {
		mediaFiles = make(map[string]*media)
	}

	err := filepath.Walk(path, visit)

	if err != nil {
		log.Error().
			Str("path", path).
			Str("error", err.Error()).
			Msg("There was an error searching for files in the path.")
	}
}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		if f.Name() == ".icons" {
			log.Debug().
				Str("folder", ".icons").
				Msg("Skipping folder")
			return filepath.SkipDir
		} else {
			return nil
		}
	}

	log.Debug().
		Str("file", path).
		Msg("Found file")

	ext := filepath.Ext(path)

	if isImage(ext) {

		image := new(media)
		image.id = uuid()
		image.name = filepath.Base(path)
		image.path = path
		image.iconPath = filepath.Dir(path) + "/.icons/" + image.name

		log.Info().
			Str("name", image.name).
			Str("id", image.id).
			Msg("Adding image to server.")

		mediaFiles[image.id] = image

	} else {
		log.Debug().
			Str("file", path).
			Msg("Not a supported file type")
	}

	return nil
}

var imageTypes = [...]string{".tif", ".tiff", ".gif", ".jpeg", ".jpg", ".jif", ".jp2", ".jpx", ".png"}

// Checks is a extension is a supported type
func isImage(extension string) bool {
	for _, v := range imageTypes {
		if strings.Compare(v, extension) == 0 {
			return true
		}
	}
	return false
}

// Temporary solution for uuid
func uuid() string {
	out, err := exec.Command("/usr/bin/uuidgen").Output()
	if err != nil {
		log.Fatal().
			Str("command", "/usr/bin/uuidgen").
			Msg("There was an error generating the uuid.")
	}

	//n := bytes.IndexByte(out, 0)
	s := string(out)
	s = strings.TrimSpace(s)
	return s
}
