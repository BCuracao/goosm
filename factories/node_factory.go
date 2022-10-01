package factories

import "os"

type Node struct {
	Id  int     `xml:"id,attr"`
	Lat float32 `xml:"lat,attr"`
	Lon float32 `xml:"lon,attr"`
}

/*
func GetNodes(docxFile *os.File) []Node {
	decoder := xml.NewDecoder(docxFile)
	nodes := make([]Node, 0)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "node" {
				var node Node
				decoder.DecodeElement(&node, &se)
				nodes = append(nodes, node)
			}
		}
	}
	return nodes
}
*/

func GetNodes(docxFile *os.File)
