package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type media struct {
	name     string
	id       string
	path     string
	iconPath string
}

var mediaFiles map[string]media

func loadMedia(path string) {
	if mediaFiles == nil {
		mediaFiles = make(map[string]media)
	}

	err := filepath.Walk(path, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	fmt.Printf("Visited: %s\t\t", path)
	ext := filepath.Ext(path)

	if isImage(ext) {
		fmt.Print("Supported Image.\n")
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
