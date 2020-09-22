package graph

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

type graph struct {
	title         string
	width, height int32
	fscreen       bool
	running       bool
	glCtx         *sdl.GLContext
	win           *sdl.Window
	rend          Renderer
	figures       *figureList
}

func New(title string, w, h int32, fscreen bool) (*graph, error) {
	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_EVENTS)
	if err != nil {
		panic(err)
	}

	winFlags := uint32(sdl.WINDOW_ALLOW_HIGHDPI | sdl.WINDOW_OPENGL | sdl.WINDOW_RESIZABLE)
	if fscreen {
		winFlags = sdl.WINDOW_FULLSCREEN | winFlags
	}

	win, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, winFlags)
	if err != nil {
		panic(err)
	}

	glCtx, err := win.GLCreateContext()
	if err != nil {
		panic(err)
	}

	rend, err := NewGLRenderer(w, h, win)
	if err != nil {
		panic(err)
	}

	gl.Viewport(0, 0, w, h)

	return &graph{title, w, h, fscreen, false, &glCtx, win, rend, CreateFigureList()}, nil
}

func (g *graph) Run() error {
	defer g.Destroy()

	g.running = true

	for g.running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				g.running = false
			}
		}

		g.rend.Clear()

		for _, f := range g.figures.Slice() {
			if err := f.Render(g.rend); err != nil {
				return err
			}
		}

		g.rend.Update()
	}

	return nil
}

func (g *graph) Destroy() error {
	defer sdl.Quit()

	if err := g.win.Destroy(); err != nil {
		return err
	}

	for _, f := range g.figures.Slice() {
		if err := f.Destroy(); err != nil {
			return err
		}
	}

	return nil
}

func (g *graph) AppendFigure(f Figure) error {
	if err := g.figures.Append(f); err != nil {
		return fmt.Errorf("cannot append figure f to graph: %s", err)
	}

	return nil
}
