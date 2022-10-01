package utility

import (
	"encoding/xml"
	"os"
)

type Node struct {
	Id  int     `xml:"id,attr"`
	Lat float32 `xml:"lat,attr"`
	Lon float32 `xml:"lon,attr"`
}

type Way struct {
	Id  int  `xml:"id,attr"`
	Nd  []Nd `xml:"nd"`
	Tag byte `xml:"k,attr"`
}

type Nd struct {
	Ref int `xml:"ref,attr"`
}

var Nodes = make([]Node, 0)
var Ways = make([]Way, 0)
var Nds = make([]Nd, 0)

func decodeXml(xmlFile *os.File) {
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
				Nodes = append(Nodes, node)
			} else if tokenType.Name.Local == "way" {
				var way Way
				decoder.DecodeElement(&way, &tokenType)
				Ways = append(Ways, way)
			} else if tokenType.Name.Local == "nd" {
				var nd Nd
				decoder.DecodeElement(&nd, &tokenType)
				Nds = append(Nds, nd)
			}
		}
	}
}
