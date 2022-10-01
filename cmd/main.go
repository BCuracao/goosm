package main

import (
	"fmt"
	"os"

	Node "github.com/BCuracao/goosm/factories"
)

var nodes = make([]Node.Node, 0)

type Data struct {
}

func main() {

	docxFile, err := os.Open("D:\\go_workspace\\projects\\gotract\\osm\\dubai.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println("file successfully opened")
	defer docxFile.Close()

	nodes = Node.GetNodes(docxFile)

}
