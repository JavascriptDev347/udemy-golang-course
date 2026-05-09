package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "Downloads/static/images"
	// Mkdir and MkdirAll create new directory
	if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
		log.Fatal(err)
	}
	// RemoveAll delete all files that belong to Downloads
	if err := os.RemoveAll("Downloads"); err != nil {
		log.Fatal(err)
	}

}
