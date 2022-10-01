package factories

import (
	"encoding/xml"
	"os"
)

type Node struct {
	Id  int     `xml:"id,attr"`
	Lat float32 `xml:"lat,attr"`
	Lon float32 `xml:"lon,attr"`
}

func GetNodes(xmlFile *os.File) []Node {
	nodes := make([]Node, 0)
	decoder := xml.NewDecoder(xmlFile)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "node" {
				var node Node
				decoder.DecodeElement(&node, &tokenType)
				nodes = append(nodes, node)
			}
		}
	}
	return nodes
}
