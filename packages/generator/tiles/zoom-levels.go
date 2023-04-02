package tiles

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
)

func CreateBaseZoomLevel(inputDir string, inputZoom int) {
	fmt.Println("Base level")
	files, err := ioutil.ReadDir(inputDir)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		length := len(file.Name())

		if file.Name()[length-8:length] == "depr.png" && !file.IsDir() {
			x, err := strconv.Atoi(file.Name()[11:15])

			if err != nil {
				log.Fatal(err)
			}

			y, err := strconv.Atoi(file.Name()[16:20])

			if err != nil {
				log.Fatal(err)
			}

			xTile, yTile := lambert93ToTileNum(x*1000, y*1000, inputZoom)
			inputFile, err := ioutil.ReadFile(filepath.Join("in", file.Name()))

			if err != nil {
				fmt.Println(err)
				return
			}

			createPathIfDoesntExist("out", fmt.Sprint(inputZoom), fmt.Sprint(xTile))
			destination := filepath.Join("out", fmt.Sprint(inputZoom), fmt.Sprint(xTile), fmt.Sprint(yTile)+".png")
			err = ioutil.WriteFile(destination, inputFile, 0644)

			if err != nil {
				fmt.Println("Error creating", destination)
				fmt.Println(err)
				return
			}
		}

	}
}

func loopOnLevel(zoomLevel int, callback func(zoomLevel int, xdir, png string)) {
	xDirs, err := ioutil.ReadDir(filepath.Join("out", fmt.Sprint(zoomLevel)))

	if err != nil {
		log.Fatal(err)
	}

	for _, xDir := range xDirs {
		if !xDir.IsDir() {
			continue
		}

		pngs, err := ioutil.ReadDir(filepath.Join("out", fmt.Sprint(zoomLevel), fmt.Sprint(xDir.Name())))

		if err != nil {
			log.Fatal(err)
		}

		for _, png := range pngs {
			if png.IsDir() {
				continue
			}

			callback(zoomLevel, xDir.Name(), png.Name())
		}
	}
}

func CreateLowerZoomLevel(zoomLevel int) {
	fmt.Println("Create zoom", zoomLevel)

	loopOnLevel(zoomLevel-1, func(zoomLvl int, xDir, png string) {
		pngPath := filepath.Join("out", fmt.Sprint(zoomLvl), xDir, png)
		img, err := mergi.Import(impexp.NewFileImporter(pngPath))

		if err != nil {
			log.Fatal(err)
		}

		topLeftImage, _ := mergi.Crop(img, image.Pt(0, 0), image.Pt(img.Bounds().Max.X/2, img.Bounds().Max.Y/2))
		topRightImage, _ := mergi.Crop(img, image.Pt(img.Bounds().Max.X/2, 0), image.Pt(img.Bounds().Max.X/2, img.Bounds().Max.Y/2))
		bottomLeftImage, _ := mergi.Crop(img, image.Pt(0, img.Bounds().Max.Y/2), image.Pt(img.Bounds().Max.X/2, img.Bounds().Max.Y/2))
		bottomRightImage, _ := mergi.Crop(img, image.Pt(img.Bounds().Max.X/2, img.Bounds().Max.Y/2), image.Pt(img.Bounds().Max.X/2, img.Bounds().Max.Y/2))

		xTile, _ := strconv.Atoi(xDir)
		yTile, _ := strconv.Atoi(strings.Replace(png, ".png", "", 1))

		xTopLeft := 2*(xTile-1) + 1
		yTopLeft := 2*(yTile-1) + 1

		topLeftPath := []string{"out", fmt.Sprint(zoomLvl + 1), fmt.Sprint(xTopLeft), fmt.Sprint(yTopLeft) + ".png"}
		createPathIfDoesntExist("out", fmt.Sprint(zoomLvl+1), fmt.Sprint(xTopLeft))
		mergi.Export(impexp.NewFileExporter(topLeftImage, filepath.Join(topLeftPath...)))

		topRightPath := []string{"out", fmt.Sprint(zoomLvl + 1), fmt.Sprint(xTopLeft + 1), fmt.Sprint(yTopLeft) + ".png"}
		createPathIfDoesntExist("out", fmt.Sprint(zoomLvl+1), fmt.Sprint(xTopLeft+1))
		mergi.Export(impexp.NewFileExporter(topRightImage, filepath.Join(topRightPath...)))

		bottomLeftPath := []string{"out", fmt.Sprint(zoomLvl + 1), fmt.Sprint(xTopLeft), fmt.Sprint(yTopLeft+1) + ".png"}
		createPathIfDoesntExist("out", fmt.Sprint(zoomLvl+1), fmt.Sprint(xTopLeft))
		mergi.Export(impexp.NewFileExporter(bottomLeftImage, filepath.Join(bottomLeftPath...)))

		bottomRightPath := []string{"out", fmt.Sprint(zoomLvl + 1), fmt.Sprint(xTopLeft + 1), fmt.Sprint(yTopLeft+1) + ".png"}
		createPathIfDoesntExist("out", fmt.Sprint(zoomLvl+1), fmt.Sprint(xTopLeft+1))
		mergi.Export(impexp.NewFileExporter(bottomRightImage, filepath.Join(bottomRightPath...)))
	})
}

func ResizeZoomLevel(zoomLevel int, isLast bool) {
	fmt.Println("Resizing", zoomLevel)

	loopOnLevel(zoomLevel, func(zoomLvl int, xDir, png string) {
		pngPath := filepath.Join("out", fmt.Sprint(zoomLvl), xDir, png)
		img, err := mergi.Import(impexp.NewFileImporter(pngPath))

		if err != nil {
			log.Fatal(err)
		}

		resizedImg, _ := mergi.Resize(img, TILE_PIXEL_SIZE, TILE_PIXEL_SIZE)

		if isLast {
			highQualityPngPath := strings.Replace(pngPath, ".png", "", 1) + "_hd.png"
			os.Rename(pngPath, highQualityPngPath)
		}

		mergi.Export(impexp.NewFileExporter(resizedImg, pngPath))
	})
}

func CreateUpperZoomLevel(zoomLevel int) {
	fmt.Println("Creating zoom", zoomLevel)

	tilesToGenerate := make(map[int]map[int]int)

	loopOnLevel(zoomLevel+1, func(zoomLvl int, xDir, png string) {
		xTile, _ := strconv.Atoi(xDir)
		yTile, _ := strconv.Atoi(strings.Replace(png, ".png", "", 1))

		xParent, yParent := findParentTile(xTile, yTile)
		_, xfound := tilesToGenerate[xParent]

		if !xfound {
			tilesToGenerate[xParent] = make(map[int]int)
		}

		_, yfound := tilesToGenerate[xParent][yParent]

		if !yfound {
			tilesToGenerate[xParent][yParent] = 0
		}
	})

	for xTile, yTilesMap := range tilesToGenerate {
		for yTile, _ := range yTilesMap {
			xTopLeft := 2*(xTile-1) + 1
			yTopLeft := 2*(yTile-1) + 1
		}
	}
}
