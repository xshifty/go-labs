package graph

import "errors"

type Figure interface {
	Id() string
	Render(rend Renderer) error
	Destroy() error
}

type figureList struct {
	figures []Figure
}

func CreateFigureList() *figureList {
	return &figureList{[]Figure{}}
}

func (c *figureList) Append(f Figure) error {
	if f == nil {
		return errors.New("figure f cannot be nil")
	}

	c.figures = append(c.figures, f)

	return nil
}

func (c *figureList) Slice() []Figure {
	return c.figures
}
