package main

import (
	"fmt"

	"github.com/xshifty/go-labs/xml"
)

func main() {
	xmlString := `
<svg>
	<rect></rect>
	<circle></circle>
</svg>
`

	document, err := xml.CreateDocument([]byte(xmlString))
	if err != nil {
		panic(err)
	}

	root := document.Root()

	fmt.Printf("- %s\n", root.Name())

	for _, e := range root.Children().All() {
		switch e.Type() {
		case xml.ElementTypeNode:
			fmt.Printf("\t- %s\n", e.(xml.XMLElementWithName).Name())
			break
		}
	}

	fmt.Println(document)
}
