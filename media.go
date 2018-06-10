package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type media struct {
	name     string
	id       string
	path     string
	iconPath string
}

var mediaFiles map[string]*media

func loadMedia(path string) {
	if mediaFiles == nil {
		mediaFiles = make(map[string]*media)
	}

	err := filepath.Walk(path, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		if f.Name() == ".icons" {
			return filepath.SkipDir
		} else {
			return nil
		}
	}

	fmt.Printf("Visited: %s\t\t", path)
	ext := filepath.Ext(path)

	if isImage(ext) {
		fmt.Print("Supported Image.\n")

		image := new(media)
		image.id = uuid()
		image.name = filepath.Base(path)
		image.path = path
		image.iconPath = filepath.Dir(path) + "/.icons/" + image.name

		fmt.Printf("Adding media:\n\tname: %s\n\tid: %s\n\tpath: %s\n", image.name, image.id, image.path)
		mediaFiles[image.id] = image

	} else {
		fmt.Print("Not a supported image type.\n")
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
		log.Fatal(err)
	}

	//n := bytes.IndexByte(out, 0)
	s := string(out)
	s = strings.TrimSpace(s)
	return s
}
