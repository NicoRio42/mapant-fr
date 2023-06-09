package tiles

import (
	"errors"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
)

var MIN_X float64 = 343646
var MAX_X float64 = 1704354
var MIN_Y float64 = 5619537
var MAX_Y float64 = 7667537
var MAX_TILE_SIZE float64 = 500 * math.Pow(2, 12)

func lambert93ToTileNum(x, y, zoom int) (int, int) {
	var xTile = int(((float64(x) - MIN_X) / MAX_TILE_SIZE) * math.Pow(2, float64(zoom)))
	var yTile = int(((MAX_Y - float64(y)) / MAX_TILE_SIZE) * math.Pow(2, float64(zoom)))
	return xTile, yTile
}

func createDirIfDoesntExist(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(path, err)
		}
	}
}

func createPathIfDoesntExist(dirNames ...string) {
	for i := 1; i <= len(dirNames); i++ {
		createDirIfDoesntExist(filepath.Join(dirNames[0:i]...))
	}
}

func findParentTile(xChild, yChild int) (int, int) {
	var xParent int
	var yParent int

	// Might be overkill because of integer division behavior
	if (xChild-1)%2 == 0 {
		xParent = ((xChild - 1) / 2) + 1
	} else {
		xParent = ((xChild - 2) / 2) + 1
	}

	if (yChild-1)%2 == 0 {
		yParent = ((yChild - 1) / 2) + 1
	} else {
		yParent = ((yChild - 2) / 2) + 1
	}

	return xParent, yParent
}

func doesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)
	return !os.IsNotExist(error)
}

func getTileIfExistsOrTransparentImage(path string) image.Image {
	fmt.Println(doesFileExist(path), path)
	if doesFileExist(path) {
		img, err := mergi.Import(impexp.NewFileImporter(path))

		if err != nil {
			log.Fatal(err)
		}

		return img
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{int(TILE_PIXEL_SIZE), int(TILE_PIXEL_SIZE)}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	return img
}
