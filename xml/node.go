package xml

type xmlElementCollection struct {
	elements []XMLElement
}

func (c *xmlElementCollection) Length() int {
	return len(c.elements)
}

func (c *xmlElementCollection) IsEmpty() bool {
	return c.Length() == 0
}

func (c *xmlElementCollection) Last() XMLElement {
	return c.elements[c.Length() - 1]
}

func (c *xmlElementCollection) Push(element XMLElement) {
	c.elements = append(c.elements, element)
}

func (c *xmlElementCollection) Pop() XMLElement {
	e := c.Last()

	c.elements = c.elements[:c.Length() - 1]

	return e
}

func (c *xmlElementCollection) All() []XMLElement {
	return c.elements
}

type xmlNode struct {
	name       string
	attributes []xmlAttribute
	children   xmlElementCollection
}

func createNode(name string) xmlNode {
	return xmlNode{name: name}
}

func (n *xmlNode) Type() elementType {
	return ElementTypeNode
}

func (n *xmlNode) Name() string {
	return n.name
}

func (n *xmlNode) Value() string {
	return ""
}

func (n *xmlNode) Children() *xmlElementCollection {
	return &n.children
}
