package main

import "os"

func main() {

	docxFile, err := os.Open("test.docx")
	if err != nil {
		panic(err)
	}

	defer docxFile.Close()

}
