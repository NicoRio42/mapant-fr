package tiles

func CreateTiles(inputDir string, inputZoom, minZoom, maxZoom int) {

	// CreateBaseZoomLevel(inputDir, inputZoom)

	// for zoomLevel := inputZoom + 1; zoomLevel <= maxZoom; zoomLevel++ {
	// 	CreateLowerZoomLevel(zoomLevel)
	// }

	// for zoomLevel := inputZoom; zoomLevel <= maxZoom; zoomLevel++ {
	// 	ResizeZoomLevel(zoomLevel, zoomLevel == maxZoom)
	// }

	CreateUpperZoomLevel(10)
	// for zoomLevel := inputZoom - 1; zoomLevel <= maxZoom; zoomLevel-- {
	// 	CreateUpperZoomLevel(zoomLevel)
	// }
}
