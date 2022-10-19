package utility

import (
	"encoding/xml"
	"os"
)

type Node struct {
	Id  float64 `xml:"id,attr"`
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

type Way struct {
	Id   int   `xml:"id,attr"`
	Nd   []Nd  `xml:"nd"`
	Tags []Tag `xml:"tag"`
}

type Nd struct {
	Ref float64 `xml:"ref,attr"`
	//TODO Add the relevant nodes here
	WayNode Node
}

type Bound struct {
	MinLat float64 `xml:"minlat,attr"`
	MinLon float64 `xml:"minlon,attr"`
	MaxLat float64 `xml:"maxlat,attr"`
	MaxLon float64 `xml:"maxlon,attr"`
}

type Tag struct {
	K string `xml:"k,attr"`
	V string `xml:"v,attr"`
}

var Nodes = make([]Node, 0)
var Ways = make([]Way, 0)
var Bounds Bound

func DecodeXml(xmlFile *os.File) {
	decoder := xml.NewDecoder(xmlFile)

	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "bounds" {
				var bounds Bound
				decoder.DecodeElement(&bounds, &tokenType)
				Bounds = bounds
			} else if tokenType.Name.Local == "node" {
				var node Node
				decoder.DecodeElement(&node, &tokenType)
				Nodes = append(Nodes, node)
			} else if tokenType.Name.Local == "way" {
				var way Way
				decoder.DecodeElement(&way, &tokenType)
				Ways = append(Ways, way)
			}
		}
	}
	FindWayNdRefValueInNodesId()
}

func FindWayNdRefValueInNodesId() {
	for i := 0; i < len(Ways); i++ {
		for j := 0; j < len(Ways[i].Nd); j++ {
			for k := 0; k < len(Nodes); k++ {
				if Ways[i].Nd[j].Ref == Nodes[k].Id {
					Ways[i].Nd[j].WayNode = Nodes[k]
				}
			}
		}
	}
}
