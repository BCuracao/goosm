package factories

type Way struct {
	Id  int  `xml:"id,attr"`
	Nd  []Nd `xml:"nd"`
	Tag byte `xml:"k,attr"`
}

type Nd struct {
	Ref int `xml:"ref,attr"`
}
