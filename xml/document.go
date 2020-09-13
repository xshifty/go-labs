package xml

import (
	"bytes"
	"errors"
	"fmt"
)

type XMLDocument interface {
	XMLElement
	Version() string
	Encoding() string
	Root() xmlNode
}

type xmlDocument struct {
	version  string
	encoding string
	rootNode xmlNode
}

func (d *xmlDocument) Version() string {
	return d.version
}

func (d *xmlDocument) Encoding() string {
	return d.encoding
}

func (d *xmlDocument) Root() xmlNode {
	return d.rootNode
}

func (d *xmlDocument) Type() elementType {
	return ElementTypeDocument
}

func CreateDocument(source []byte) (XMLDocument, error) {
	var (
		doc = xmlDocument{
			version:  "1.0",
			encoding: "utf-8",
		}
		states        = parserStateStack{}
		elements      = elementStack{[]XMLElement{&doc}}
		token         = ""
		line          = 1
		position      = 0
		separator     = byte(' ')
		symbolParsers = map[byte]func(element XMLElement, states *parserStateStack, elements *elementStack) error{
			'>': parseGreaterThanStep,
			'<': parseLessThanStep,
			'"': parseDoubleQuoteStep,
			'=': parseEqualStep,
			'/': parseSlashStep,
		}
	)

	for _, b := range bytes.TrimSpace(source) {
		fmt.Printf("Document %#v\n\n", doc)
		currentElement := elements.Last()
		position++
		switch b {
		case '<', '>', '"', '/', '=':
			if len(token) > 0 {
				err := parseTokenStep(token, separator, currentElement, &states, &elements)
				if err != nil {
					return nil, fmt.Errorf("error found at line %d position %d: %s", line, position, err)
				}
			}

			token = ""

			err := symbolParsers[b](currentElement, &states, &elements)
			if err != nil {
				return nil, fmt.Errorf("error found at line %d position %d: %s", line, position, err)
			}

			//err = parseTokenStep(string(b), separator, currentElement, &states, &elements)
			//if err != nil {
			//	return nil, fmt.Errorf("error found at line %d position %d: %s", line, position, err)
			//}

			break

		case '\n', '\t', ' ':
			if b == '\n' {
				line = line + 1
				position = 0
			}

			if len(token) > 0 {
				err := parseTokenStep(token, separator, &doc, &states, &elements)
				if err != nil {
					return nil, fmt.Errorf("error found at line %d position %d: %s", line, position, err)
				}
			}

			token = ""
			separator = b
			break

		default:
			token = token + string(b)
		}
	}

	fmt.Printf("---[ FINAL DOCUMENT STRUCTURE ]---\n%#v\n---[ FINAL DOCUMENT STRUCTURE ]---\n\n", doc)

	return &doc, nil
}

func parseSlashStep(element XMLElement, states *parserStateStack, elements *elementStack) error {
	if !states.IsEmpty() {
		if states.Last().Is(parserStateTagOpenStarted) {
			states.Push(parserStateTagCloseNamingStarted)

			for i := 0; i < elements.Length(); i++ {
				if !elements.Last().Type().Is(ElementTypeNode) {
					elements.Pop()
					continue
				}

				return nil
			}
		}
	}

	return errors.New("unexpected slash symbol")
}

func parseDoubleQuoteStep(element XMLElement, states *parserStateStack, elements *elementStack) error {
	if !states.IsEmpty() {
		return nil
	}

	return errors.New("unexpected double quote symbol")
}

func parseEqualStep(element XMLElement, states *parserStateStack, elements *elementStack) error {
	fmt.Printf("Parsing slash token\nState stack\t%+v\nElement stack\t%#v\n\n", states.stack, elements.Last())

	if !states.IsEmpty() {
		return nil
	}

	return errors.New("unexpected equal symbol")
}

func parseLessThanStep(element XMLElement, states *parserStateStack, elements *elementStack) error {
	fmt.Printf("Parsing less than token\nState stack\t%+v\nElement stack\t%#v\n\n", states.stack, elements.Last())

	if states.IsEmpty() {
		states.Push(parserStateTagOpenStarted)

		return nil
	}

	if states.Last().Is(parserStateTagOpenEnded) {
		states.Push(parserStateTagOpenStarted)

		return nil
	}

	if states.Last().Is(parserStateTagCloseEnded) {
		states.Push(parserStateTagOpenStarted)

		return nil
	}

	return errors.New("unexpected less than symbol (<)")
}

func parseGreaterThanStep(element XMLElement, states *parserStateStack, elements *elementStack) error {
	fmt.Printf("Parsing greater than token\nState stack\t%+v\nElement stack\t%#v\n\n", states.stack, elements.Last())

	if states.Last().Is(parserStateTagOpenNamingEnded) {
		states.Push(parserStateTagOpenEnded)

		return nil
	}

	if states.Last().Is(parserStateTagCloseNamingEnded) {
		states.Push(parserStateTagCloseEnded)
		elements.Pop()

		return nil
	}

	return errors.New("unexpected greater than symbol (>)")
}

func parseTokenStep(token string, separator byte, element XMLElement, states *parserStateStack, elements *elementStack) error {
	fmt.Printf("Parsing token\t%s\nState stack\t%+v\nElement stack\t%#v\n\n", token, states.stack, elements.Last())

	if !states.IsEmpty() {
		if element.Type().Is(ElementTypeDocument) {
			doc := element.(*xmlDocument)

			if states.Last().Is(parserStateTagOpenStarted) {
				doc.rootNode = createNode(token)
				elements.Push(&doc.rootNode)
				states.Push(parserStateTagOpenNamingEnded)

				return nil
			}
		}

		if element.Type().Is(ElementTypeNode) {
			if states.Last().Is(parserStateTagOpenStarted) {
				newElement := createNode(token)

				elements.Push(&newElement)
				element.(*xmlNode).Children().Push(elements.Last())

				states.Push(parserStateTagOpenNamingEnded)

				return nil
			}

			if states.Last().Is(parserStateTagCloseStarted) {
				nodeName := element.(*xmlNode).Name()
				if nodeName != token {
					return fmt.Errorf("unexpected closing tag name '%s': expected '%s' instead", token, nodeName)
				}

				return nil
			}

			if states.Last().Is(parserStateTagCloseNamingStarted) {
				nodeName := element.(*xmlNode).Name()
				if nodeName != token {
					return fmt.Errorf("unexpected closing tag name '%s': expected '%s' instead", token, nodeName)
				}

				states.Push(parserStateTagCloseNamingEnded)

				return nil
			}
		}
	}

	return fmt.Errorf("unexpected text token %s", token)
}
