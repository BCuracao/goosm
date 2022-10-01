package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	Utility "github.com/BCuracao/goosm/utility"
	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {

	xmlFile, err := os.Open("/Users/bastianbreindl/go/src/projects/goosm/osm/dubai.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println("File successfully opened")
	defer xmlFile.Close()

	Utility.DecodeXml(xmlFile)

	nodes := Utility.Nodes
	ways := Utility.Ways

	// Outputs for debugging purposes only
	fmt.Println("The node: ", nodes[4000])
	fmt.Println("The way: ", ways[2])
	fmt.Println("The nd ref: ", ways[2].Nd)

	drawImage(ways)
}

func drawImage(ways []Utility.Way) {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	for i := 0; i < len(ways); i++ {
		for j := 1; j <= len(ways[i].Nd)-1; j++ {
			// Draw a closed shape
			gc.BeginPath()    // Initialize a new path
			gc.MoveTo(10, 10) // Move to a position to start the new path
			gc.LineTo(ways[i].Nd[j-1].Ref, ways[i].Nd[j].Ref)
			gc.Close()
			gc.FillStroke()
		}
	}

	// Save to file
	draw2dimg.SaveToPngFile("hello.png", dest)
}
