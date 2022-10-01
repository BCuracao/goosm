package factories

import (
	"encoding/xml"
	"os"
)

type Way struct {
	Id  int  `xml:"id,attr"`
	Nd  []Nd `xml:"nd"`
	Tag byte `xml:"k,attr"`
}

type Nd struct {
	Ref int `xml:"ref,attr"`
}

func GetWays(xmlFile *os.File) []Way {
	ways := make([]Way, 0)
	decoder := xml.NewDecoder(xmlFile)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "way" {
				var way Way
				decoder.DecodeElement(&way, &tokenType)
				ways = append(ways, way)
			}
		}
	}
	return ways
}
