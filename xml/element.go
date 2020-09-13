package xml

type elementType string

func (t elementType) Is(typeName elementType) bool {
	return t == typeName
}

const (
	ElementTypeNode      elementType = "node"
	ElementTypeAttribute elementType = "attribute"
	ElementTypeText      elementType = "text"
	ElementTypeDocument  elementType = "document"
)

type XMLElement interface {
	Type() elementType
}

type XMLElementWithName interface {
	XMLElement
	Name() string
}

type XMLElementWithValue interface {
	XMLElement
	Value() string
}

type XMLElementWithAttributes interface {
	XMLElement
	Attributes() []xmlAttribute
}

type XMLElementWithNameAndAttributes interface {
	XMLElementWithName
	XMLElementWithAttributes
}

type elementStack struct {
	stack []XMLElement
}

func (s *elementStack) Length() int {
	return len(s.stack)
}

func (s *elementStack) IsEmpty() bool {
	return s.Length() < 1
}

func (s *elementStack) Last() XMLElement {
	if s.IsEmpty() {
		return nil
	}

	return s.stack[s.Length() - 1]
}

func (s *elementStack) Push(element XMLElement) {
	s.stack = append(s.stack, element)
}

func (s *elementStack) Pop() XMLElement {
	if s.IsEmpty() {
		return nil
	}

	e := s.Last()
	s.stack = s.stack[:s.Length() - 1]

	return e
}
