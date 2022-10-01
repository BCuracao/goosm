package main

import (
	"fmt"
	"os"

	Factory "github.com/BCuracao/goosm/factories"
)

var nodes = make([]Factory.Node, 0)
var ways = make([]Factory.Way, 0)

type Data struct {
}

func main() {

	xmlFile, err := os.Open("/Users/bastianbreindl/go/src/projects/goosm/osm/dubai.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println("File successfully opened")
	defer xmlFile.Close()

	//nodes = Factory.GetNodes(xmlFile)
	//ways = Factory.GetWays(xmlFile)

	nodes = Utility.decodeXml(xmlFile)

	fmt.Println(nodes[4000])
	fmt.Println(ways[2])
}
