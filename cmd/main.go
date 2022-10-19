package main

import (
	"fmt"
	"math/rand"
	"os"

	Utility "github.com/BCuracao/goosm/utility"
	"github.com/fogleman/gg"
)

var nodes []Utility.Node
var ways []Utility.Way
var bounds Utility.Bound

const W = 1024
const H = 1024

var dc = gg.NewContext(W, H)

func main() {
	xmlFile, err := os.Open("D:\\go_workspace\\projects\\gotract\\osm\\dubai_small.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println("File successfully opened")
	defer xmlFile.Close()

	Utility.DecodeXml(xmlFile)

	ways = Utility.Ways
	nodes = Utility.Nodes

	// Outputs for debugging purposes only
	fmt.Println("The node: ", nodes[400])
	fmt.Println("The way: ", ways[2])
	fmt.Println("The nd ref: ", ways[2].Nd)

	//dc := gg.NewContext(W, H)

	dc.SetRGB(0, 0, 0)
	dc.Clear()

	drawImage()
}

func drawBounds(bounds Utility.Bound) {
	lat1 := bounds.MinLat
	lon1 := bounds.MinLon
	x1, y1 := Utility.ConvertLatLonToXY(lat1, lon1)
	lat2 := bounds.MaxLat
	lon2 := bounds.MaxLon
	x2, y2 := Utility.ConvertLatLonToXY(lat2, lon2)

	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	a := rand.Float64()*0.5 + 0.5
	w := 0.2
	dc.SetRGBA(r, g, b, a)
	dc.SetLineWidth(w)
	dc.DrawLine(x1, y1, x2, y2)
	dc.Stroke()
}

func drawWays(ways []Utility.Way) {
	for i := 0; i < len(ways); i++ {
		for j := 1; j <= len(ways[i].Nd)-1; j++ {
			lat1 := ways[i].Nd[j-1].WayNode.Lat
			lon1 := ways[i].Nd[j-1].WayNode.Lon
			x1, y1 := Utility.ConvertLatLonToXY(lat1, lon1)
			lat2 := ways[i].Nd[j].WayNode.Lat
			lon2 := ways[i].Nd[j].WayNode.Lon
			x2, y2 := Utility.ConvertLatLonToXY(lat2, lon2)

			r := rand.Float64()
			g := rand.Float64()
			b := rand.Float64()
			a := rand.Float64()*0.5 + 0.5
			w := 0.2
			dc.SetRGBA(r, g, b, a)
			dc.SetLineWidth(w)
			dc.DrawLine(x1, y1, x2, y2)
			dc.Stroke()
		}
	}
}

func drawImage() {
	drawBounds(bounds)
	drawWays(ways)
	dc.SavePNG("out.png")
}
