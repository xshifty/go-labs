package graph

import (
	"errors"
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer interface {
	Clear()
	Update()
}

type glRenderer struct {
	width, height int32
	win           *sdl.Window
}

func NewGLRenderer(w, h int32, win *sdl.Window) (*glRenderer, error) {
	if win == nil {
		return nil, errors.New("win pointer must not be nil")
	}

	if err := gl.Init(); err != nil {
		return nil, fmt.Errorf("cannot initiate opengl: %s", err)
	}

	return &glRenderer{w, h, win}, nil
}

func (r *glRenderer) Clear() {
	gl.ClearColor(0.0, 0.0, 0.5, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (r *glRenderer) Update() {
	r.win.GLSwap()
}
