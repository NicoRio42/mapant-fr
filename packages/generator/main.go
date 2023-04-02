package main

import (
	"log"
	"os"
	"path/filepath"

	"mapant-fr-generator/tiles"
)

func main() {
	workingDir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	tiles.CreateTiles(filepath.Join(workingDir, "in"), 11, 1, 13)
}
